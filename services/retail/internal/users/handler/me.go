package handler

import (
	"net/http"
	"time"

	"shopwise/retail/internal/platform/apperror"
	"shopwise/retail/internal/platform/middleware"
	"shopwise/retail/internal/users/domain"
	"shopwise/retail/internal/users/usecase"

	"github.com/gin-gonic/gin"
)

type UserPreferencesResponse struct {
	BudgetSensitivity   string   `json:"budgetSensitivity"`
	PreferredCategories []string `json:"preferredCategories"`
	BrandOpenness       string   `json:"brandOpenness"`
	DefaultCurrency     string   `json:"defaultCurrency"`
}

type MeResponse struct {
	ID              string                  `json:"id"`
	Email           string                  `json:"email"`
	Name            string                  `json:"name"`
	EmailVerifiedAt *time.Time              `json:"emailVerifiedAt"`
	Preferences     UserPreferencesResponse `json:"preferences"`
}

type UpdateMeRequest struct {
	Name                *string   `json:"name,omitempty"`
	BudgetSensitivity   *string   `json:"budgetSensitivity,omitempty"`
	PreferredCategories *[]string `json:"preferredCategories,omitempty"`
	BrandOpenness       *string   `json:"brandOpenness,omitempty"`
	DefaultCurrency     *string   `json:"defaultCurrency,omitempty"`
}

// GetMe godoc
//
//	@Summary	Get current user profile
//	@Tags		users
//	@Security	BearerAuth
//	@Produce	json
//	@Success	200	{object}	MeResponse
//	@Failure	401	{object}	ErrorResponse
//	@Router		/users/me [get]
func (h *Handler) GetMe(c *gin.Context) {
	userID, ok := middleware.UserID(c)
	if !ok {
		_ = c.Error(apperror.Unauthorized("authentication required", nil))
		return
	}
	profile, err := h.svc.GetProfile(c.Request.Context(), userID)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, toMeResponse(profile))
}

// PatchMe godoc
//
//	@Summary	Update profile and preferences
//	@Tags		users
//	@Security	BearerAuth
//	@Accept		json
//	@Produce	json
//	@Param		body	body		UpdateMeRequest	true	"Profile update"
//	@Success	200		{object}	MeResponse
//	@Failure	401		{object}	ErrorResponse
//	@Router		/users/me [patch]
func (h *Handler) PatchMe(c *gin.Context) {
	userID, ok := middleware.UserID(c)
	if !ok {
		_ = c.Error(apperror.Unauthorized("authentication required", nil))
		return
	}
	var req UpdateMeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}
	profile, err := h.svc.UpdateProfile(c.Request.Context(), userID, usecase.UpdateProfileInput{
		Name:                req.Name,
		BudgetSensitivity:   req.BudgetSensitivity,
		PreferredCategories: req.PreferredCategories,
		BrandOpenness:       req.BrandOpenness,
		DefaultCurrency:     req.DefaultCurrency,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, toMeResponse(profile))
}

func toMeResponse(profile *domain.UserProfile) MeResponse {
	return MeResponse{
		ID:              profile.ID.String(),
		Email:           profile.Email,
		Name:            profile.Name,
		EmailVerifiedAt: profile.EmailVerifiedAt,
		Preferences: UserPreferencesResponse{
			BudgetSensitivity:   profile.Preferences.BudgetSensitivity,
			PreferredCategories: profile.Preferences.PreferredCategories,
			BrandOpenness:       profile.Preferences.BrandOpenness,
			DefaultCurrency:     profile.Preferences.DefaultCurrency,
		},
	}
}
