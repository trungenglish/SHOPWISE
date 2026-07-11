package handler

import (
	"net/http"

	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/platform/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) ListPreferences(c *gin.Context) {
	uid, anonID := getIdentity(c)
	if uid == nil {
		if anonID != nil {
			// Anonymous users don't have stored preferences in the DB per this design, or they could but model requires UserID
			c.JSON(http.StatusOK, gin.H{"data": []domain.UserPreference{}})
			return
		}
		c.Error(apperror.Unauthorized("unauthorized", nil))
		return
	}

	prefs, err := h.service.ListPreferences(c.Request.Context(), *uid)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prefs})
}

type UpdatePreferenceRequest struct {
	Category string `json:"category" binding:"required"`
	Value    string `json:"value" binding:"required"` // JSON encoded string
}

func (h *Handler) UpdatePreference(c *gin.Context) {
	uid, _ := getIdentity(c)
	if uid == nil {
		c.Error(apperror.Unauthorized("unauthorized", nil))
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperror.Validation("invalid preference id", err))
		return
	}

	var req UpdatePreferenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.Validation("invalid request body", err))
		return
	}

	pref := &domain.UserPreference{
		ID:       id,
		UserID:   *uid,
		Category: req.Category,
		Value:    req.Value,
	}

	if err := h.service.UpdatePreference(c.Request.Context(), pref); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, pref)
}

func (h *Handler) DeletePreference(c *gin.Context) {
	uid, _ := getIdentity(c)
	if uid == nil {
		c.Error(apperror.Unauthorized("unauthorized", nil))
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperror.Validation("invalid preference id", err))
		return
	}

	if err := h.service.DeletePreference(c.Request.Context(), id, *uid); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
