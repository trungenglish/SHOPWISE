package usecase_test

import (
	"testing"
	"time"

	"shopwise/apps/server/internal/identity/usecase"

	"github.com/google/uuid"
)

func TestJWTServiceIssueAndVerify(t *testing.T) {
	svc := usecase.NewJWTService("test-secret-key", 15*time.Minute)
	userID := uuid.New()

	token, expiresIn, err := svc.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("IssueAccessToken: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token")
	}
	if expiresIn <= 0 {
		t.Fatalf("expiresIn = %d", expiresIn)
	}

	subject, err := svc.Verify(token)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}
	if subject != userID.String() {
		t.Fatalf("subject = %q, want %q", subject, userID.String())
	}
}

func TestJWTServiceRejectsWrongSecret(t *testing.T) {
	issuer := usecase.NewJWTService("secret-a", time.Minute)
	verifier := usecase.NewJWTService("secret-b", time.Minute)
	userID := uuid.New()

	token, _, err := issuer.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("IssueAccessToken: %v", err)
	}
	if _, err := verifier.Verify(token); err == nil {
		t.Fatal("expected verify error for wrong secret")
	}
}

func TestJWTServiceRejectsMalformedToken(t *testing.T) {
	svc := usecase.NewJWTService("test-secret", time.Minute)
	if _, err := svc.Verify("not-a-jwt"); err == nil {
		t.Fatal("expected error for malformed token")
	}
}

func TestJWTServiceRejectsExpiredToken(t *testing.T) {
	svc := usecase.NewJWTService("test-secret", -time.Minute)
	userID := uuid.New()

	token, _, err := svc.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("IssueAccessToken: %v", err)
	}
	if _, err := svc.Verify(token); err == nil {
		t.Fatal("expected expired token error")
	}
}
