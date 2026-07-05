package usecase_test

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"shopwise/apps/server/internal/platform/apperror"
	"shopwise/apps/server/internal/users/domain"
	"shopwise/apps/server/internal/users/usecase"

	"github.com/google/uuid"
)

type mockRepo struct {
	users map[uuid.UUID]*domain.User
}

func newMockRepo() *mockRepo {
	return &mockRepo{users: make(map[uuid.UUID]*domain.User)}
}

func (m *mockRepo) Create(_ context.Context, user *domain.User) error {
	for _, existing := range m.users {
		if existing.Email == user.Email {
			return domain.ErrDuplicateEmail
		}
	}
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	copyUser := *user
	m.users[user.ID] = &copyUser
	return nil
}

func (m *mockRepo) GetByID(_ context.Context, id uuid.UUID) (*domain.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	copyUser := *user
	return &copyUser, nil
}

func (m *mockRepo) List(_ context.Context, limit, offset int) ([]domain.User, error) {
	items := make([]domain.User, 0, len(m.users))
	for _, user := range m.users {
		items = append(items, *user)
	}
	if offset >= len(items) {
		return []domain.User{}, nil
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	return items[offset:end], nil
}

func (m *mockRepo) Update(_ context.Context, user *domain.User) error {
	if _, ok := m.users[user.ID]; !ok {
		return domain.ErrNotFound
	}
	for id, existing := range m.users {
		if id != user.ID && existing.Email == user.Email {
			return domain.ErrDuplicateEmail
		}
	}
	copyUser := *user
	m.users[user.ID] = &copyUser
	return nil
}

func (m *mockRepo) Delete(_ context.Context, id uuid.UUID) error {
	if _, ok := m.users[id]; !ok {
		return domain.ErrNotFound
	}
	delete(m.users, id)
	return nil
}

func (m *mockRepo) GetProfile(_ context.Context, userID uuid.UUID) (*domain.UserProfile, error) {
	user, ok := m.users[userID]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return &domain.UserProfile{
		User:        *user,
		Preferences: domain.DefaultPreferences(userID),
	}, nil
}

func (m *mockRepo) UpdateProfileName(_ context.Context, userID uuid.UUID, name string) error {
	user, ok := m.users[userID]
	if !ok {
		return domain.ErrNotFound
	}
	user.Name = name
	return nil
}

func (m *mockRepo) UpsertPreferences(_ context.Context, prefs domain.UserPreferences) error {
	if _, ok := m.users[prefs.UserID]; !ok {
		return domain.ErrNotFound
	}
	return nil
}

type mockEnqueuer struct {
	calls []string
	err   error
}

func (m *mockEnqueuer) EnqueueWelcomeEmail(_ context.Context, userID, email string) error {
	m.calls = append(m.calls, userID+":"+email)
	return m.err
}

func TestServiceCreateEnqueuesWelcomeEmail(t *testing.T) {
	repo := newMockRepo()
	enqueuer := &mockEnqueuer{}
	svc := usecase.NewService(repo, enqueuer, slog.New(slog.NewTextHandler(os.Stdout, nil)))

	user, err := svc.Create(context.Background(), usecase.CreateInput{
		Email: "jane@example.com",
		Name:  "Jane",
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if len(enqueuer.calls) != 1 {
		t.Fatalf("expected 1 enqueue call, got %d", len(enqueuer.calls))
	}
	if user.Email != "jane@example.com" {
		t.Fatalf("email = %q", user.Email)
	}
}

func TestServiceCreateDuplicateEmail(t *testing.T) {
	repo := newMockRepo()
	svc := usecase.NewService(repo, &mockEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil)))

	_, err := svc.Create(context.Background(), usecase.CreateInput{Email: "jane@example.com", Name: "Jane"})
	if err != nil {
		t.Fatalf("first Create() error = %v", err)
	}

	_, err = svc.Create(context.Background(), usecase.CreateInput{Email: "jane@example.com", Name: "Janet"})
	var appErr *apperror.AppError
	if !errors.As(err, &appErr) || appErr.Code != "CONFLICT" {
		t.Fatalf("expected CONFLICT, got %v", err)
	}
}

func TestServiceGetByIDNotFound(t *testing.T) {
	svc := usecase.NewService(newMockRepo(), &mockEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil)))

	_, err := svc.GetByID(context.Background(), uuid.New())
	var appErr *apperror.AppError
	if !errors.As(err, &appErr) || appErr.Code != "NOT_FOUND" {
		t.Fatalf("expected NOT_FOUND, got %v", err)
	}
}
