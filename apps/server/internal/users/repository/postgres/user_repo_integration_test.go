//go:build integration

package postgres_test

import (
	"context"
	"testing"

	"shopwise/apps/server/internal/platform/testutil/integration"
	"shopwise/apps/server/internal/users/domain"
	userpostgres "shopwise/apps/server/internal/users/repository/postgres"

	"github.com/google/uuid"
)

func TestRepositoryCRUD(t *testing.T) {
	db, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	repo := userpostgres.NewRepository(db)
	ctx := context.Background()

	user := &domain.User{
		Email: "jane@example.com",
		Name:  "Jane Doe",
	}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if user.ID == uuid.Nil {
		t.Fatal("expected generated id")
	}

	got, err := repo.GetByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if got.Email != user.Email {
		t.Fatalf("email = %q", got.Email)
	}

	user.Name = "Jane Updated"
	if err := repo.Update(ctx, user); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	users, err := repo.List(ctx, 10, 0)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("list len = %d", len(users))
	}

	if err := repo.Delete(ctx, user.ID); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	_, err = repo.GetByID(ctx, user.ID)
	if err == nil {
		t.Fatal("expected not found after delete")
	}
}

func TestRepositoryDuplicateEmail(t *testing.T) {
	db, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	repo := userpostgres.NewRepository(db)
	ctx := context.Background()

	if err := repo.Create(ctx, &domain.User{Email: "dup@example.com", Name: "One"}); err != nil {
		t.Fatalf("first Create() error = %v", err)
	}

	err := repo.Create(ctx, &domain.User{Email: "dup@example.com", Name: "Two"})
	if err == nil {
		t.Fatal("expected duplicate email error")
	}
	if err != domain.ErrDuplicateEmail {
		t.Fatalf("error = %v", err)
	}
}
