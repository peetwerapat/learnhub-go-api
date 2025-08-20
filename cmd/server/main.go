package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/peetwerapat/learnhub-go-api/config"
	_ "github.com/peetwerapat/learnhub-go-api/docs"
	"github.com/peetwerapat/learnhub-go-api/pkg/router"
)

// @title LearnHub Go API
// @version 1.0
// @description This is LearnHub API documentation.
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	log.Println("Starting application")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	r := router.InitRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
