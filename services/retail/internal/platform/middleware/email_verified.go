package middleware

import (
	"context"

	"shopwise/retail/internal/identity/domain"
	identityusecase "shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmailVerifiedChecker interface {
	GetByID(ctx context.Context, id uuid.UUID) (*domain.AuthUser, error)
}

func RequireEmailVerified(users EmailVerifiedChecker) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := UserID(c)
		if !ok {
			_ = c.Error(apperror.Unauthorized("authentication required", nil))
			c.Abort()
			return
		}
		user, err := users.GetByID(c.Request.Context(), userID)
		if err != nil {
			_ = c.Error(apperror.Unauthorized("authentication required", err))
			c.Abort()
			return
		}
		if user.EmailVerifiedAt == nil {
			_ = c.Error(apperror.Forbidden("email verification required", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}

// Ensure EmailVerifiedChecker is satisfied by identity usecase port.
var _ EmailVerifiedChecker = (identityusecase.UserAuthRepository)(nil)
