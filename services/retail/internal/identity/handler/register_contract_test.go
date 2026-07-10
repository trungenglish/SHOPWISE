package handler_test

import (
	"net/http"
	"testing"

	"shopwise/retail/internal/platform/testutil"
	"shopwise/retail/internal/platform/testutil/contract"
)

func TestRegisterContractSuccess(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/identity/register", "POST")
	contract.AssertResponseCode(t, spec, "/identity/register", "POST", 201)

	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/register", map[string]string{
		"email":    "new@example.com",
		"password": "password123",
		"name":     "New User",
	}, nil)
	testutil.AssertStatus(t, recorder, http.StatusCreated)
	if repo.verifyCalls != 1 {
		t.Fatalf("verifyCalls = %d, want 1", repo.verifyCalls)
	}
}

func TestRegisterContractDuplicateEmail(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertResponseCode(t, spec, "/identity/register", "POST", 409)

	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))

	body := map[string]string{
		"email":    "dup@example.com",
		"password": "password123",
		"name":     "Dup User",
	}
	testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/register", body, nil)
	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/register", body, nil)
	testutil.AssertStatus(t, recorder, http.StatusConflict)
}

func TestRegisterContractValidationError(t *testing.T) {
	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/register", map[string]string{
		"email":    "bad-email",
		"password": "short",
		"name":     "",
	}, nil)
	testutil.AssertStatus(t, recorder, http.StatusBadRequest)
}
