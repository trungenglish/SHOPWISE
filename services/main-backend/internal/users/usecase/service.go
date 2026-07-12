package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"shopwise/retail/internal/platform/apperror"
	"shopwise/retail/internal/users/domain"

	"github.com/google/uuid"
)

const (
	defaultListLimit = 20
	maxListLimit     = 100
)

type Service struct {
	repo            UserRepository
	enqueuer        WelcomeEmailEnqueuer
	log             *slog.Logger
	accountDeletion *AccountDeletionDeps
}

func NewService(repo UserRepository, enqueuer WelcomeEmailEnqueuer, log *slog.Logger) *Service {
	return &Service{
		repo:     repo,
		enqueuer: enqueuer,
		log:      log.With("component", "users-service"),
	}
}

type CreateInput struct {
	Email string
	Name  string
	Phone string
}

type UpdateInput struct {
	Email *string
	Name  *string
	Phone *string
}

func (s *Service) Create(ctx context.Context, input CreateInput) (*domain.User, error) {
	email := strings.TrimSpace(input.Email)
	name := strings.TrimSpace(input.Name)
	phone := strings.TrimSpace(input.Phone)
	if email == "" || name == "" {
		return nil, apperror.Validation("email and name are required", nil)
	}

	user := &domain.User{
		Email: email,
		Name:  name,
		Phone: phone,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			return nil, apperror.Conflict("email already registered", err)
		}
		return nil, apperror.Internal("failed to create user", err)
	}

	if s.enqueuer != nil {
		if err := s.enqueuer.EnqueueWelcomeEmail(ctx, user.ID.String(), user.Email); err != nil {
			s.log.Error("failed to enqueue welcome email",
				slog.String("user_id", user.ID.String()),
				slog.Any("error", err),
			)
		}
	}

	return user, nil
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, apperror.NotFound("user not found", err)
		}
		return nil, apperror.Internal("failed to get user", err)
	}
	return user, nil
}

func (s *Service) GetVerifiedNotificationIdentity(ctx context.Context, userID uuid.UUID) (string, error) {
	user, err := s.GetByID(ctx, userID)
	if err != nil {
		return "", err
	}
	if user.Phone == "" {
		return "", errors.New("user has no verified phone number")
	}
	return user.Phone, nil
}

func (s *Service) List(ctx context.Context, limit, offset int) ([]domain.User, error) {
	if limit <= 0 {
		limit = defaultListLimit
	}
	if limit > maxListLimit {
		limit = maxListLimit
	}
	if offset < 0 {
		offset = 0
	}

	users, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, apperror.Internal("failed to list users", err)
	}
	return users, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateInput) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, apperror.NotFound("user not found", err)
		}
		return nil, apperror.Internal("failed to get user", err)
	}

	if input.Email != nil {
		email := strings.TrimSpace(*input.Email)
		if email == "" {
			return nil, apperror.Validation("email cannot be empty", nil)
		}
		user.Email = email
	}
	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		if name == "" {
			return nil, apperror.Validation("name cannot be empty", nil)
		}
		user.Name = name
	}
	if input.Phone != nil {
		phone := strings.TrimSpace(*input.Phone)
		user.Phone = phone
	}

	if err := s.repo.Update(ctx, user); err != nil {
		if errors.Is(err, domain.ErrDuplicateEmail) {
			return nil, apperror.Conflict("email already registered", err)
		}
		if errors.Is(err, domain.ErrNotFound) {
			return nil, apperror.NotFound("user not found", err)
		}
		return nil, apperror.Internal("failed to update user", err)
	}

	updated, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, apperror.Internal("failed to reload user", fmt.Errorf("get after update: %w", err))
	}
	return updated, nil
}

func (s *Service) UpsertGuestUser(ctx context.Context, input CreateInput) (*domain.User, error) {
	// First attempt to create
	user, err := s.Create(ctx, input)
	if err != nil {
		var appErr *apperror.AppError
		isConflict := errors.As(err, &appErr) && appErr.Code == "CONFLICT"
		if isConflict || strings.Contains(err.Error(), "already registered") {
			// If conflict, find the user and update name and phone
			// Since we don't have GetByEmail in this simplified domain, we can only list and filter or rely on repo.
			// Wait, the repository might not have GetByEmail exposed, so let's find the user.
			users, listErr := s.repo.List(ctx, 1000, 0)
			if listErr != nil {
				return nil, apperror.Internal("failed to list users for upsert", listErr)
			}
			var existing *domain.User
			for i := range users {
				if users[i].Email == strings.TrimSpace(input.Email) {
					existing = &users[i]
					break
				}
			}
			if existing != nil {
				// Update existing
				existing.Name = strings.TrimSpace(input.Name)
				existing.Phone = strings.TrimSpace(input.Phone)
				if updateErr := s.repo.Update(ctx, existing); updateErr != nil {
					return nil, apperror.Internal("failed to update existing guest user", updateErr)
				}
				return existing, nil
			}
		}
		return nil, err
	}
	return user, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return apperror.NotFound("user not found", err)
		}
		return apperror.Internal("failed to delete user", err)
	}
	return nil
}
