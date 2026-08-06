package handler

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"shopwise/retail/internal/decision_memory/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	maxInteractionBody = 16 << 10
	maxFreeTextLength  = 1000
	interactionLease   = 120 * time.Second
)

type InteractionRequest struct {
	SchemaVersion string             `json:"schema_version" binding:"required"`
	InteractionID string             `json:"interaction_id" binding:"required"`
	SourceTurnID  string             `json:"source_turn_id" binding:"required"`
	ComponentID   string             `json:"component_id" binding:"required"`
	Event         string             `json:"event" binding:"required"`
	Action        string             `json:"action" binding:"required"`
	Payload       InteractionPayload `json:"payload" binding:"required"`
}

type InteractionPayload struct {
	QuestionID        string   `json:"question_id"`
	SelectedOptionIDs []string `json:"selected_option_ids"`
	FreeText          string   `json:"free_text"`
}

type storedQuestionEnvelope struct {
	Type     string `json:"type"`
	TurnID   string `json:"turn_id"`
	Question struct {
		ID      string `json:"id"`
		Mode    string `json:"mode"`
		Options []struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"options"`
	} `json:"question"`
	UIOperations []struct {
		Component interactionComponent `json:"component"`
	} `json:"ui_operations"`
}

type interactionComponent struct {
	ID       string                 `json:"id"`
	Children []interactionComponent `json:"children"`
}

func (h *Handler) InteractionStream(c *gin.Context) {
	sessionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}
	session, ok := h.getOwnedSession(c, sessionID)
	if !ok {
		return
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxInteractionBody)
	var request InteractionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid interaction request"})
		return
	}
	if request.SchemaVersion != "1.0" || request.Event != "submit" ||
		request.Action != "question.answer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported interaction"})
		return
	}
	interactionID, err := uuid.Parse(request.InteractionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid interaction id"})
		return
	}
	requestJSON, _ := json.Marshal(request)
	hash := sha256.Sum256(requestJSON)
	now := time.Now().UTC()
	interaction, acquired, err := h.service.BeginInteraction(
		c.Request.Context(),
		&domain.SessionInteraction{
			SessionID: sessionID, InteractionID: interactionID,
			PayloadHash: hex.EncodeToString(hash[:]), Status: "pending",
			LeaseUntil: now.Add(interactionLease), CreatedAt: now, UpdatedAt: now,
		},
	)
	if errors.Is(err, domain.ErrInteractionConflict) {
		c.JSON(http.StatusConflict, gin.H{"error": "interaction payload conflict"})
		return
	}
	if errors.Is(err, domain.ErrInteractionInProgress) {
		c.JSON(http.StatusConflict, gin.H{"error": "interaction in progress", "retryable": true})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to persist interaction"})
		return
	}
	if !acquired && interaction.ResponseEnvelope != "" {
		writeEnvelopeSSE(c, []byte(interaction.ResponseEnvelope))
		return
	}
	message, err := validateQuestionInteraction(session, request)
	if err != nil {
		_ = h.service.FinishInteraction(c, sessionID, interactionID, "failed", nil)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, ok := h.prepareAIChatPayload(c, sessionID, message)
	if !ok {
		_ = h.service.FinishInteraction(c, sessionID, interactionID, "failed", nil)
		return
	}
	responseBody, err := h.callAIRuntime(c, "/api/v1/chat", payload)
	if err != nil {
		_ = h.service.FinishInteraction(c, sessionID, interactionID, "failed", nil)
		c.JSON(http.StatusBadGateway, gin.H{"error": "AI Runtime returned an error"})
		return
	}
	var envelope agentEnvelope
	if json.Unmarshal(responseBody, &envelope) != nil || !validEnvelope(envelope) {
		_ = h.service.FinishInteraction(c, sessionID, interactionID, "failed", nil)
		c.JSON(http.StatusBadGateway, gin.H{"error": "AI Runtime returned an invalid response"})
		return
	}
	if err := h.persistAssistantMessage(c, sessionID, envelope.Message, responseBody); err != nil {
		_ = h.service.FinishInteraction(c, sessionID, interactionID, "failed", nil)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save AI response"})
		return
	}
	if err := h.service.FinishInteraction(
		c, sessionID, interactionID, "completed", responseBody,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to finish interaction"})
		return
	}
	writeEnvelopeSSE(c, responseBody)
}

func (h *Handler) callAIRuntime(c *gin.Context, path string, payload []byte) ([]byte, error) {
	request, err := http.NewRequestWithContext(
		c.Request.Context(), http.MethodPost, h.aiRuntimeURL+path, bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AI Runtime status %d", response.StatusCode)
	}
	return body, nil
}

func validateQuestionInteraction(
	session *domain.DecisionSession,
	request InteractionRequest,
) (string, error) {
	if request.SchemaVersion != "1.0" || request.Event != "submit" ||
		request.Action != "question.answer" {
		return "", errors.New("unsupported interaction")
	}
	if _, err := uuid.Parse(request.InteractionID); err != nil {
		return "", errors.New("invalid interaction id")
	}
	if len(request.Payload.FreeText) > maxFreeTextLength {
		return "", errors.New("free text is too long")
	}

	var envelope storedQuestionEnvelope
	found := false
	for index := len(session.Messages) - 1; index >= 0; index-- {
		message := session.Messages[index]
		if message.Role != "assistant" || message.ReasoningGraph == "" {
			continue
		}
		if json.Unmarshal([]byte(message.ReasoningGraph), &envelope) == nil {
			found = true
		}
		break
	}
	if !found || envelope.Type != "question" || envelope.Question.ID == "" {
		return "", errors.New("no active clarification question")
	}
	if request.InteractionID != envelope.Question.ID ||
		request.Payload.QuestionID != envelope.Question.ID ||
		request.SourceTurnID != envelope.TurnID {
		return "", errors.New("interaction does not match the active question")
	}
	if request.ComponentID != envelope.Question.ID &&
		!containsComponent(envelope.UIOperations, request.ComponentID) {
		return "", errors.New("unknown component")
	}
	if envelope.Question.Mode == "single" && len(request.Payload.SelectedOptionIDs) > 1 {
		return "", errors.New("question accepts one option")
	}

	labelsByID := make(map[string]string, len(envelope.Question.Options))
	for _, option := range envelope.Question.Options {
		labelsByID[option.ID] = option.Label
	}
	answers := make([]string, 0, len(request.Payload.SelectedOptionIDs)+1)
	seen := make(map[string]bool)
	for _, optionID := range request.Payload.SelectedOptionIDs {
		label, exists := labelsByID[optionID]
		if !exists || seen[optionID] {
			return "", errors.New("invalid selected option")
		}
		seen[optionID] = true
		answers = append(answers, label)
	}
	if freeText := strings.TrimSpace(request.Payload.FreeText); freeText != "" {
		answers = append(answers, freeText)
	}
	if len(answers) == 0 {
		return "", errors.New("an answer is required")
	}
	return "Answer to the clarification: " + strings.Join(answers, "; "), nil
}

func containsComponent(operations []struct {
	Component interactionComponent `json:"component"`
}, componentID string) bool {
	var contains func(interactionComponent) bool
	contains = func(component interactionComponent) bool {
		if component.ID == componentID {
			return true
		}
		for _, child := range component.Children {
			if contains(child) {
				return true
			}
		}
		return false
	}
	for _, operation := range operations {
		if contains(operation.Component) {
			return true
		}
	}
	return false
}

func writeEnvelopeSSE(c *gin.Context, envelopeJSON []byte) {
	var envelope struct {
		Message      string            `json:"message"`
		UIOperations []json.RawMessage `json:"ui_operations"`
	}
	_ = json.Unmarshal(envelopeJSON, &envelope)
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Status(http.StatusOK)
	token, _ := json.Marshal(gin.H{"text": envelope.Message})
	_, _ = fmt.Fprintf(c.Writer, "event: token\ndata: %s\n\n", token)
	for _, operation := range envelope.UIOperations {
		_, _ = fmt.Fprintf(c.Writer, "event: ui_operation\ndata: %s\n\n", operation)
	}
	_, _ = fmt.Fprintf(c.Writer, "event: envelope\ndata: %s\n\n", envelopeJSON)
	_, _ = fmt.Fprint(c.Writer, "event: done\ndata: {}\n\n")
}
