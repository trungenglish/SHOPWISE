package main

import (
	"log/slog"
	"os"

	"shopwise/apps/server/internal/platform/config"
	"shopwise/apps/server/internal/platform/database"
	"shopwise/apps/server/internal/platform/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	log := logger.New(cfg.AppName, cfg.LogLevel)

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode)
	if err != nil {
		log.Error("failed to connect database", slog.Any("error", err))
		os.Exit(1)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Error("database close failed", slog.Any("error", err))
		}
	}()
}
