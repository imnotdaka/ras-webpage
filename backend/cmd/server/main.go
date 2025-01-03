package main

import (
	"log/slog"
	"net/http"
	"os"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/cmd/server/config"

	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
	"github.com/imnotdaka/RAS-webpage/internal/database"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/handlers"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/session"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/subscription"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
	mpconfig "github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	db, err := database.NewDB(cfg.DB)
	if err != nil {
		return err
	}

	mpclient := mercadopago.NewClient(&mpconfig.Config{
		AccessToken: cfg.MPAccessToken,
		Requester:   &http.Client{},
	})
	if mpclient == nil {
		slog.Info("failed to initialize MercadoPago client")
		return mercadopago.ErrEmptyClient
	}

	userRepo := user.NewRepo(db)
	planRepo := plan.NewRepo(db)
	subsRepo := subscription.NewRepo(db)
	sessionRepo := session.NewRepo(db)

	auth := authenticator.NewAuth(cfg.AccessTokenSecret, cfg.RefreshTokenSecret, sessionRepo)

	app := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	app.Use(cors.New(config))

	authorizedRoutes := app.Group("/", handlers.JWTMiddleware(auth))

	app.POST("/user", handlers.CreateUserHandler(userRepo, auth))
	app.POST("/auth/user", handlers.LoginHandler(userRepo, auth))
	app.GET("/auth/refresh", handlers.RefreshHandler(auth))
	app.GET("/auth/me", handlers.AuthMeHandler(userRepo, auth))
	authorizedRoutes.PUT("/user/:id", handlers.UpdateUserHandler(userRepo))
	authorizedRoutes.DELETE("/user/:id", handlers.DeleteUserHandler(userRepo))
	authorizedRoutes.POST("/auth/logout", handlers.LogOutHandler(sessionRepo))

	authorizedRoutes.POST("/preapproval_plan", handlers.CreatePlanHandler(mpclient, planRepo))
	authorizedRoutes.GET("/get_plans", handlers.GetAllPlanHandler(planRepo))
	authorizedRoutes.POST("/create_suscription", handlers.CreateSubscriptionHandler(mpclient, planRepo, subsRepo))
	authorizedRoutes.GET("/subscription/:id", handlers.GetSubscriptionHandler(mpclient))
	authorizedRoutes.GET("/subscription", handlers.GetSubscriptionByUserIDHandler(subsRepo))
	authorizedRoutes.PUT("/cancel_subscription", handlers.CancelMPSubscriptionHandler(mpclient, subsRepo))
	app.POST("/webhook", handlers.WebhookHandler(mpclient, subsRepo, planRepo, cfg))

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}
