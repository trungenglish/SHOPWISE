package handler_test

import (
	"net/http"
	"testing"

	"shopwise/retail/internal/identity/handler"
	"shopwise/retail/internal/platform/testutil"
	"shopwise/retail/internal/platform/testutil/contract"
)

func TestLoginContractSuccess(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/identity/login", "POST")
	contract.AssertResponseCode(t, spec, "/identity/login", "POST", 200)

	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))

	testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/register", map[string]string{
		"email":    "login@example.com",
		"password": "password123",
		"name":     "Login User",
	}, nil)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/login", map[string]string{
		"email":    "login@example.com",
		"password": "password123",
	}, nil)
	testutil.AssertStatus(t, recorder, http.StatusOK)

	var resp handler.TokenResponse
	testutil.AssertJSON(t, recorder, &resp)
	if resp.AccessToken == "" || resp.RefreshToken == "" || resp.ExpiresIn <= 0 {
		t.Fatalf("unexpected token response: %+v", resp)
	}
}

func TestLoginContractInvalidCredentials(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertResponseCode(t, spec, "/identity/login", "POST", 401)

	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/login", map[string]string{
		"email":    "missing@example.com",
		"password": "password123",
	}, nil)
	testutil.AssertStatus(t, recorder, http.StatusUnauthorized)
}

func TestRefreshContractRotatesToken(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/identity/refresh", "POST")
	contract.AssertResponseCode(t, spec, "/identity/refresh", "POST", 200)

	repo := newAuthRepoStub()
	router := newIdentityRouter(newTestIdentityHandler(repo))

	testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/register", map[string]string{
		"email":    "refresh@example.com",
		"password": "password123",
		"name":     "Refresh User",
	}, nil)
	loginRecorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/login", map[string]string{
		"email":    "refresh@example.com",
		"password": "password123",
	}, nil)
	var loginResp handler.TokenResponse
	testutil.AssertJSON(t, loginRecorder, &loginResp)

	refreshRecorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/identity/refresh", map[string]string{
		"refreshToken": loginResp.RefreshToken,
	}, nil)
	testutil.AssertStatus(t, refreshRecorder, http.StatusOK)

	var refreshResp handler.TokenResponse
	testutil.AssertJSON(t, refreshRecorder, &refreshResp)
	if refreshResp.AccessToken == "" {
		t.Fatal("expected new access token")
	}
}
