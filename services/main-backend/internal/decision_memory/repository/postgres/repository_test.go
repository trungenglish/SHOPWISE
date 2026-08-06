//go:build integration

package postgres_test

import (
	"context"
	"testing"
	"time"

	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/decision_memory/repository/postgres"
	"shopwise/retail/internal/platform/testutil/integration"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// setupDB initializes a test database connection. Note: Adjust based on real project setup.
// Assuming a generic mock or in-memory sqlite/postgres setup is available in the workspace.
func setupDB(t *testing.T) *gorm.DB {
	db, cleanup := integration.SetupIntegrationDB(t)
	t.Cleanup(cleanup)
	db.AutoMigrate(&postgres.DecisionSession{}, &postgres.SessionMessage{}, &postgres.ExtractedPreference{})
	return db
}

func TestRepository_UpdateSession_LWW(t *testing.T) {
	db := setupDB(t)
	if db == nil { return }
	repo := postgres.NewRepository(db)
	ctx := context.Background()

	uid := uuid.New()
	session := &domain.DecisionSession{
		ID:        uuid.New(),
		UserID:    &uid,
		Title:     "LWW Test",
		Status:    "active",
		CreatedAt: time.Now().UTC().Add(-1 * time.Hour),
		UpdatedAt: time.Now().UTC().Add(-30 * time.Minute),
	}
	err := repo.CreateSession(ctx, session)
	require.NoError(t, err)

	// Valid update
	clientTsValid := time.Now().UTC()
	err = repo.UpdateSession(ctx, session, clientTsValid)
	assert.NoError(t, err)

	// Fetch updated session
	updated, err := repo.GetSession(ctx, session.ID)
	require.NoError(t, err)

	// Conflict update (older client timestamp than server)
	clientTsConflict := updated.UpdatedAt.Add(-1 * time.Minute)
	err = repo.UpdateSession(ctx, updated, clientTsConflict)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "conflict")
}

func TestRepository_BranchSession_DeepCopy(t *testing.T) {
	db := setupDB(t)
	if db == nil { return }
	repo := postgres.NewRepository(db)
	ctx := context.Background()

	uid := uuid.New()
	original := &domain.DecisionSession{
		ID:        uuid.New(),
		UserID:    &uid,
		Title:     "Original",
		Status:    "active",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	require.NoError(t, repo.CreateSession(ctx, original))

	msg1 := &domain.SessionMessage{
		ID:        uuid.New(),
		SessionID: original.ID,
		Role:      "user",
		Content:   "Hello",
		CreatedAt: time.Now().UTC(),
	}
	require.NoError(t, repo.AddMessage(ctx, msg1))

	branched, err := repo.BranchSession(ctx, original.ID, "Branched")
	require.NoError(t, err)
	assert.NotEqual(t, original.ID, branched.ID)
	assert.Equal(t, "Branched", branched.Title)
	assert.Equal(t, original.ID, *branched.ParentSessionID)

	// Verify messages are copied
	branchedFull, err := repo.GetSession(ctx, branched.ID)
	require.NoError(t, err)
	assert.Len(t, branchedFull.Messages, 1)
	assert.Equal(t, "Hello", branchedFull.Messages[0].Content)
	assert.NotEqual(t, msg1.ID, branchedFull.Messages[0].ID) // new ID
}
