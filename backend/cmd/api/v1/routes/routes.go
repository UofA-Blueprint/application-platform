package routes

import (
	"github.com/gin-gonic/gin"
	"backend/cmd/api/v1/handlers"
)

func SetupRoutes(r *gin.Engine, healthHandler *handlers.HealthHandler) {
	// Root route
	r.GET("/", healthHandler.HelloWorld)
	
	// Health check route
	r.GET("/health", healthHandler.HealthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", healthHandler.HelloWorld)
		v1.GET("/health", healthHandler.HealthCheck)
	}
}