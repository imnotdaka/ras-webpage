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

	app := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	app.Use(cors.New(config))

	app.POST("/user", handlers.CreateUserHandler(user.NewRepo(db), auth))
	app.POST("/auth/user", handlers.Login(user.NewRepo(db), auth))
	app.POST("/auth/jwt", handlers.JWTLogin(user.NewRepo(db), auth))
	app.PUT("/user/:id", handlers.UpdateUserHandler(user.NewRepo(db)))
	app.DELETE("/user/:id", handlers.DeleteUserHandler(user.NewRepo(db)))

	app.POST("/preapproval_plan", handlers.CreatePlanHandler(mpclient, plan.NewRepo(db)))
	app.GET("/get_plans", mercadopago.GetAll(plan.NewRepo(db)))
	app.POST("/create_suscription", handlers.CreateSuscriptionHandler(mpclient, plan.NewRepo(db)))
	app.POST("/webhook", mercadopago.Webhook())

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}
