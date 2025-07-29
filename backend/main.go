package main

import (
	"log/slog"
	"net/http"
	"os"
	"trip-planner-backend/config"
	"trip-planner-backend/db"

	"github.com/gin-gonic/gin"
)

func init() {
	// Load config
	err := config.Load()
	if err != nil {
		slog.Error("Could not load environment variables", "error", err.Error())
		os.Exit(1)
	}

	// Configure logger
	appConfig, err := config.GetConfig()
	if err != nil {
		slog.Error("Could not get app environment", "error", err.Error())
		os.Exit(1)
	}
	if appConfig.Environment == "dev" {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
	slog.Debug("Loaded env successfully")

	// Connect to database
	err = db.Connect()
	if err != nil {
		slog.Error("Could not connect to database", "error", err.Error())
		os.Exit(1)
	}
	slog.Debug("Connected to the database")

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Could not run migrations", "error", err.Error())
		os.Exit(1)
	}
	slog.Debug("Migrations run successfully")
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
