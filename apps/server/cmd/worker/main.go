package main

import (
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"shopwise/apps/server/internal/platform/cache"
	"shopwise/apps/server/internal/platform/config"
	"shopwise/apps/server/internal/platform/database"
	"shopwise/apps/server/internal/platform/job"
	"shopwise/apps/server/internal/platform/logger"

	"github.com/hibiken/asynq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	log := logger.New(cfg.AppName+"-worker", cfg.LogLevel)

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode)
	if err != nil {
		log.Error("failed to connect database", slog.Any("error", err))
		os.Exit(1)
	}

	redisClient, err := cache.Connect(cfg.RedisURL)
	if err != nil {
		log.Error("failed to connect redis", slog.Any("error", err))
		os.Exit(1)
	}

	opts, err := asynq.ParseRedisURI(cfg.RedisURL)
	if err != nil {
		log.Error("failed to parse redis url", slog.Any("error", err))
		os.Exit(1)
	}

	server := asynq.NewServer(opts, asynq.Config{
		Concurrency: 5,
		Queues: map[string]int{
			"default": 5,
		},
	})

	mux := asynq.NewServeMux()

	scheduler := asynq.NewScheduler(opts, nil)
	if _, err := scheduler.Register("@every 30m", asynq.NewTask(job.TypePriceCheck, nil)); err != nil {
		log.Error("failed to register price check schedule", slog.Any("error", err))
		os.Exit(1)
	}

	go func() {
		log.Info("worker starting", slog.String("event_type", "worker_startup"))
		if err := server.Run(mux); err != nil && !errors.Is(err, asynq.ErrServerClosed) {
			log.Error("worker stopped unexpectedly", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	go func() {
		if err := scheduler.Run(); err != nil && !errors.Is(err, asynq.ErrServerClosed) {
			log.Error("scheduler stopped unexpectedly", slog.Any("error", err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Info("worker shutting down", slog.String("event_type", "worker_shutdown"))
	scheduler.Shutdown()
	server.Shutdown()

	if err := database.Close(db); err != nil {
		log.Error("database close failed", slog.Any("error", err))
	}
	if err := redisClient.Close(); err != nil {
		log.Error("redis close failed", slog.Any("error", err))
	}
}
