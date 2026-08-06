package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"

	"shopwise/retail/internal/resume_session/usecase"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *usecase.Service
}

func NewHandler(svc *usecase.Service) *Handler {
	return &Handler{svc: svc}
}

// ResumeSession Handles GET /api/v1/session/resume
func (h *Handler) ResumeSession(c *gin.Context) {
	tokenStr := c.Query("token")
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token_missing", "message": "Resume token is required."})
		return
	}

	ip := c.ClientIP()
	ua := c.Request.UserAgent()
	hashBytes := sha256.Sum256([]byte(ip + ua))
	hashStr := hex.EncodeToString(hashBytes[:])

	session, err := h.svc.ResumeSession(c.Request.Context(), tokenStr, &hashStr)
	if err != nil {
		if errors.Is(err, usecase.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token_expired", "message": "This resume link has expired. Please request a new one."})
			return
		}
		if errors.Is(err, usecase.ErrTokenConsumed) {
			c.JSON(http.StatusConflict, gin.H{"error": "token_consumed", "message": "This resume link has already been used."})
			return
		}
		if errors.Is(err, usecase.ErrTokenRevoked) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token_revoked", "message": "This resume link was revoked because a newer one was generated."})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token_invalid", "message": "Invalid resume link."})
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *Handler) NotifyInactivity(c *gin.Context) {
	sessionIDStr := c.Param("id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid_session_id"})
		return
	}

	err = h.svc.TriggerInactivityNotification(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification triggered (mocked Zalo message sent)"})
}

func RegisterRoutes(r gin.IRouter, h *Handler) {
	r.GET("/resume", h.ResumeSession)
	r.POST("/:id/notify-inactivity", h.NotifyInactivity)
}
