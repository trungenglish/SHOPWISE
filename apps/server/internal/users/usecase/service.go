package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"shopwise/apps/server/internal/platform/apperror"
	"shopwise/apps/server/internal/users/domain"

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
}

type UpdateInput struct {
	Email *string
	Name  *string
}

func (s *Service) Create(ctx context.Context, input CreateInput) (*domain.User, error) {
	email := strings.TrimSpace(input.Email)
	name := strings.TrimSpace(input.Name)
	if email == "" || name == "" {
		return nil, apperror.Validation("email and name are required", nil)
	}

	user := &domain.User{
		Email: email,
		Name:  name,
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

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return apperror.NotFound("user not found", err)
		}
		return apperror.Internal("failed to delete user", err)
	}
	return nil
}
