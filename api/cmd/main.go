package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imnotdaka/RAS-webpage/cmd/config"
	"github.com/imnotdaka/RAS-webpage/internal/database"
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

	app := gin.New()

	app.POST("/user", handlers.CreateUserHandler(user.NewRepo(db)))
	app.GET("/user/:id", handlers.WithJWTAuth(handlers.GetUserByIdHandler(user.NewRepo(db)), user.NewRepo(db)))
	app.PUT("/user/:id", handlers.UpdateUserHandler(user.NewRepo(db)))
	app.DELETE("/user/:id", handlers.DeleteUserHandler(user.NewRepo(db)))

	err = app.Run()
	if err != nil {
		return err
	}

	return nil
}
