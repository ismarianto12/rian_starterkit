package main

import (
	"log"
	"os"
	"raenRestApi/config"
	"raenRestApi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables from .env file
	if err := os.Setenv("GIN_MODE", "release"); err != nil {
		log.Fatal("Error setting GIN_MODE:", err)
	}

	// Initialize database connection
	config.Connect()
	db := config.GetDB()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
