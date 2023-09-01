package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/Cristi571/golang_api/config"
	"github.com/Cristi571/golang_api/handlers"
	"github.com/Cristi571/golang_api/repos"
	"github.com/Cristi571/golang_api/services"
)

func main() {
	server := echo.New()

	// load config
	config := config.Load()
	userRepo := repos.NewUserRepository(config.DbConn)
	userService := services.NewUserService(userRepo)

	healthHandler := handlers.NewHealthHandler()
	// server.GET("/", healthHandler.HomePage)
	server.GET("/live", healthHandler.IsAlive)

	// REMOVE THAT ENDPOINT
	userHandler := handlers.NewUserHandler(userService)
	server.GET("/users/:id", userHandler.Get)

	// TODO: Register a new endpoint for POST user
	// userHandler := handlers.NewUserHandler(userService)
	server.POST("users/:id", userHandler.Post)

	if err := server.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
