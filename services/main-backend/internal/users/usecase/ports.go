package usecase

import (
	"context"

	"shopwise/retail/internal/users/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	List(ctx context.Context, limit, offset int) ([]domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetProfile(ctx context.Context, userID uuid.UUID) (*domain.UserProfile, error)
	UpdateProfileName(ctx context.Context, userID uuid.UUID, name string) error
	UpsertPreferences(ctx context.Context, prefs domain.UserPreferences) error
}

type WelcomeEmailEnqueuer interface {
	EnqueueWelcomeEmail(ctx context.Context, userID, email string) error
}
