package handler_test

import (
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	identityusecase "shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/testutil"
	"shopwise/retail/internal/users/domain"
	"shopwise/retail/internal/users/handler"
	"shopwise/retail/internal/users/usecase"

	"github.com/google/uuid"
)

func TestPatchMeUpdatesPhone(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	repository := &handlerRepoStub{profile: &domain.UserProfile{
		User:        domain.User{ID: userID, Email: "customer@example.com", Name: "Customer"},
		Preferences: domain.DefaultPreferences(userID),
	}}
	checkoutHandler := handler.NewHandler(usecase.NewService(
		repository,
		noopEnqueuer{},
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	))
	router := newUsersRouter(checkoutHandler)
	jwtService := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	token, _, err := jwtService.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("issue access token: %v", err)
	}

	recorder := testutil.PerformRequest(t, router, http.MethodPatch, "/api/v1/users/me", map[string]any{
		"phone": "0912345678",
	}, map[string]string{"Authorization": "Bearer " + token})

	testutil.AssertStatus(t, recorder, http.StatusOK)
	var response handler.MeResponse
	testutil.AssertJSON(t, recorder, &response)
	if response.Phone != "0912345678" {
		t.Fatalf("phone = %q, want 0912345678", response.Phone)
	}
}
