//go:build integration

package postgres_test

import (
	"context"
	"testing"
	"time"

	identitypostgres "shopwise/retail/internal/identity/repository/postgres"
	integration "shopwise/retail/internal/platform/testutil/integration"

	"github.com/google/uuid"
)

func TestIdentityRepositoryIntegration(t *testing.T) {
	db, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	repo := identitypostgres.NewRepository(db)
	ctx := context.Background()

	user, err := repo.CreateWithPassword(ctx, "identity@example.com", "Identity User", "hash")
	if err != nil {
		t.Fatalf("CreateWithPassword() error = %v", err)
	}

	got, err := repo.GetByEmail(ctx, "identity@example.com")
	if err != nil {
		t.Fatalf("GetByEmail() error = %v", err)
	}
	if got.ID != user.ID {
		t.Fatalf("expected user id %v, got %v", user.ID, got.ID)
	}

	if err := repo.Link(ctx, user.ID, "google", "google-subject-1"); err != nil {
		t.Fatalf("Link() error = %v", err)
	}

	verifiedAt := time.Now().UTC()
	if err := repo.MarkEmailVerified(ctx, user.ID, verifiedAt); err != nil {
		t.Fatalf("MarkEmailVerified() error = %v", err)
	}

	refreshHash := "refresh-token-hash"
	expiresAt := time.Now().UTC().Add(24 * time.Hour)
	if err := repo.Store(ctx, user.ID, refreshHash, expiresAt); err != nil {
		t.Fatalf("Store() error = %v", err)
	}

	storedUserID, err := repo.FindValid(ctx, refreshHash, time.Now().UTC())
	if err != nil {
		t.Fatalf("FindValid() error = %v", err)
	}
	if storedUserID != user.ID {
		t.Fatalf("expected refresh user id %v, got %v", user.ID, storedUserID)
	}

	if err := repo.DeleteUserData(ctx, nil, user.ID); err != nil {
		t.Fatalf("DeleteUserData() error = %v", err)
	}

	if _, err := repo.FindValid(ctx, refreshHash, time.Now().UTC()); err == nil {
		t.Fatal("expected refresh token to be deleted")
	}
}

func TestIdentityRepositoryDeleteUserDataClearsTokens(t *testing.T) {
	db, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	repo := identitypostgres.NewRepository(db)
	ctx := context.Background()
	userID := uuid.New()

	if err := repo.Store(ctx, userID, "token-to-delete", time.Now().UTC().Add(time.Hour)); err != nil {
		t.Fatalf("Store() error = %v", err)
	}
	if err := repo.DeleteUserData(ctx, nil, userID); err != nil {
		t.Fatalf("DeleteUserData() error = %v", err)
	}
	if _, err := repo.FindValid(ctx, "token-to-delete", time.Now().UTC()); err == nil {
		t.Fatal("expected refresh token to be deleted")
	}
}
