package handler_test

import (
	"net/http"
	"testing"
	"time"

	"context"

	"shopwise/retail/internal/identity/domain"
	"shopwise/retail/internal/platform/testutil"
	"shopwise/retail/internal/platform/testutil/contract"
)

func TestVerifyEmailContractSuccess(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/identity/verify-email", "POST")
	contract.AssertResponseCode(t, spec, "/identity/verify-email", "POST", 200)

	repo := newAuthRepoStub()
	user, err := repo.CreateWithPassword(context.Background(), "verify-contract@example.com", "User", "hash")
	if err != nil {
		t.Fatalf("CreateWithPassword: %v", err)
	}

	rawToken := "verify-contract-token"
	repo.tokens[domain.HashToken(rawToken)] = verifyTokenStub{
		userID:    user.ID,
		expiresAt: time.Now().UTC().Add(time.Hour),
	}

	router := newIdentityRouter(newTestIdentityHandler(repo))
	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/verify-email", map[string]string{
		"token": rawToken,
	}, nil)
	testutil.AssertStatus(t, recorder, http.StatusOK)
}

func TestVerifyEmailContractInvalidToken(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertResponseCode(t, spec, "/identity/verify-email", "POST", 400)

	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))
	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/verify-email", map[string]string{
		"token": "invalid-token",
	}, nil)
	testutil.AssertStatus(t, recorder, http.StatusBadRequest)
}
