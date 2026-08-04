package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/decision_memory/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type chatRepository struct {
	session domain.DecisionSession
}

func (r *chatRepository) CreateSession(context.Context, *domain.DecisionSession) error { return nil }
func (r *chatRepository) GetSession(_ context.Context, _ uuid.UUID) (*domain.DecisionSession, error) {
	copy := r.session
	copy.Messages = append([]domain.SessionMessage(nil), r.session.Messages...)
	return &copy, nil
}
func (r *chatRepository) ListSessions(context.Context, *uuid.UUID, *string, int, int) ([]domain.DecisionSession, error) {
	return nil, nil
}
func (r *chatRepository) UpdateSession(context.Context, *domain.DecisionSession, time.Time) error {
	return nil
}
func (r *chatRepository) DeleteSession(context.Context, uuid.UUID) error { return nil }
func (r *chatRepository) CountAnonymousSessions(context.Context, string) (int64, error) {
	return 0, nil
}
func (r *chatRepository) DeleteOldestAnonymousSession(context.Context, string) error { return nil }
func (r *chatRepository) AddMessage(_ context.Context, message *domain.SessionMessage) error {
	r.session.Messages = append(r.session.Messages, *message)
	return nil
}
func (r *chatRepository) ListPreferences(context.Context, uuid.UUID) ([]domain.UserPreference, error) {
	return nil, nil
}
func (r *chatRepository) UpdatePreference(context.Context, *domain.UserPreference) error { return nil }
func (r *chatRepository) DeletePreference(context.Context, uuid.UUID, uuid.UUID) error {
	return nil
}
func (r *chatRepository) BranchSession(context.Context, uuid.UUID, string) (*domain.DecisionSession, error) {
	return nil, nil
}
func (r *chatRepository) ArchiveInactiveSessions(context.Context, time.Time) (int64, error) {
	return 0, nil
}

func TestChatForwardsHistoryAndPersistsAssistantResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sessionID := uuid.New()
	anonymousID := "anon-1"
	repository := &chatRepository{session: domain.DecisionSession{
		ID:          sessionID,
		AnonymousID: &anonymousID,
		Messages: []domain.SessionMessage{
			{Role: "user", Content: "I need a laptop"},
			{
				Role:           "assistant",
				Content:        "What will you use it for?",
				ReasoningGraph: `{"type":"recommendation","message":"Pick","decision":{"products":[{"id":"product-1"}]}}`,
			},
		},
	}}
	var forwarded struct {
		SessionID            string   `json:"session_id"`
		AllowedComparisonIDs []string `json:"allowed_comparison_ids"`
		Messages             []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
	}
	aiRuntime := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/api/v1/chat" {
			t.Fatalf("AI path = %q", request.URL.Path)
		}
		if err := json.NewDecoder(request.Body).Decode(&forwarded); err != nil {
			t.Fatalf("decode forwarded request: %v", err)
		}
		response.Header().Set("Content-Type", "application/json")
		_, _ = response.Write([]byte(`{"type":"question","message":"What is your budget?"}`))
	}))
	defer aiRuntime.Close()

	handler := NewHandler(usecase.NewService(repository), aiRuntime.URL)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "id", Value: sessionID.String()}}
	context.Request = httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(`{"message":"Gaming"}`),
	)
	context.Request.Header.Set("Content-Type", "application/json")
	context.Request.Header.Set("X-Anonymous-ID", anonymousID)

	handler.Chat(context)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
	}
	if len(forwarded.Messages) != 3 || forwarded.Messages[2].Content != "Gaming" {
		t.Fatalf("forwarded messages = %#v", forwarded.Messages)
	}
	if len(forwarded.AllowedComparisonIDs) != 1 || forwarded.AllowedComparisonIDs[0] != "product-1" {
		t.Fatalf("allowed comparison ids = %#v", forwarded.AllowedComparisonIDs)
	}
	if got := repository.session.Messages[len(repository.session.Messages)-1]; got.Role != "assistant" || got.Content != "What is your budget?" {
		t.Fatalf("saved assistant message = %#v", got)
	}
}

func TestChatRejectsSessionOwnedByAnotherAnonymousUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sessionID := uuid.New()
	ownerID := "owner"
	repository := &chatRepository{session: domain.DecisionSession{
		ID:          sessionID,
		AnonymousID: &ownerID,
	}}
	aiCalled := false
	aiRuntime := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		aiCalled = true
	}))
	defer aiRuntime.Close()

	handler := NewHandler(usecase.NewService(repository), aiRuntime.URL)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "id", Value: sessionID.String()}}
	context.Request = httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(`{"message":"Gaming"}`),
	)
	context.Request.Header.Set("Content-Type", "application/json")
	context.Request.Header.Set("X-Anonymous-ID", "other-user")

	handler.Chat(context)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", recorder.Code)
	}
	if aiCalled {
		t.Fatal("AI Runtime was called for an unauthorized session")
	}
}

func TestGetSessionRejectsAnotherAnonymousUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sessionID := uuid.New()
	ownerID := "owner"
	repository := &chatRepository{session: domain.DecisionSession{
		ID:          sessionID,
		AnonymousID: &ownerID,
		Messages:    []domain.SessionMessage{{Role: "user", Content: "private history"}},
	}}
	handler := NewHandler(usecase.NewService(repository), "http://unused")
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "id", Value: sessionID.String()}}
	context.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	context.Request.Header.Set("X-Anonymous-ID", "other-user")

	handler.GetSession(context)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want 404; body = %s", recorder.Code, recorder.Body.String())
	}
}

func TestChatStreamForwardsSSEAndPersistsAssistantResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sessionID := uuid.New()
	anonymousID := "anon-stream"
	repository := &chatRepository{session: domain.DecisionSession{
		ID:          sessionID,
		AnonymousID: &anonymousID,
	}}
	aiRuntime := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/api/v1/chat/stream" {
			t.Fatalf("AI path = %q", request.URL.Path)
		}
		response.Header().Set("Content-Type", "text/event-stream")
		_, _ = response.Write([]byte("event: token\ndata: {\"text\":\"What is your budget?\"}\n\nevent: done\ndata: {}\n\n"))
	}))
	defer aiRuntime.Close()

	handler := NewHandler(usecase.NewService(repository), aiRuntime.URL)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Params = gin.Params{{Key: "id", Value: sessionID.String()}}
	context.Request = httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(`{"message":"Gaming"}`),
	)
	context.Request.Header.Set("Content-Type", "application/json")
	context.Request.Header.Set("X-Anonymous-ID", anonymousID)

	handler.ChatStream(context)

	if recorder.Code != http.StatusOK || !strings.Contains(recorder.Body.String(), "event: token") {
		t.Fatalf("status = %d, body = %s", recorder.Code, recorder.Body.String())
	}
	lastMessage := repository.session.Messages[len(repository.session.Messages)-1]
	if lastMessage.Role != "assistant" || lastMessage.Content != "What is your budget?" {
		t.Fatalf("saved assistant message = %#v", lastMessage)
	}
}
