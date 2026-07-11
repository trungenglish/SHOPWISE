package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"shopwise/retail/internal/platform/apperror"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type GenerateResumeTokenRequest struct {
	ExpiresInSeconds int `json:"expires_in_seconds"`
}

type GenerateResumeTokenResponse struct {
	Token string `json:"token"`
}

type ResumableSessionClaims struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id,omitempty"`
	AnonID    string `json:"anon_id,omitempty"`
	jwt.RegisteredClaims
}

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("fallback-secret-for-development")
	}
	return []byte(secret)
}

func (h *Handler) GenerateResumeToken(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperror.Validation("invalid session id", err))
		return
	}

	var req GenerateResumeTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperror.Validation("invalid request body", err))
		return
	}

	// Verify ownership logic (simplified for stub, normally we verify session belongs to user)
	uid, anonID := getIdentity(c)
	if uid == nil && anonID == nil {
		c.Error(apperror.Unauthorized("unauthorized", nil))
		return
	}

	// Check if session belongs to user
	session, err := h.service.GetSession(c.Request.Context(), id)
	if err != nil {
		c.Error(apperror.NotFound("session not found", err))
		return
	}

	var u, a string
	if uid != nil {
		if session.UserID == nil || *session.UserID != *uid {
			c.Error(apperror.Unauthorized("unauthorized session access", nil))
			return
		}
		u = uid.String()
	} else if anonID != nil {
		if session.AnonymousID == nil || *session.AnonymousID != *anonID {
			c.Error(apperror.Unauthorized("unauthorized session access", nil))
			return
		}
		a = *anonID
	}

	exp := req.ExpiresInSeconds
	if exp <= 0 {
		exp = 3600 // default 1 hour
	}

	claims := ResumableSessionClaims{
		SessionID: id.String(),
		UserID:    u,
		AnonID:    a,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Duration(exp) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "decision-memory-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		c.Error(fmt.Errorf("failed to sign token: %w", err))
		return
	}

	c.JSON(http.StatusOK, GenerateResumeTokenResponse{Token: tokenString})
}

func (h *Handler) ResolveResumeToken(c *gin.Context) {
	tokenString := c.Param("token")
	if tokenString == "" {
		c.Error(apperror.Validation("missing token", nil))
		return
	}

	token, err := jwt.ParseWithClaims(tokenString, &ResumableSessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})

	if err != nil {
		c.Error(apperror.Unauthorized("invalid or expired token", err))
		return
	}

	if claims, ok := token.Claims.(*ResumableSessionClaims); ok && token.Valid {
		// Token is valid. We return the session ID.
		c.JSON(http.StatusOK, gin.H{
			"session_id": claims.SessionID,
			"user_id":    claims.UserID,
			"anon_id":    claims.AnonID,
		})
		return
	}

	c.Error(apperror.Unauthorized("invalid token claims", nil))
}
