//go:build integration

package integration_test

import (
	"testing"

	integration "shopwise/retail/internal/platform/testutil/integration"
)

func TestIntegrationSuiteBootstraps(t *testing.T) {
	db, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()
	if db == nil {
		t.Fatal("expected database connection")
	}
}
