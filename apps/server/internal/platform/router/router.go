package router

import (
	"log/slog"

	adminhandler "shopwise/apps/server/internal/administration/handler"
	fileshandler "shopwise/apps/server/internal/files/handler"
	identityhandler "shopwise/apps/server/internal/identity/handler"
	"shopwise/apps/server/internal/platform/config"
	"shopwise/apps/server/internal/platform/health"
	"shopwise/apps/server/internal/platform/middleware"
	usershandler "shopwise/apps/server/internal/users/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Modules struct {
	Health         *health.Handler
	Identity       *identityhandler.Handler
	Users          *usershandler.Handler
	Files          *fileshandler.Handler
	Administration *adminhandler.Handler
}

type Dependencies struct {
	Config         *config.Config
	Log            *slog.Logger
	TokenVerifier  middleware.TokenVerifier
	EmailChecker   middleware.EmailVerifiedChecker
	GuestRateLimit gin.HandlerFunc
	Modules        Modules
}

func New(deps Dependencies) *gin.Engine {
	gin.SetMode(deps.Config.GinMode)

	debugMode := deps.Config.GinMode == gin.DebugMode

	engine := gin.New()
	engine.Use(middleware.Recovery(deps.Log, debugMode))
	engine.Use(middleware.RequestLogger(deps.Log))
	engine.Use(middleware.CORS(deps.Config.AllowedOrigins))
	engine.Use(middleware.ErrorHandler(deps.Log, debugMode))

	engine.GET("/health", deps.Modules.Health.Health)

	if debugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1 := engine.Group("/api/v1")
	identityGroup := v1.Group("/identity")
	identityhandler.RegisterRoutes(identityGroup, deps.Modules.Identity)
	identityhandler.RegisterGoogleRoutes(identityGroup, deps.Modules.Identity, deps.Config.WebAppURL)
	usershandler.RegisterRoutes(v1.Group("/users"), deps.Modules.Users, deps.TokenVerifier)
	fileshandler.RegisterRoutes(v1.Group("/files"), deps.Modules.Files)
	adminhandler.RegisterRoutes(v1.Group("/admin"), deps.Modules.Administration)

	return engine
}
