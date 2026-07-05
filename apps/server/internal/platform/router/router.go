package router

import (
	"log/slog"

	"shopwise/apps/server/internal/platform/config"
	"shopwise/apps/server/internal/platform/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Dependencies struct {
	Config *config.Config
	Log    *slog.Logger
}

// New creates a new Gin engine with global middlewares and basic routes.
func New(deps Dependencies) *gin.Engine {
	gin.SetMode(deps.Config.GinMode)

	debugMode := deps.Config.GinMode == gin.DebugMode

	engine := gin.New()
	engine.Use(middleware.Recovery(deps.Log, debugMode))
	engine.Use(middleware.RequestLogger(deps.Log))
	engine.Use(middleware.CORS(deps.Config.AllowedOrigins))
	engine.Use(middleware.ErrorHandler(deps.Log, debugMode))

	if debugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return engine
}
