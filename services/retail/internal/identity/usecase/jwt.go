package usecase

import (
	"errors"
	"fmt"
	"time"

	"shopwise/retail/internal/identity/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secret    []byte
	accessTTL time.Duration
}

type accessClaims struct {
	jwt.RegisteredClaims
}

func NewJWTService(secret string, accessTTL time.Duration) *JWTService {
	return &JWTService{
		secret:    []byte(secret),
		accessTTL: accessTTL,
	}
}

func (s *JWTService) IssueAccessToken(userID uuid.UUID) (string, int, error) {
	now := time.Now().UTC()
	expiresAt := now.Add(s.accessTTL)
	claims := accessClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", 0, fmt.Errorf("sign access token: %w", err)
	}
	return signed, int(s.accessTTL.Seconds()), nil
}

func (s *JWTService) Verify(token string) (string, error) {
	parsed, err := jwt.ParseWithClaims(token, &accessClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.secret, nil
	})
	if err != nil {
		return "", fmt.Errorf("%w: %w", domain.ErrInvalidCredentials, err)
	}
	claims, ok := parsed.Claims.(*accessClaims)
	if !ok || !parsed.Valid {
		return "", domain.ErrInvalidCredentials
	}
	if claims.Subject == "" {
		return "", domain.ErrInvalidCredentials
	}
	return claims.Subject, nil
}

func ParseUserID(subject string) (uuid.UUID, error) {
	id, err := uuid.Parse(subject)
	if err != nil {
		return uuid.Nil, domain.ErrInvalidCredentials
	}
	return id, nil
}

func IsTokenExpired(err error) bool {
	return errors.Is(err, jwt.ErrTokenExpired)
}
