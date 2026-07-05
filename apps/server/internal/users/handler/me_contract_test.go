package handler_test

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	identityusecase "shopwise/apps/server/internal/identity/usecase"
	"shopwise/apps/server/internal/platform/middleware"
	"shopwise/apps/server/internal/platform/testutil"
	"shopwise/apps/server/internal/platform/testutil/contract"
	"shopwise/apps/server/internal/users/domain"
	"shopwise/apps/server/internal/users/handler"
	"shopwise/apps/server/internal/users/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type profileRepoStub struct {
	profiles map[uuid.UUID]*domain.UserProfile
}

func newProfileRepoStub() *profileRepoStub {
	return &profileRepoStub{profiles: map[uuid.UUID]*domain.UserProfile{}}
}

func (s *profileRepoStub) seedProfile(userID uuid.UUID) *domain.UserProfile {
	now := time.Now().UTC()
	profile := &domain.UserProfile{
		User: domain.User{
			ID:        userID,
			Email:     "me@example.com",
			Name:      "Me User",
			CreatedAt: now,
			UpdatedAt: now,
		},
		Preferences: domain.DefaultPreferences(userID),
	}
	s.profiles[userID] = profile
	return profile
}

func (s *profileRepoStub) Create(context.Context, *domain.User) error { return nil }
func (s *profileRepoStub) GetByID(context.Context, uuid.UUID) (*domain.User, error) {
	return nil, domain.ErrNotFound
}
func (s *profileRepoStub) List(context.Context, int, int) ([]domain.User, error) {
	return nil, nil
}
func (s *profileRepoStub) Update(context.Context, *domain.User) error { return domain.ErrNotFound }
func (s *profileRepoStub) Delete(_ context.Context, id uuid.UUID) error {
	delete(s.profiles, id)
	return nil
}

func (s *profileRepoStub) GetProfile(_ context.Context, userID uuid.UUID) (*domain.UserProfile, error) {
	profile, ok := s.profiles[userID]
	if !ok {
		return nil, domain.ErrNotFound
	}
	copy := *profile
	return &copy, nil
}

func (s *profileRepoStub) UpdateProfileName(_ context.Context, userID uuid.UUID, name string) error {
	profile, ok := s.profiles[userID]
	if !ok {
		return domain.ErrNotFound
	}
	profile.Name = name
	return nil
}

func (s *profileRepoStub) UpsertPreferences(_ context.Context, prefs domain.UserPreferences) error {
	profile, ok := s.profiles[prefs.UserID]
	if !ok {
		return domain.ErrNotFound
	}
	profile.Preferences = prefs
	return nil
}

func newMeRouter(h *handler.Handler, jwtSvc *identityusecase.JWTService) *gin.Engine {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := testutil.NewTestRouter()
	router.Use(middleware.ErrorHandler(log, true))
	v1 := router.Group("/api/v1")
	handler.RegisterRoutes(v1.Group("/users"), h, jwtSvc)
	return router
}

func TestGetMeContractSuccess(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/users/me", "GET")
	contract.AssertResponseCode(t, spec, "/users/me", "GET", 200)

	repo := newProfileRepoStub()
	userID := uuid.New()
	repo.seedProfile(userID)
	jwtSvc := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	token, _, err := jwtSvc.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("IssueAccessToken: %v", err)
	}

	h := handler.NewHandler(usecase.NewService(repo, noopEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil))))
	router := newMeRouter(h, jwtSvc)

	recorder := testutil.PerformRequest(t, router, http.MethodGet, "/api/v1/users/me", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	testutil.AssertStatus(t, recorder, http.StatusOK)

	var resp handler.MeResponse
	testutil.AssertJSON(t, recorder, &resp)
	if resp.Email != "me@example.com" || resp.Preferences.BudgetSensitivity == "" {
		t.Fatalf("unexpected me response: %+v", resp)
	}
}

func TestGetMeContractUnauthorized(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertResponseCode(t, spec, "/users/me", "GET", 401)

	repo := newProfileRepoStub()
	jwtSvc := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	h := handler.NewHandler(usecase.NewService(repo, noopEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil))))
	router := newMeRouter(h, jwtSvc)

	recorder := testutil.PerformRequest(t, router, http.MethodGet, "/api/v1/users/me", nil, nil)
	testutil.AssertStatus(t, recorder, http.StatusUnauthorized)
}

func TestPatchMeContractSuccess(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/users/me", "PATCH")
	contract.AssertResponseCode(t, spec, "/users/me", "PATCH", 200)

	repo := newProfileRepoStub()
	userID := uuid.New()
	repo.seedProfile(userID)
	jwtSvc := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	token, _, err := jwtSvc.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("IssueAccessToken: %v", err)
	}

	h := handler.NewHandler(usecase.NewService(repo, noopEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil))))
	router := newMeRouter(h, jwtSvc)

	recorder := testutil.PerformRequest(t, router, http.MethodPatch, "/api/v1/users/me", map[string]any{
		"name":              "Updated Name",
		"budgetSensitivity": "high",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	testutil.AssertStatus(t, recorder, http.StatusOK)

	var resp handler.MeResponse
	testutil.AssertJSON(t, recorder, &resp)
	if resp.Name != "Updated Name" || resp.Preferences.BudgetSensitivity != "high" {
		t.Fatalf("unexpected patch response: %+v", resp)
	}
}
