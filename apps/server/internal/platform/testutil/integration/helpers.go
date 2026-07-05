//go:build integration

package integration

import (
	"context"
	"strings"
	"testing"

	identitypostgres "shopwise/apps/server/internal/identity/repository/postgres"
	userpostgres "shopwise/apps/server/internal/users/repository/postgres"

	"github.com/testcontainers/testcontainers-go"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupIntegrationDB(t *testing.T) (*gorm.DB, func()) {
	t.Helper()
	ctx := context.Background()

	container, err := tcpostgres.Run(ctx,
		"postgres:18.4-bookworm",
		tcpostgres.WithDatabase("app"),
		tcpostgres.WithUsername("app"),
		tcpostgres.WithPassword("app"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
		),
	)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "docker") {
			t.Skipf("docker unavailable for integration tests: %v", err)
		}
		t.Fatalf("start postgres container: %v", err)
	}

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("connection string: %v", err)
	}

	db, err := gorm.Open(gormpostgres.Open(connStr), &gorm.Config{})
	if err != nil {
		t.Fatalf("open gorm: %v", err)
	}

	migrations := []func(*gorm.DB) error{
		userpostgres.Migrate,
		identitypostgres.Migrate,
	}
	for _, migrate := range migrations {
		if err := migrate(db); err != nil {
			t.Fatalf("migrate: %v", err)
		}
	}

	cleanup := func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Logf("terminate container: %v", err)
		}
	}
	return db, cleanup
}
