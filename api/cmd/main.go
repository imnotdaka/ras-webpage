package main

import (
	"log/slog"
	"os"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/cmd/config"

	"github.com/imnotdaka/RAS-webpage/internal/clients/mercadopago"
	"github.com/imnotdaka/RAS-webpage/internal/database"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/handlers"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/plan"
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
		Requester:   cfg.HTTPClient,
	})
	if mpclient == nil {
		slog.Info("failed to initialize MercadoPago client")
		return mercadopago.ErrEmptyClient
	}

	auth := authenticator.NewAuth(os.Getenv(cfg.JWTSecret))

	app := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	app.Use(cors.New(config))

	authorizedRoutes := app.Group("/", handlers.JWTMiddleware(auth))

	app.POST("/user", handlers.CreateUserHandler(user.NewRepo(db), auth))
	app.POST("/auth/user", handlers.Login(user.NewRepo(db), auth))
	authorizedRoutes.PUT("/user/:id", handlers.UpdateUserHandler(user.NewRepo(db)))
	authorizedRoutes.DELETE("/user/:id", handlers.DeleteUserHandler(user.NewRepo(db)))

	authorizedRoutes.POST("/preapproval_plan", handlers.CreatePlanHandler(mpclient, plan.NewRepo(db)))
	authorizedRoutes.GET("/get_plans", handlers.GetAllPlanHandler(plan.NewRepo(db)))
	authorizedRoutes.POST("/create_suscription", handlers.CreateSubscriptionHandler(mpclient, plan.NewRepo(db), subscription.NewRepo(db)))
	authorizedRoutes.GET("/subscription/:id", handlers.GetSubscriptionHandler(mpclient))
	authorizedRoutes.PUT("/subscription", handlers.UpdateMPSubscriptionHandler(mpclient))
	app.POST("/webhook", handlers.Webhook(mpclient))

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}
