package main

import (
	"log"

	_ "github.com/peetwerapat/learnhub-go-api/docs"
	"github.com/peetwerapat/learnhub-go-api/internal/infrastructure/db"
	"github.com/peetwerapat/learnhub-go-api/internal/infrastructure/di"
	"github.com/peetwerapat/learnhub-go-api/internal/infrastructure/router"
	"github.com/peetwerapat/learnhub-go-api/pkg/config"
)

// @title LearnHub Go API
// @version 1.0
// @description This is LearnHub API documentation.
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Println("Starting application")

	cfg := config.Load()

	dbConn, err := db.ConnectDatabase(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// entities := []interface{}{}
	// if err := db.AutoMigrate(dbConn, entities...); err != nil {
	// 	log.Fatalf("Database migration failed: %v", err)
	// }

	appUC := di.InitApp(dbConn)

	r := router.InitRouter(appUC)

	log.Printf("Server starting on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
