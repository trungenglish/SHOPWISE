package usecase

import (
	"context"
	"fmt"
	"time"

	"shopwise/retail/internal/decision_memory/domain"

	"github.com/google/uuid"
)

type Repository interface {
	CreateSession(ctx context.Context, session *domain.DecisionSession) error
	GetSession(ctx context.Context, id uuid.UUID) (*domain.DecisionSession, error)
	ListSessions(ctx context.Context, userID *uuid.UUID, anonymousID *string, limit, offset int) ([]domain.DecisionSession, error)
	UpdateSession(ctx context.Context, session *domain.DecisionSession, clientTimestamp time.Time) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	CountAnonymousSessions(ctx context.Context, anonymousID string) (int64, error)
	DeleteOldestAnonymousSession(ctx context.Context, anonymousID string) error
	AddMessage(ctx context.Context, msg *domain.SessionMessage) error
	ListPreferences(ctx context.Context, userID uuid.UUID) ([]domain.UserPreference, error)
	UpdatePreference(ctx context.Context, pref *domain.UserPreference) error
	DeletePreference(ctx context.Context, id uuid.UUID, userID uuid.UUID) error
	BranchSession(ctx context.Context, id uuid.UUID, newTitle string) (*domain.DecisionSession, error)
	ArchiveInactiveSessions(ctx context.Context, before time.Time) (int64, error)
}

type Service struct {
	repo Repository
}

type interactionRepository interface {
	BeginInteraction(context.Context, *domain.SessionInteraction) (*domain.SessionInteraction, bool, error)
	FinishInteraction(context.Context, uuid.UUID, uuid.UUID, string, []byte) error
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateSession creates a new session. If it's an anonymous user, enforces the 10-session limit.
func (s *Service) CreateSession(ctx context.Context, userID *uuid.UUID, anonymousID *string, title string) (*domain.DecisionSession, error) {
	if userID == nil && anonymousID == nil {
		return nil, fmt.Errorf("must provide either userID or anonymousID")
	}

	if userID == nil && anonymousID != nil {
		// Enforce anonymous limit
		count, err := s.repo.CountAnonymousSessions(ctx, *anonymousID)
		if err != nil {
			return nil, fmt.Errorf("failed to count anonymous sessions: %w", err)
		}
		if count >= 10 {
			if err := s.repo.DeleteOldestAnonymousSession(ctx, *anonymousID); err != nil {
				return nil, fmt.Errorf("failed to evict oldest anonymous session: %w", err)
			}
		}
	}

	session := &domain.DecisionSession{
		ID:          uuid.New(),
		UserID:      userID,
		AnonymousID: anonymousID,
		Title:       title,
		Status:      "active",
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := s.repo.CreateSession(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (s *Service) GetSession(ctx context.Context, id uuid.UUID) (*domain.DecisionSession, error) {
	return s.repo.GetSession(ctx, id)
}

func (s *Service) ListSessions(ctx context.Context, userID *uuid.UUID, anonymousID *string, limit, offset int) ([]domain.DecisionSession, error) {
	return s.repo.ListSessions(ctx, userID, anonymousID, limit, offset)
}

func (s *Service) UpdateSession(ctx context.Context, session *domain.DecisionSession, clientTimestamp time.Time) error {
	return s.repo.UpdateSession(ctx, session, clientTimestamp)
}

func (s *Service) RenameSession(ctx context.Context, id uuid.UUID, title string, clientTimestamp time.Time) error {
	session, err := s.repo.GetSession(ctx, id)
	if err != nil {
		return err
	}
	session.Title = title
	return s.repo.UpdateSession(ctx, session, clientTimestamp)
}

func (s *Service) DeleteSession(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSession(ctx, id)
}

func (s *Service) ListPreferences(ctx context.Context, userID uuid.UUID) ([]domain.UserPreference, error) {
	return s.repo.ListPreferences(ctx, userID)
}

func (s *Service) UpdatePreference(ctx context.Context, pref *domain.UserPreference) error {
	return s.repo.UpdatePreference(ctx, pref)
}

func (s *Service) DeletePreference(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	return s.repo.DeletePreference(ctx, id, userID)
}

func (s *Service) BranchSession(ctx context.Context, id uuid.UUID, newTitle string) (*domain.DecisionSession, error) {
	return s.repo.BranchSession(ctx, id, newTitle)
}

func (s *Service) RestoreSession(ctx context.Context, id uuid.UUID, clientTimestamp time.Time) error {
	session, err := s.repo.GetSession(ctx, id)
	if err != nil {
		return err
	}
	session.Status = "active"
	return s.repo.UpdateSession(ctx, session, clientTimestamp)
}

func (s *Service) AddMessage(ctx context.Context, msg *domain.SessionMessage) error {
	return s.repo.AddMessage(ctx, msg)
}

func (s *Service) BeginInteraction(
	ctx context.Context,
	interaction *domain.SessionInteraction,
) (*domain.SessionInteraction, bool, error) {
	repository, ok := s.repo.(interactionRepository)
	if !ok {
		return nil, false, fmt.Errorf("interaction persistence is unavailable")
	}
	return repository.BeginInteraction(ctx, interaction)
}

func (s *Service) FinishInteraction(
	ctx context.Context,
	sessionID uuid.UUID,
	interactionID uuid.UUID,
	status string,
	response []byte,
) error {
	repository, ok := s.repo.(interactionRepository)
	if !ok {
		return fmt.Errorf("interaction persistence is unavailable")
	}
	return repository.FinishInteraction(ctx, sessionID, interactionID, status, response)
}
