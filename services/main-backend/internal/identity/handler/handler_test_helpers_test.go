package handler_test

import (
	"context"
	"log/slog"
	"os"
	"time"

	"shopwise/retail/internal/identity/domain"
	"shopwise/retail/internal/identity/handler"
	"shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/middleware"
	"shopwise/retail/internal/platform/testutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type authRepoStub struct {
	users       map[string]*domain.AuthUser
	byID        map[uuid.UUID]*domain.AuthUser
	refresh     map[string]uuid.UUID
	verifyCalls int
	tokens      map[string]verifyTokenStub
}

type verifyTokenStub struct {
	userID    uuid.UUID
	expiresAt time.Time
	usedAt    *time.Time
}

func newAuthRepoStub() *authRepoStub {
	return &authRepoStub{
		users:   map[string]*domain.AuthUser{},
		byID:    map[uuid.UUID]*domain.AuthUser{},
		refresh: map[string]uuid.UUID{},
		tokens:  map[string]verifyTokenStub{},
	}
}

func (s *authRepoStub) CreateWithPassword(_ context.Context, email, name, passwordHash string) (*domain.AuthUser, error) {
	if _, ok := s.users[email]; ok {
		return nil, domain.ErrDuplicateEmail
	}
	user := &domain.AuthUser{
		ID:           uuid.New(),
		Email:        email,
		Name:         name,
		PasswordHash: &passwordHash,
	}
	s.users[email] = user
	s.byID[user.ID] = user
	return user, nil
}

func (s *authRepoStub) GetByEmail(_ context.Context, email string) (*domain.AuthUser, error) {
	user, ok := s.users[email]
	if !ok {
		return nil, domain.ErrInvalidCredentials
	}
	return user, nil
}

func (s *authRepoStub) GetByID(_ context.Context, id uuid.UUID) (*domain.AuthUser, error) {
	user, ok := s.byID[id]
	if !ok {
		return nil, domain.ErrInvalidCredentials
	}
	return user, nil
}

func (s *authRepoStub) MarkEmailVerified(_ context.Context, userID uuid.UUID, verifiedAt time.Time) error {
	user, ok := s.byID[userID]
	if !ok {
		return domain.ErrInvalidCredentials
	}
	user.EmailVerifiedAt = &verifiedAt
	return nil
}

func (s *authRepoStub) FindByProvider(context.Context, string, string) (*domain.AuthUser, error) {
	return nil, domain.ErrInvalidCredentials
}

func (s *authRepoStub) Link(context.Context, uuid.UUID, string, string) error {
	return nil
}

func (s *authRepoStub) CreateGoogleUser(context.Context, string, string, string, string, *time.Time) (*domain.AuthUser, error) {
	return nil, domain.ErrInvalidCredentials
}

func (s *authRepoStub) Store(_ context.Context, userID uuid.UUID, tokenHash string, _ time.Time) error {
	s.refresh[tokenHash] = userID
	return nil
}

func (s *authRepoStub) FindValid(_ context.Context, tokenHash string, _ time.Time) (uuid.UUID, error) {
	userID, ok := s.refresh[tokenHash]
	if !ok {
		return uuid.Nil, domain.ErrInvalidRefreshToken
	}
	return userID, nil
}

func (s *authRepoStub) Revoke(_ context.Context, tokenHash string, _ time.Time) error {
	delete(s.refresh, tokenHash)
	return nil
}

func (s *authRepoStub) Create(_ context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error {
	s.verifyCalls++
	s.tokens[tokenHash] = verifyTokenStub{userID: userID, expiresAt: expiresAt}
	return nil
}

func (s *authRepoStub) FindValidByHash(_ context.Context, tokenHash string, now time.Time) (uuid.UUID, error) {
	record, ok := s.tokens[tokenHash]
	if !ok {
		return uuid.Nil, domain.ErrInvalidVerificationToken
	}
	if record.usedAt != nil {
		return uuid.Nil, domain.ErrVerificationTokenUsed
	}
	if !record.expiresAt.After(now) {
		return uuid.Nil, domain.ErrVerificationTokenExpired
	}
	return record.userID, nil
}

func (s *authRepoStub) MarkUsed(_ context.Context, tokenHash string, usedAt time.Time) error {
	record, ok := s.tokens[tokenHash]
	if !ok || record.usedAt != nil {
		return domain.ErrInvalidVerificationToken
	}
	record.usedAt = &usedAt
	s.tokens[tokenHash] = record
	return nil
}

func newIdentityRouter(h *handler.Handler) *gin.Engine {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := testutil.NewTestRouter()
	router.Use(middleware.ErrorHandler(log, true))
	v1 := router.Group("/api/v1")
	handler.RegisterRoutes(v1.Group("/identity"), h)
	return router
}

func newTestIdentityHandler(repo *authRepoStub) *handler.Handler {
	jwtSvc := usecase.NewJWTService("test-secret", 15*time.Minute)
	svc := usecase.NewService(repo, repo, repo, repo, jwtSvc, time.Hour, nil)
	return handler.NewHandler(svc)
}
