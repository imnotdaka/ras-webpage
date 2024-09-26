package main

import (
	"os"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/cmd/config"
	"github.com/imnotdaka/RAS-webpage/internal/database"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/authenticator"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/handlers"
	"github.com/imnotdaka/RAS-webpage/internal/rautosport/user"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	db, err := database.NewDB(cfg.DB)
	if err != nil {
		return err
	}

	JWT_SECRET := "JWT_SECRET"
	auth := authenticator.NewAuth(os.Getenv(JWT_SECRET))

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

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}
