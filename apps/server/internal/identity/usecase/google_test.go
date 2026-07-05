package usecase_test

import (
	"context"
	"testing"
	"time"

	"shopwise/apps/server/internal/identity/domain"
	"shopwise/apps/server/internal/identity/usecase"

	"github.com/google/uuid"
)

type googleExchangerStub struct {
	info *usecase.GoogleUserInfo
	err  error
}

func (s googleExchangerStub) Exchange(context.Context, string) (*usecase.GoogleUserInfo, error) {
	return s.info, s.err
}

func (s googleExchangerStub) AuthCodeURL(state string) string {
	return "https://accounts.google.com/o/oauth2/auth?state=" + state
}

type memoryAuthRepo struct {
	users      map[string]*domain.AuthUser
	byID       map[uuid.UUID]*domain.AuthUser
	identities map[string]uuid.UUID
}

func newMemoryAuthRepo() *memoryAuthRepo {
	return &memoryAuthRepo{
		users:      map[string]*domain.AuthUser{},
		byID:       map[uuid.UUID]*domain.AuthUser{},
		identities: map[string]uuid.UUID{},
	}
}

func (m *memoryAuthRepo) CreateWithPassword(_ context.Context, email, name, passwordHash string) (*domain.AuthUser, error) {
	if _, ok := m.users[email]; ok {
		return nil, domain.ErrDuplicateEmail
	}
	user := &domain.AuthUser{
		ID:           uuid.New(),
		Email:        email,
		Name:         name,
		PasswordHash: &passwordHash,
	}
	m.users[email] = user
	m.byID[user.ID] = user
	return user, nil
}

func (m *memoryAuthRepo) GetByEmail(_ context.Context, email string) (*domain.AuthUser, error) {
	user, ok := m.users[email]
	if !ok {
		return nil, domain.ErrInvalidCredentials
	}
	return user, nil
}

func (m *memoryAuthRepo) GetByID(_ context.Context, id uuid.UUID) (*domain.AuthUser, error) {
	user, ok := m.byID[id]
	if !ok {
		return nil, domain.ErrInvalidCredentials
	}
	return user, nil
}

func (m *memoryAuthRepo) MarkEmailVerified(_ context.Context, userID uuid.UUID, verifiedAt time.Time) error {
	user, ok := m.byID[userID]
	if !ok {
		return domain.ErrInvalidCredentials
	}
	user.EmailVerifiedAt = &verifiedAt
	return nil
}

func (m *memoryAuthRepo) FindByProvider(_ context.Context, provider, subject string) (*domain.AuthUser, error) {
	key := provider + ":" + subject
	userID, ok := m.identities[key]
	if !ok {
		return nil, domain.ErrInvalidCredentials
	}
	return m.GetByID(context.Background(), userID)
}

func (m *memoryAuthRepo) Link(_ context.Context, userID uuid.UUID, provider, subject string) error {
	m.identities[provider+":"+subject] = userID
	return nil
}

func (m *memoryAuthRepo) CreateGoogleUser(_ context.Context, email, name, provider, subject string, verifiedAt *time.Time) (*domain.AuthUser, error) {
	if existing, ok := m.users[email]; ok {
		_ = m.Link(context.Background(), existing.ID, provider, subject)
		if verifiedAt != nil && existing.EmailVerifiedAt == nil {
			existing.EmailVerifiedAt = verifiedAt
		}
		return existing, nil
	}
	user := &domain.AuthUser{
		ID:              uuid.New(),
		Email:           email,
		Name:            name,
		EmailVerifiedAt: verifiedAt,
		PasswordHash:    nil,
	}
	m.users[email] = user
	m.byID[user.ID] = user
	_ = m.Link(context.Background(), user.ID, provider, subject)
	return user, nil
}

func (m *memoryAuthRepo) Store(_ context.Context, userID uuid.UUID, tokenHash string, _ time.Time) error {
	_ = userID
	_ = tokenHash
	return nil
}

func (m *memoryAuthRepo) FindValid(_ context.Context, _ string, _ time.Time) (uuid.UUID, error) {
	return uuid.Nil, domain.ErrInvalidRefreshToken
}

func (m *memoryAuthRepo) Revoke(_ context.Context, _ string, _ time.Time) error {
	return nil
}

func (m *memoryAuthRepo) Create(_ context.Context, _ uuid.UUID, _ string, _ time.Time) error {
	return nil
}

func (m *memoryAuthRepo) FindValidByHash(_ context.Context, _ string, _ time.Time) (uuid.UUID, error) {
	return uuid.Nil, domain.ErrInvalidVerificationToken
}

func (m *memoryAuthRepo) MarkUsed(_ context.Context, _ string, _ time.Time) error {
	return nil
}

func TestGoogleCallbackCreatesGoogleOnlyUser(t *testing.T) {
	repo := newMemoryAuthRepo()
	jwtSvc := usecase.NewJWTService("secret", time.Minute)
	svc := usecase.NewService(repo, repo, repo, repo, jwtSvc, time.Hour, googleExchangerStub{
		info: &usecase.GoogleUserInfo{
			Subject:       "google-sub-1",
			Email:         "google@example.com",
			Name:          "Google User",
			EmailVerified: true,
		},
	})

	tokens, err := svc.GoogleCallback(context.Background(), "auth-code")
	if err != nil {
		t.Fatalf("GoogleCallback: %v", err)
	}
	if tokens.AccessToken == "" || tokens.RefreshToken == "" {
		t.Fatal("expected token pair")
	}

	user, err := repo.GetByEmail(context.Background(), "google@example.com")
	if err != nil {
		t.Fatalf("GetByEmail: %v", err)
	}
	if user.PasswordHash != nil {
		t.Fatal("google-only user must not have password hash")
	}
	if user.EmailVerifiedAt == nil {
		t.Fatal("expected email verified for google user")
	}
}

func TestGoogleCallbackLinksExistingEmailUser(t *testing.T) {
	repo := newMemoryAuthRepo()
	hash := "hashed"
	existing, err := repo.CreateWithPassword(context.Background(), "user@example.com", "User", hash)
	if err != nil {
		t.Fatalf("CreateWithPassword: %v", err)
	}

	jwtSvc := usecase.NewJWTService("secret", time.Minute)
	svc := usecase.NewService(repo, repo, repo, repo, jwtSvc, time.Hour, googleExchangerStub{
		info: &usecase.GoogleUserInfo{
			Subject:       "google-sub-2",
			Email:         "user@example.com",
			Name:          "User",
			EmailVerified: true,
		},
	})

	tokens, err := svc.GoogleCallback(context.Background(), "auth-code")
	if err != nil {
		t.Fatalf("GoogleCallback: %v", err)
	}
	if tokens.AccessToken == "" {
		t.Fatal("expected access token")
	}

	linked, err := repo.FindByProvider(context.Background(), "google", "google-sub-2")
	if err != nil {
		t.Fatalf("FindByProvider: %v", err)
	}
	if linked.ID != existing.ID {
		t.Fatalf("linked user id = %s, want %s", linked.ID, existing.ID)
	}
}

func TestHashTokenDeterministic(t *testing.T) {
	a := domain.HashToken("refresh-token")
	b := domain.HashToken("refresh-token")
	if a != b {
		t.Fatal("hash should be deterministic")
	}
}
