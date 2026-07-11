package health

import (
	"context"
	"net/http"
	"time"

	"shopwise/retail/internal/platform/cache"
	"shopwise/retail/internal/platform/database"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewHandler(db *gorm.DB, redisClient *redis.Client) *Handler {
	return &Handler{db: db, redis: redisClient}
}

type Response struct {
	Status string `json:"status"`
	DB     string `json:"db"`
	Redis  string `json:"redis"`
}

// Health godoc
//
//	@Summary	Health check
//	@Tags		platform
//	@Produce	json
//	@Success	200	{object}	Response
//	@Failure	503	{object}	Response
//	@Router		/health [get]
func (h *Handler) Health(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	dbStatus := "ok"
	if err := database.Ping(h.db); err != nil {
		dbStatus = "error"
	}

	redisStatus := "ok"
	if err := cache.Ping(ctx, h.redis); err != nil {
		redisStatus = "error"
	}

	status := http.StatusOK
	overall := "ok"
	if dbStatus != "ok" || redisStatus != "ok" {
		status = http.StatusServiceUnavailable
		overall = "degraded"
	}

	c.JSON(status, Response{
		Status: overall,
		DB:     dbStatus,
		Redis:  redisStatus,
	})
}
