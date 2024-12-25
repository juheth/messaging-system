package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juheth/messaging-system/internal/infrastructure/database"
)

func LoadConfig() database.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return database.Config{
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Database: os.Getenv("DATABASE_NAME"),
	}
}
