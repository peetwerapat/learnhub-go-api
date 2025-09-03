package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	AppPort    string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		AppPort:    os.Getenv("APP_PORT"),
	}

	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}

	return cfg
}
