package handler_test

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	identityusecase "shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/platform/testutil"
	"shopwise/retail/internal/platform/testutil/contract"
	"shopwise/retail/internal/users/domain"
	"shopwise/retail/internal/users/handler"
	"shopwise/retail/internal/users/usecase"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accountDeleteRepoStub struct {
	*profileRepoStub
}

func newAccountDeleteRepoStub() *accountDeleteRepoStub {
	return &accountDeleteRepoStub{profileRepoStub: newProfileRepoStub()}
}

func (s *accountDeleteRepoStub) Delete(_ context.Context, id uuid.UUID) error {
	delete(s.profiles, id)
	return nil
}

func (s *accountDeleteRepoStub) DeleteAccount(_ context.Context, _ *gorm.DB, userID uuid.UUID) error {
	if _, ok := s.profiles[userID]; !ok {
		return domain.ErrNotFound
	}
	delete(s.profiles, userID)
	return nil
}

type noopCascadeDeleter struct{}

func (noopCascadeDeleter) DeleteUserData(context.Context, *gorm.DB, uuid.UUID) error {
	return nil
}

type noopOwnershipDeleter struct{}

func (noopOwnershipDeleter) DeleteUserData(context.Context, *gorm.DB, uuid.UUID) ([]string, error) {
	return nil, nil
}

func newDeleteAccountService(repo *accountDeleteRepoStub) *usecase.Service {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return usecase.NewService(repo, noopEnqueuer{}, log).WithAccountDeletion(usecase.AccountDeletionDeps{
		Identity:     noopCascadeDeleter{},
		Advisor:      noopCascadeDeleter{},
		Wishlist:     noopCascadeDeleter{},
		Ownership:    noopOwnershipDeleter{},
		Notification: noopCascadeDeleter{},
		Users:        repo,
	})
}

func TestDeleteAccountContractSuccess(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertPathMethod(t, spec, "/users/me", "DELETE")
	contract.AssertResponseCode(t, spec, "/users/me", "DELETE", 204)

	repo := newAccountDeleteRepoStub()
	userID := uuid.New()
	repo.seedProfile(userID)
	jwtSvc := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	token, _, err := jwtSvc.IssueAccessToken(userID)
	if err != nil {
		t.Fatalf("IssueAccessToken: %v", err)
	}

	h := handler.NewHandler(newDeleteAccountService(repo))
	router := newMeRouter(h, jwtSvc)

	deleteRecorder := testutil.PerformRequest(t, router, http.MethodDelete, "/api/v1/users/me", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	testutil.AssertStatus(t, deleteRecorder, http.StatusNoContent)

	getRecorder := testutil.PerformRequest(t, router, http.MethodGet, "/api/v1/users/me", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	testutil.AssertStatus(t, getRecorder, http.StatusNotFound)
}

func TestDeleteAccountContractUnauthorized(t *testing.T) {
	spec := contract.LoadSpec(t)
	contract.AssertResponseCode(t, spec, "/users/me", "DELETE", 401)

	repo := newAccountDeleteRepoStub()
	jwtSvc := identityusecase.NewJWTService("test-secret", 15*time.Minute)
	h := handler.NewHandler(newDeleteAccountService(repo))
	router := newMeRouter(h, jwtSvc)

	recorder := testutil.PerformRequest(t, router, http.MethodDelete, "/api/v1/users/me", nil, nil)
	testutil.AssertStatus(t, recorder, http.StatusUnauthorized)
}
