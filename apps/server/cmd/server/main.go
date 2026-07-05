// @title			Boilerplate Server API
// @version		1.0
// @description	Modular monolith REST API for React web and native clients.
// @host			localhost:18080
// @BasePath		/api/v1
// @schemes		http
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

	_ "shopwise/apps/server/docs"
	adminhandler "shopwise/apps/server/internal/administration/handler"
	adminusecase "shopwise/apps/server/internal/administration/usecase"
	fileshandler "shopwise/apps/server/internal/files/handler"
	localstorage "shopwise/apps/server/internal/files/repository/local"
	filesusecase "shopwise/apps/server/internal/files/usecase"
	identityhandler "shopwise/apps/server/internal/identity/handler"
	identitypostgres "shopwise/apps/server/internal/identity/repository/postgres"
	identityusecase "shopwise/apps/server/internal/identity/usecase"
	"shopwise/apps/server/internal/platform/cache"
	"shopwise/apps/server/internal/platform/config"
	"shopwise/apps/server/internal/platform/database"
	"shopwise/apps/server/internal/platform/health"
	"shopwise/apps/server/internal/platform/job"
	"shopwise/apps/server/internal/platform/logger"
	"shopwise/apps/server/internal/platform/middleware"
	"shopwise/apps/server/internal/platform/router"
	usershandler "shopwise/apps/server/internal/users/handler"
	userpostgres "shopwise/apps/server/internal/users/repository/postgres"
	usersusecase "shopwise/apps/server/internal/users/usecase"

	langfuse "github.com/git-hulk/langfuse-go"
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

	if err := userpostgres.Migrate(db); err != nil {
		log.Error("failed to migrate users", slog.Any("error", err))
		os.Exit(1)
	}

	if err := identitypostgres.Migrate(db); err != nil {
		log.Error("failed to migrate identity", slog.Any("error", err))
		os.Exit(1)
	}

	redisClient, err := cache.Connect(cfg.RedisURL)
	if err != nil {
		log.Error("failed to connect redis", slog.Any("error", err))
		os.Exit(1)
	}

	if err := database.EnsureStartupKey(context.Background(), db); err != nil {
		log.Error("failed to seed metadata", slog.Any("error", err))
		os.Exit(1)
	}

	jobClient, err := job.NewClient(cfg.RedisURL)
	if err != nil {
		log.Error("failed to create job client", slog.Any("error", err))
		os.Exit(1)
	}

	userRepo := userpostgres.NewRepository(db)

	identityRepo := identitypostgres.NewRepository(db)
	jwtSvc := identityusecase.NewJWTService(cfg.JWTSecret, cfg.JWTAccessTTL)
	googleSvc := identityusecase.NewGoogleOAuthService(cfg)
	identitySvc := identityusecase.NewService(
		identityRepo,
		identityRepo,
		identityRepo,
		identityRepo,
		jwtSvc,
		cfg.JWTRefreshTTL,
		googleSvc,
	).WithDevLogging(cfg.GinMode == "debug", log)
	identityH := identityhandler.NewHandler(identitySvc)

	fileStorage, err := localstorage.NewStorage(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init file storage", slog.Any("error", err))
		os.Exit(1)
	}

	useStubLLM := cfg.LLMAPIKey == ""
	if !useStubLLM {
		var lfClient *langfuse.Langfuse
		if cfg.LangfusePublicKey != "" && cfg.LangfuseSecretKey != "" {
			lfClient = langfuse.NewClient(cfg.LangfuseHost, cfg.LangfusePublicKey, cfg.LangfuseSecretKey)
			defer lfClient.Close()
		}
	}

	guestRateLimit := middleware.NewIPRateLimiter(30, time.Minute).Middleware()

	userSvc := usersusecase.NewService(userRepo, jobClient, log).WithAccountDeletion(usersusecase.AccountDeletionDeps{
		DB:       db,
		Identity: identityRepo,
		Users:    userRepo,
		Storage:  fileStorage,
	})
	userH := usershandler.NewHandler(userSvc)

	engine := router.New(router.Dependencies{
		Config:         cfg,
		Log:            log,
		TokenVerifier:  jwtSvc,
		EmailChecker:   identityRepo,
		GuestRateLimit: guestRateLimit,
		Modules: router.Modules{
			Health:         health.NewHandler(db, redisClient),
			Identity:       identityH,
			Users:          userH,
			Files:          fileshandler.NewHandler(filesusecase.NewService()),
			Administration: adminhandler.NewHandler(adminusecase.NewService(log)),
		},
	})

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	logger.SysStartup(log, cfg.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server stopped unexpectedly", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.SysShutdown(log)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown failed", slog.Any("error", err))
	}

	if err := database.Close(db); err != nil {
		log.Error("database close failed", slog.Any("error", err))
	}

	if err := redisClient.Close(); err != nil {
		log.Error("redis close failed", slog.Any("error", err))
	}

	if err := jobClient.Close(); err != nil {
		log.Error("job client close failed", slog.Any("error", err))
	}
}
