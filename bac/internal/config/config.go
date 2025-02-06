package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
}

func Load() (*Config, error) {
	// Load environment variables from .env file
	envFile := ".env.local"
	if os.Getenv("APP_ENV") == "production" {
		envFile = ".env.production"
	}

	if err := godotenv.Load(envFile); err != nil {
		return nil, fmt.Errorf("failed to load %s: %w", envFile, err)
	}

	// Ensure mandatory variables are present
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required but not set")
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        getEnvWithDefault("PORT", "3000"),
		Environment: getEnvWithDefault("ENV", "development"),
	}, nil
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
