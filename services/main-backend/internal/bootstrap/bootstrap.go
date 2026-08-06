package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"shopwise/retail/internal/accessories"
	"shopwise/retail/internal/accessories/phongvu"
	adminhandler "shopwise/retail/internal/administration/handler"
	adminusecase "shopwise/retail/internal/administration/usecase"
	cataloghandler "shopwise/retail/internal/catalog/handler"
	decisionmemoryhandler "shopwise/retail/internal/decision_memory/handler"
	decisionmemorypostgres "shopwise/retail/internal/decision_memory/repository/postgres"
	decisionmemoryusecase "shopwise/retail/internal/decision_memory/usecase"
	fileshandler "shopwise/retail/internal/files/handler"
	localstorage "shopwise/retail/internal/files/repository/local"
	filesusecase "shopwise/retail/internal/files/usecase"
	identityhandler "shopwise/retail/internal/identity/handler"
	identitypostgres "shopwise/retail/internal/identity/repository/postgres"
	identityusecase "shopwise/retail/internal/identity/usecase"
	ordershandler "shopwise/retail/internal/orders/handler"
	orderspostgres "shopwise/retail/internal/orders/repository/postgres"
	ordersusecase "shopwise/retail/internal/orders/usecase"
	"shopwise/retail/internal/platform/cache"
	"shopwise/retail/internal/platform/config"
	"shopwise/retail/internal/platform/database"
	"shopwise/retail/internal/platform/health"
	"shopwise/retail/internal/platform/job"
	"shopwise/retail/internal/platform/logger"
	"shopwise/retail/internal/platform/middleware"
	"shopwise/retail/internal/platform/router"
	resumesessionhandler "shopwise/retail/internal/resume_session/handler"
	zaloprovider "shopwise/retail/internal/resume_session/provider/zalo"
	resumesessionpostgres "shopwise/retail/internal/resume_session/repository/postgres"
	resumesessionusecase "shopwise/retail/internal/resume_session/usecase"
	storeshandler "shopwise/retail/internal/stores/handler"
	usershandler "shopwise/retail/internal/users/handler"
	userpostgres "shopwise/retail/internal/users/repository/postgres"
	usersusecase "shopwise/retail/internal/users/usecase"
	"shopwise/retail/internal/promotions"
	promotionsrepo "shopwise/retail/internal/promotions/repository"
)

