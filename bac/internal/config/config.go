// internal/config/config.go
package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
}

func Load() (*Config, error) {
	// Load .env file
	envFile := ".env.local"
	if os.Getenv("APP_ENV") == "production" {
		envFile = ".env.production"
	}
	godotenv.Load(envFile)

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        getEnvWithDefault("PORT", "8080"),
		Environment: getEnvWithDefault("ENV", "development"),
	}, nil
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
