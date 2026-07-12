package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secret string
}

type resumeClaims struct {
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: secret}
}

// IssueResumeToken generates a single-use token valid for 24 hours
func (s *JWTService) IssueResumeToken(sessionID uuid.UUID, userID *uuid.UUID) (string, uuid.UUID, time.Time, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)
	jti := uuid.New()

	sub := sessionID.String()
	if userID != nil {
		sub = userID.String()
	}

	claims := resumeClaims{
		SessionID: sessionID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti.String(),
			Subject:   sub,
			Issuer:    "shopwise-resume-service",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return "", uuid.Nil, time.Time{}, fmt.Errorf("sign resume token: %w", err)
	}

	return signed, jti, expiresAt, nil
}

// Verify validates the token and returns the sessionID and JTI
func (s *JWTService) Verify(tokenStr string) (uuid.UUID, uuid.UUID, error) {
	parsed, err := jwt.ParseWithClaims(tokenStr, &resumeClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("verify token: %w", err)
	}
	claims, ok := parsed.Claims.(*resumeClaims)
	if !ok || !parsed.Valid {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid claims")
	}

	sessionID, err := uuid.Parse(claims.SessionID)
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid session_id format")
	}

	jti, err := uuid.Parse(claims.ID)
	if err != nil {
		return uuid.Nil, uuid.Nil, fmt.Errorf("invalid jti format")
	}

	return sessionID, jti, nil
}
