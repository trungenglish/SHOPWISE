package handler_test

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	identityusecase "shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/middleware"
	"shopwise/retail/internal/platform/testutil"
	"shopwise/retail/internal/users/domain"
	"shopwise/retail/internal/users/handler"
	"shopwise/retail/internal/users/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handlerRepoStub struct {
	createFn func(ctx context.Context, user *domain.User) error
}

func (s *handlerRepoStub) Create(ctx context.Context, user *domain.User) error {
	if s.createFn != nil {
		return s.createFn(ctx, user)
	}
	return nil
}

func (s *handlerRepoStub) GetByID(context.Context, uuid.UUID) (*domain.User, error) {
	return nil, domain.ErrNotFound
}

func (s *handlerRepoStub) List(context.Context, int, int) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (s *handlerRepoStub) Update(context.Context, *domain.User) error {
	return domain.ErrNotFound
}

func (s *handlerRepoStub) Delete(context.Context, uuid.UUID) error {
	return domain.ErrNotFound
}

func (s *handlerRepoStub) GetProfile(context.Context, uuid.UUID) (*domain.UserProfile, error) {
	return nil, domain.ErrNotFound
}

func (s *handlerRepoStub) UpdateProfileName(context.Context, uuid.UUID, string) error {
	return domain.ErrNotFound
}

func (s *handlerRepoStub) UpsertPreferences(context.Context, domain.UserPreferences) error {
	return domain.ErrNotFound
}

type noopEnqueuer struct{}

func (noopEnqueuer) EnqueueWelcomeEmail(context.Context, string, string) error { return nil }

func newUsersRouter(h *handler.Handler) *gin.Engine {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	jwtSvc := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	router := testutil.NewTestRouter()
	router.Use(middleware.ErrorHandler(log, true))
	v1 := router.Group("/api/v1")
	handler.RegisterRoutes(v1.Group("/users"), h, jwtSvc)
	return router
}

func TestHandlerCreateSuccess(t *testing.T) {
	now := time.Now().UTC()
	id := uuid.New()

	h := handler.NewHandler(usecase.NewService(&handlerRepoStub{
		createFn: func(_ context.Context, user *domain.User) error {
			user.ID = id
			user.CreatedAt = now
			user.UpdatedAt = now
			return nil
		},
	}, noopEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil))))
	router := newUsersRouter(h)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/users", map[string]string{
		"email": "jane@example.com",
		"name":  "Jane",
	}, nil)

	testutil.AssertStatus(t, recorder, http.StatusCreated)

	var resp handler.UserResponse
	testutil.AssertJSON(t, recorder, &resp)
	if resp.Email != "jane@example.com" {
		t.Fatalf("email = %q", resp.Email)
	}
}

func TestHandlerCreateValidationError(t *testing.T) {
	h := handler.NewHandler(usecase.NewService(&handlerRepoStub{}, noopEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil))))
	router := newUsersRouter(h)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/users", map[string]string{
		"email": "",
		"name":  "",
	}, nil)

	testutil.AssertStatus(t, recorder, http.StatusBadRequest)
}

func TestHandlerListEmpty(t *testing.T) {
	h := handler.NewHandler(usecase.NewService(&handlerRepoStub{}, noopEnqueuer{}, slog.New(slog.NewTextHandler(os.Stdout, nil))))
	router := newUsersRouter(h)

	recorder := testutil.PerformRequest(t, router, http.MethodGet, "/api/v1/users", nil, nil)
	testutil.AssertStatus(t, recorder, http.StatusOK)

	var resp handler.UserListResponse
	testutil.AssertJSON(t, recorder, &resp)
	if len(resp.Items) != 0 {
		t.Fatalf("expected empty list, got %d", len(resp.Items))
	}
}
