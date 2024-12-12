package main

import (
	"log"

	"auth0-api-service/config"
	"auth0-api-service/handlers"
	"auth0-api-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	authService := services.NewAuthService(cfg)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Setup router
	r := gin.Default()

	// Routes
	r.POST("/signup", authHandler.Signup)
	r.POST("/signin", authHandler.Signin)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
