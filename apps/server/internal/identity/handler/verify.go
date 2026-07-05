package handler

import (
	"net/http"

	"shopwise/apps/server/internal/platform/apperror"

	"github.com/gin-gonic/gin"
)

type VerifyEmailRequest struct {
	Token string `json:"token" binding:"required"`
}

// VerifyEmail godoc
//
//	@Summary	Verify email with token
//	@Tags		identity
//	@Accept		json
//	@Produce	json
//	@Param		body	body	VerifyEmailRequest	true	"Verify email"
//	@Success	200
//	@Failure	400	{object}	response.ProblemDetail
//	@Router		/identity/verify-email [post]
func (h *Handler) VerifyEmail(c *gin.Context) {
	var req VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(apperror.Validation("invalid request body", err))
		return
	}
	if err := h.svc.VerifyEmail(c.Request.Context(), req.Token); err != nil {
		_ = c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}
