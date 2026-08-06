package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/decision_memory/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	testQuestionID = "11111111-1111-4111-8111-111111111111"
	testTurnID     = "22222222-2222-4222-8222-222222222222"
	testButtonID   = "33333333-3333-4333-8333-333333333333"
)

func activeQuestionSession() *domain.DecisionSession {
	return &domain.DecisionSession{Messages: []domain.SessionMessage{{
		Role: "assistant",
		ReasoningGraph: `{
          "type":"question",
          "turn_id":"` + testTurnID + `",
          "question":{"id":"` + testQuestionID + `","mode":"single","options":[
            {"id":"under-25m","label":"Under 25 million VND"},
            {"id":"25m-40m","label":"25–40 million VND"},
            {"id":"flexible","label":"Flexible budget"}
          ]},
          "ui_operations":[{"component":{"id":"` + testQuestionID + `","children":[
            {"id":"` + testButtonID + `","children":[]}
          ]}}]
        }`,
	}}}
}

func TestValidateQuestionInteraction(t *testing.T) {
	request := InteractionRequest{
		SchemaVersion: "1.0", InteractionID: testQuestionID,
		SourceTurnID: testTurnID, ComponentID: testButtonID,
		Event: "submit", Action: "question.answer",
		Payload: InteractionPayload{
			QuestionID:        testQuestionID,
			SelectedOptionIDs: []string{"under-25m"},
		},
	}

	message, err := validateQuestionInteraction(activeQuestionSession(), request)

	require.NoError(t, err)
	require.Contains(t, message, "Under 25 million VND")
}

func TestValidateQuestionInteractionRejectsUnknownOption(t *testing.T) {
	request := InteractionRequest{
		SchemaVersion: "1.0", InteractionID: testQuestionID,
		SourceTurnID: testTurnID, ComponentID: testButtonID,
		Event: "submit", Action: "question.answer",
		Payload: InteractionPayload{
			QuestionID:        testQuestionID,
			SelectedOptionIDs: []string{"invented"},
		},
	}

	_, err := validateQuestionInteraction(activeQuestionSession(), request)

	require.ErrorContains(t, err, "invalid selected option")
}

func TestInteractionRejectsSessionOwnedByAnotherAnonymousUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	sessionID := uuid.New()
	ownerID := "owner"
	repository := &chatRepository{session: domain.DecisionSession{
		ID: sessionID, AnonymousID: &ownerID,
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
		strings.NewReader(`{"schema_version":"1.0"}`),
	)
	context.Request.Header.Set("Content-Type", "application/json")
	context.Request.Header.Set("X-Anonymous-ID", "other-user")

	handler.InteractionStream(context)

	require.Equal(t, http.StatusNotFound, recorder.Code)
	require.False(t, aiCalled)
}
