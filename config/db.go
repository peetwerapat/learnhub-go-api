package config

import (
	"fmt"
	"os"

	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = database.AutoMigrate(&domain.User{}, &domain.Content{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	DB = database
	return nil
}
