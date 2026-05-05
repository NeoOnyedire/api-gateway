package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"api-gateway/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	// Health check
	router.GET("/health", handlers.HealthCheck)

	log.Println("Starting API Gateway on port:", port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}