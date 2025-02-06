package main

import (
	"bac/internal/api"
	"bac/internal/config"
	"bac/internal/database"
	"bac/internal/utils"
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
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

	// Set up migration path
	baseDir, _ := filepath.Abs(".")
	migrationPath := "file://" + filepath.Join(baseDir, "internal", "database", "migration")

	// Run migrations
	if err := database.RunMigrations(cfg.DatabaseURL, migrationPath); err != nil {
		logger.Fatal("Failed to run migrations:", err)
	}

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to initialize database:", err)
	}

	// Initialize server
	server := api.NewServer(db, cfg)

	// Setup graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Start(); err != nil {
			logger.Fatal("Server failed:", err)
		}
	}()

	<-stop
	logger.Info("Shutting down server...")

	// Gracefully shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Failed to shut down server gracefully:", err)
	}

	logger.Info("Server shut down successfully.")
}