// Run initializes and starts the HTTP server.
func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	log := logger.New(cfg.AppName, cfg.LogLevel)

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	redisClient, err := cache.Connect(cfg.RedisURL)
	if err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}

	jobClient, err := job.NewClient(cfg.RedisURL)
	if err != nil {
		return fmt.Errorf("failed to create job client: %w", err)
	}

	userRepo := userpostgres.NewRepository(db)
	identityRepo := identitypostgres.NewRepository(db)
	orderRepo := orderspostgres.NewRepository(db)
	decisionRepo := decisionmemorypostgres.NewRepository(db)
	accessoryRepo := accessories.NewRepository(db)

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
		return fmt.Errorf("failed to init file storage: %w", err)
	}

	guestRateLimit := middleware.NewIPRateLimiter(30, time.Minute).Middleware()

	userSvc := usersusecase.NewService(userRepo, jobClient, log).WithAccountDeletion(usersusecase.AccountDeletionDeps{
		DB:       db,
		Identity: identityRepo,
		Users:    userRepo,
		Storage:  fileStorage,
	})
	userH := usershandler.NewHandler(userSvc)
	phongVuClient := phongvu.NewClient(http.DefaultClient, 10*time.Second)
	orderService := ordersusecase.NewService(orderRepo, orderRepo, orderRepo, orderRepo).
		WithRetailerOffers(orderRepo, phongvu.NewOfferRefresher(phongVuClient, cfg.PhongVuConnectorEnabled, time.Now), orderRepo).
		WithOrderConfirmation(jobClient)
	orderH := ordershandler.NewHandler(orderService)
	if cfg.CheckoutAuthBypass {
		orderH.WithDevelopmentAuthBypass()
	}

	promotionsRepo := promotionsrepo.NewPostgresRepository(db)
	voucherGen := promotions.NewLocalVoucherGenerator()
	promotionsSvc := promotions.NewService(promotionsRepo, voucherGen)
	promotionsH := promotions.NewHandler(promotionsSvc)

	orderH.WithPromotionSink(promotionsSvc)

	healthH := health.NewHandler(db, redisClient)
	filesH := fileshandler.NewHandler(filesusecase.NewService())
	adminH := adminhandler.NewHandler(adminusecase.NewService(log))

	decisionSvc := decisionmemoryusecase.NewService(decisionRepo)
	decisionH := decisionmemoryhandler.NewHandler(decisionSvc, cfg.AIRuntimeURL)
	accessoryH := accessories.NewHandler(accessories.NewService(accessoryRepo, time.Now))

	resumeRepo := resumesessionpostgres.NewRepository(db)
	resumeJWTSvc := resumesessionusecase.NewJWTService(cfg.JWTSecret)

	zaloProv := zaloprovider.NewProvider(cfg.ZaloOAID, cfg.ZaloAPIToken)
	retryZalo := resumesessionusecase.NewRetryableProvider(zaloProv, 1)

	resumeSvc := resumesessionusecase.NewService(resumeRepo, resumeJWTSvc, decisionSvc, userSvc, retryZalo)
	resumeH := resumesessionhandler.NewHandler(resumeSvc)

	engine := router.New(router.Dependencies{
		Config: cfg,
		Log:    log,
	})

	// Register Routes
	engine.GET("/health", healthH.Health)

	v1 := engine.Group("/api/v1")

	identityGroup := v1.Group("/identity")
	identityhandler.RegisterRoutes(identityGroup, identityH)
	identityhandler.RegisterGoogleRoutes(identityGroup, identityH, cfg.WebAppURL)

	usershandler.RegisterRoutes(v1.Group("/users"), userH, jwtSvc)
	ordershandler.RegisterRoutes(v1.Group("/checkout"), orderH, jwtSvc)
	ordershandler.RegisterListRoutes(v1.Group("/orders"), orderH, jwtSvc)
	fileshandler.RegisterRoutes(v1.Group("/files"), filesH)

	catalogH := cataloghandler.NewHandler(db)
	cataloghandler.RegisterRoutes(v1.Group("/products"), catalogH)
	accessories.RegisterRoutes(v1.Group("/accessories"), accessoryH)

	storesH := storeshandler.NewHandler()
	storeshandler.RegisterRoutes(v1.Group("/stores"), storesH)

	decisionmemoryhandler.RegisterRoutes(v1.Group("/sessions"), decisionH, jwtSvc)
	resumesessionhandler.RegisterRoutes(v1.Group("/session"), resumeH)

	promotionsH.RegisterRoutes(v1)
	if cfg.GinMode == "debug" {
		promotionsH.RegisterDevRoutes(v1)
	}

	adminGroup := v1.Group("/admin")
	adminGroup.Use(guestRateLimit)
	adminhandler.RegisterRoutes(adminGroup, adminH)

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

	return nil
}

// Migrate executes database migrations.
func Migrate() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	log := logger.New(cfg.AppName, cfg.LogLevel)

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer database.Close(db)

	log.Info("running user migrations")
	if err := userpostgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate users: %w", err)
	}

	log.Info("running identity migrations")
	if err := identitypostgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate identity: %w", err)
	}

	log.Info("running order migrations")
	if err := orderspostgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate orders: %w", err)
	}

	log.Info("running decision memory migrations")
	if err := decisionmemorypostgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate decision memory: %w", err)
	}

	log.Info("running accessory catalog migrations")
	if err := accessories.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate accessory catalog: %w", err)
	}

	log.Info("running resume session migrations")
	if err := resumesessionpostgres.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate resume session: %w", err)
	}

	log.Info("running promotions migrations")
	if err := promotionsrepo.Migrate(db); err != nil {
		return fmt.Errorf("failed to migrate promotions: %w", err)
	}

	log.Info("migrations completed successfully")
	return nil
}

// Seed seeds the database with initial metadata.
func Seed() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	log := logger.New(cfg.AppName, cfg.LogLevel)

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer database.Close(db)

	log.Info("seeding database metadata")
	if err := database.EnsureStartupKey(context.Background(), db); err != nil {
		return fmt.Errorf("failed to seed metadata: %w", err)
	}

	log.Info("seeding completed successfully")
	return nil
}
