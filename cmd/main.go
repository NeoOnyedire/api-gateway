package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/config"
	"api-gateway/internal/handlers"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
)

func main() {
	config.LoadEnv()

	port := config.GetEnv("PORT", "8080")
	target := config.GetEnv("TARGET_SERVICE", "http://localhost:9000")

	router := gin.Default()

	// Middleware
	router.Use(middleware.Logger())

	// Health
	router.GET("/health", handlers.HealthCheck)

	// Proxy route
	router.Any("/api/*path", proxy.ReverseProxy(target))

	log.Println("Gateway running on port:", port)
	log.Println("Forwarding requests to:", target)

	router.Run(":" + port)
}