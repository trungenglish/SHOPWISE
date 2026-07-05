package usecase

import (
	"context"
	"time"

	"shopwise/apps/server/internal/identity/domain"

	"github.com/google/uuid"
)

// TokenVerifier validates access tokens.
type TokenVerifier interface {
	Verify(token string) (subject string, err error)
}

type UserAuthRepository interface {
	CreateWithPassword(ctx context.Context, email, name, passwordHash string) (*domain.AuthUser, error)
	GetByEmail(ctx context.Context, email string) (*domain.AuthUser, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.AuthUser, error)
	MarkEmailVerified(ctx context.Context, userID uuid.UUID, verifiedAt time.Time) error
}

type AuthIdentityRepository interface {
	FindByProvider(ctx context.Context, provider, subject string) (*domain.AuthUser, error)
	Link(ctx context.Context, userID uuid.UUID, provider, subject string) error
	CreateGoogleUser(ctx context.Context, email, name, provider, subject string, verifiedAt *time.Time) (*domain.AuthUser, error)
}

type RefreshTokenRepository interface {
	Store(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error
	FindValid(ctx context.Context, tokenHash string, now time.Time) (uuid.UUID, error)
	Revoke(ctx context.Context, tokenHash string, revokedAt time.Time) error
}

type EmailVerificationRepository interface {
	Create(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error
	FindValidByHash(ctx context.Context, tokenHash string, now time.Time) (uuid.UUID, error)
	MarkUsed(ctx context.Context, tokenHash string, usedAt time.Time) error
}

type TokenPair struct {
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
}

type GoogleUserInfo struct {
	Subject       string
	Email         string
	Name          string
	EmailVerified bool
}

type GoogleTokenExchanger interface {
	Exchange(ctx context.Context, code string) (*GoogleUserInfo, error)
	AuthCodeURL(state string) string
}
