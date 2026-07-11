package handler

import (
	"net/http"

	"shopwise/retail/internal/platform/apperror"
	"shopwise/retail/internal/platform/middleware"

	"github.com/gin-gonic/gin"
)

// DeleteMe godoc
//
//	@Summary	Delete account and all user data (FR-037)
//	@Tags		users
//	@Security	BearerAuth
//	@Success	204
//	@Failure	401	{object}	ErrorResponse
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/users/me [delete]
func (h *Handler) DeleteMe(c *gin.Context) {
	userID, ok := middleware.UserID(c)
	if !ok {
		_ = c.Error(apperror.Unauthorized("authentication required", nil))
		return
	}

	if err := h.svc.DeleteAccount(c.Request.Context(), userID); err != nil {
		_ = c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
