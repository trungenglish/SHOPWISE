package promotions

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/checkout-incentive")
	{
		group.POST("/start", h.StartPromotion)
		group.GET("/status", h.GetStatus)
	}
}

type startPromotionRequest struct {
	TriggerType TriggerType `json:"trigger_type" binding:"required"`
}

func (h *Handler) StartPromotion(c *gin.Context) {
	var req startPromotionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-Session-ID header"})
		return
	}

	var userID *uuid.UUID
	if uidStr := c.GetHeader("X-User-ID"); uidStr != "" {
		uid, err := uuid.Parse(uidStr)
		if err == nil {
			userID = &uid
		}
	}

	var deviceID *string
	if did := c.GetHeader("X-Device-ID"); did != "" {
		deviceID = &did
	}

	if userID == nil && deviceID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must provide either X-User-ID or X-Device-ID header"})
		return
	}

	promo, err := h.service.StartPromotion(c.Request.Context(), sessionID, userID, deviceID, req.TriggerType)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"promotion_id": promo.PromotionID.String(),
		"status":       promo.Status,
		"expires_at":   promo.ExpiresAt.Format(time.RFC3339),
		"reward_type":  promo.RewardType,
		"reward_value": promo.RewardValue,
	})
}

func (h *Handler) GetStatus(c *gin.Context) {
	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing X-Session-ID header"})
		return
	}

	promo, err := h.service.GetStatus(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if promo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active promotion found for this session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        promo.Status,
		"expires_at":    promo.ExpiresAt.Format(time.RFC3339),
		"campaign_code": promo.CampaignCode,
		"reward_type":   promo.RewardType,
		"reward_value":  promo.RewardValue,
		"reward_scope":  promo.RewardScope,
	})
}
