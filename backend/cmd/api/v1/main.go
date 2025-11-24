package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "backend/docs" // Swagger docs
	"backend/cmd/api/v1/handlers"
	"backend/cmd/api/v1/routes"
)

// @title           Application Platform API
// @version         1.0
// @description     Backend API for Application Platform
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, healthHandler)

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}