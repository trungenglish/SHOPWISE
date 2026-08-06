package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shopwise/retail/internal/accessories"
	"shopwise/retail/internal/accessories/phongvu"
	"shopwise/retail/internal/platform/cache"
	"shopwise/retail/internal/platform/config"
	"shopwise/retail/internal/platform/database"
	"shopwise/retail/internal/platform/job"
	"shopwise/retail/internal/platform/logger"

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
	mux.HandleFunc(job.TypeOrderConfirmation, func(_ context.Context, task *asynq.Task) error {
		return job.HandleOrderConfirmation(cfg)(task)
	})

	scheduler := asynq.NewScheduler(opts, nil)
	if _, err := scheduler.Register("@every 30m", asynq.NewTask(job.TypePriceCheck, nil)); err != nil {
		log.Error("failed to register price check schedule", slog.Any("error", err))
		os.Exit(1)
	}
	if cfg.PhongVuConnectorEnabled {
		accessoryRepository := accessories.NewRepository(db)
		catalogClient := phongvu.NewClient(http.DefaultClient, 10*time.Second)
		fetcher := &phongVuFetcher{client: catalogClient}
		syncer := accessories.NewSyncer(fetcher, accessoryRepository, phongVuSources, time.Now)
		mux.HandleFunc(job.TypePhongVuAccessorySync, func(ctx context.Context, _ *asynq.Task) error {
			return syncer.Sync(ctx)
		})
		if _, err := scheduler.Register("@every 6h", asynq.NewTask(job.TypePhongVuAccessorySync, nil)); err != nil {
			log.Error("failed to register Phong Vu sync schedule", slog.Any("error", err))
			os.Exit(1)
		}
		initialClient := asynq.NewClient(opts)
		_, enqueueErr := initialClient.Enqueue(asynq.NewTask(job.TypePhongVuAccessorySync, nil), asynq.Unique(6*time.Hour))
		_ = initialClient.Close()
		if enqueueErr != nil && !errors.Is(enqueueErr, asynq.ErrDuplicateTask) {
			log.Error("failed to enqueue initial Phong Vu sync", slog.Any("error", enqueueErr))
		}
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

var phongVuSources = []accessories.CatalogSource{
	{URL: "https://phongvu.vn/c/chuot-co-day", Category: "mouse"},
	{URL: "https://phongvu.vn/c/ban-phim", Category: "keyboard"},
	{URL: "https://phongvu.vn/c/tai-nghe", Category: "headset"},
	{URL: "https://phongvu.vn/c/de-tan-nhiet", Category: "cooling_pad"},
	{URL: "https://phongvu.vn/c/man-hinh-may-tinh", Category: "monitor"},
	{URL: "https://phongvu.vn/c/hub-usb", Category: "hub"},
	{URL: "https://phongvu.vn/c/webcam", Category: "webcam"},
	{URL: "https://phongvu.vn/c/balo-laptop", Category: "bag"},
	{URL: "https://phongvu.vn/c/o-cung-di-dong", Category: "external_ssd"},
}

type phongVuFetcher struct{ client *phongvu.Client }

func (fetcher *phongVuFetcher) Fetch(ctx context.Context, source accessories.CatalogSource) ([]accessories.SyncedProduct, error) {
	products, err := fetcher.client.FetchCategory(ctx, source.URL, source.Category)
	if err != nil {
		return nil, err
	}
	result := make([]accessories.SyncedProduct, 0, len(products))
	for _, product := range products {
		result = append(result, accessories.SyncedProduct{
			RetailerProductID: product.RetailerProductID, Name: product.Name, Brand: product.Brand,
			Category: product.Category, ImageURL: product.ImageURL, SourceURL: product.SourceURL,
			Price: product.Price, OriginalPrice: product.OriginalPrice, InStock: product.InStock,
		})
	}
	return result, nil
}
