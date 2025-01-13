// cmd/server/main.go
package main

import (
	"bac/internal/api"
	"bac/internal/config"
	"bac/internal/database"
	"bac/internal/utils"
)

func main() {
	// Initialize logger
	logger := utils.NewLogger()
	defer logger.Cleanup()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration:", err)
	}

	// Run migrations first
	if err := database.RunMigrations(cfg.DatabaseURL); err != nil {
		logger.Fatal("Failed to run migrations:", err)
	}

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	// Initialize and start server
	server := api.NewServer(db, cfg)
	if err := server.Start(); err != nil {
		logger.Fatal("Server failed:", err)
	}
}
