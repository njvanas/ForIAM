package main

import (
	"log"
	"os"

	"github.com/ForIAM/ForIAM/backend/internal/api"
	"github.com/ForIAM/ForIAM/backend/internal/config"
	"github.com/ForIAM/ForIAM/backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Seed initial data
	if err := database.Seed(db); err != nil {
		log.Println("Warning: Failed to seed database:", err)
	}

	// Set Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize API server
	server := api.NewServer(db, cfg)
	
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting server on port %s", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}