package middleware

import (
	"strings"

	identityusecase "shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const userIDKey = "userID"

type TokenVerifier = identityusecase.TokenVerifier

func Auth(verifier TokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			_ = c.Error(apperror.Unauthorized("missing authorization header", nil))
			c.Abort()
			return
		}
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
			_ = c.Error(apperror.Unauthorized("invalid authorization header", nil))
			c.Abort()
			return
		}
		subject, err := verifier.Verify(parts[1])
		if err != nil {
			_ = c.Error(apperror.Unauthorized("invalid access token", err))
			c.Abort()
			return
		}
		userID, err := uuid.Parse(subject)
		if err != nil {
			_ = c.Error(apperror.Unauthorized("invalid access token subject", err))
			c.Abort()
			return
		}
		c.Set(userIDKey, userID)
		c.Next()
	}
}

func UserID(c *gin.Context) (uuid.UUID, bool) {
	value, ok := c.Get(userIDKey)
	if !ok {
		return uuid.Nil, false
	}
	userID, ok := value.(uuid.UUID)
	return userID, ok
}
