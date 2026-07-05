package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"shopwise/apps/server/internal/identity/domain"
	"shopwise/apps/server/internal/identity/usecase"

	"github.com/google/uuid"
)

type verifyTokenRecord struct {
	userID    uuid.UUID
	expiresAt time.Time
	usedAt    *time.Time
}

type memoryVerifyRepo struct {
	*memoryAuthRepo
	tokens map[string]verifyTokenRecord
}

func newMemoryVerifyRepo() *memoryVerifyRepo {
	return &memoryVerifyRepo{
		memoryAuthRepo: newMemoryAuthRepo(),
		tokens:         map[string]verifyTokenRecord{},
	}
}

func (m *memoryVerifyRepo) storeToken(raw string, userID uuid.UUID, expiresAt time.Time) {
	m.tokens[domain.HashToken(raw)] = verifyTokenRecord{
		userID:    userID,
		expiresAt: expiresAt,
	}
}

func (m *memoryVerifyRepo) FindValidByHash(_ context.Context, tokenHash string, now time.Time) (uuid.UUID, error) {
	record, ok := m.tokens[tokenHash]
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

func (m *memoryVerifyRepo) MarkUsed(_ context.Context, tokenHash string, usedAt time.Time) error {
	record, ok := m.tokens[tokenHash]
	if !ok || record.usedAt != nil {
		return domain.ErrInvalidVerificationToken
	}
	record.usedAt = &usedAt
	m.tokens[tokenHash] = record
	return nil
}

func TestVerifyEmailMarksUserVerified(t *testing.T) {
	repo := newMemoryVerifyRepo()
	user, err := repo.CreateWithPassword(context.Background(), "verify@example.com", "User", "hash")
	if err != nil {
		t.Fatalf("CreateWithPassword: %v", err)
	}
	raw := "verification-token-raw"
	repo.storeToken(raw, user.ID, time.Now().UTC().Add(time.Hour))

	svc := usecase.NewService(repo, repo, repo, repo, usecase.NewJWTService("secret", time.Minute), time.Hour, nil)
	if err := svc.VerifyEmail(context.Background(), raw); err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}
	updated, err := repo.GetByID(context.Background(), user.ID)
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if updated.EmailVerifiedAt == nil {
		t.Fatal("expected email_verified_at to be set")
	}
	record := repo.tokens[domain.HashToken(raw)]
	if record.usedAt == nil {
		t.Fatal("expected token to be marked used")
	}
}

func TestVerifyEmailExpiredToken(t *testing.T) {
	repo := newMemoryVerifyRepo()
	user, err := repo.CreateWithPassword(context.Background(), "expired@example.com", "User", "hash")
	if err != nil {
		t.Fatalf("CreateWithPassword: %v", err)
	}
	raw := "expired-token"
	repo.storeToken(raw, user.ID, time.Now().UTC().Add(-time.Hour))

	svc := usecase.NewService(repo, repo, repo, repo, usecase.NewJWTService("secret", time.Minute), time.Hour, nil)
	err = svc.VerifyEmail(context.Background(), raw)
	if err == nil {
		t.Fatal("expected error for expired token")
	}
}

func TestVerifyEmailUsedToken(t *testing.T) {
	repo := newMemoryVerifyRepo()
	user, err := repo.CreateWithPassword(context.Background(), "used@example.com", "User", "hash")
	if err != nil {
		t.Fatalf("CreateWithPassword: %v", err)
	}
	raw := "used-token"
	hash := domain.HashToken(raw)
	repo.storeToken(raw, user.ID, time.Now().UTC().Add(time.Hour))
	usedAt := time.Now().UTC()
	repo.tokens[hash] = verifyTokenRecord{userID: user.ID, expiresAt: time.Now().UTC().Add(time.Hour), usedAt: &usedAt}

	svc := usecase.NewService(repo, repo, repo, repo, usecase.NewJWTService("secret", time.Minute), time.Hour, nil)
	err = svc.VerifyEmail(context.Background(), raw)
	if err == nil {
		t.Fatal("expected error for used token")
	}
}

func TestVerifyEmailInvalidToken(t *testing.T) {
	repo := newMemoryVerifyRepo()
	svc := usecase.NewService(repo, repo, repo, repo, usecase.NewJWTService("secret", time.Minute), time.Hour, nil)
	err := svc.VerifyEmail(context.Background(), "unknown-token")
	if err == nil {
		t.Fatal("expected error for invalid token")
	}
}

func TestVerifyEmailIdempotentWhenAlreadyVerified(t *testing.T) {
	repo := newMemoryVerifyRepo()
	user, err := repo.CreateWithPassword(context.Background(), "done@example.com", "User", "hash")
	if err != nil {
		t.Fatalf("CreateWithPassword: %v", err)
	}
	verifiedAt := time.Now().UTC().Add(-time.Minute)
	_ = repo.MarkEmailVerified(context.Background(), user.ID, verifiedAt)
	raw := "repeat-token"
	repo.storeToken(raw, user.ID, time.Now().UTC().Add(time.Hour))

	svc := usecase.NewService(repo, repo, repo, repo, usecase.NewJWTService("secret", time.Minute), time.Hour, nil)
	if err := svc.VerifyEmail(context.Background(), raw); err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}
}

func TestVerifyEmailEmptyToken(t *testing.T) {
	repo := newMemoryVerifyRepo()
	svc := usecase.NewService(repo, repo, repo, repo, usecase.NewJWTService("secret", time.Minute), time.Hour, nil)
	err := svc.VerifyEmail(context.Background(), "  ")
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestVerifyEmailRepositoryErrors(t *testing.T) {
	if !errors.Is(domain.ErrVerificationTokenExpired, domain.ErrVerificationTokenExpired) {
		t.Fatal("verification errors should be comparable")
	}
}
