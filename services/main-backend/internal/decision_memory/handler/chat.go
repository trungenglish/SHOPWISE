package handler

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"shopwise/retail/internal/decision_memory/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatRequest struct {
	Message string `json:"message" binding:"required"`
}

type aiChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type aiChatRequest struct {
	SessionID            string          `json:"session_id"`
	Messages             []aiChatMessage `json:"messages"`
	AllowedComparisonIDs []string        `json:"allowed_comparison_ids"`
}

type agentEnvelope struct {
	Type     string          `json:"type"`
	Message  string          `json:"message"`
	Decision json.RawMessage `json:"decision,omitempty"`
}

func (h *Handler) Chat(c *gin.Context) {
	sessionID, payload, ok := h.prepareChatRequest(c)
	if !ok {
		return
	}
	proxyRequest, err := http.NewRequestWithContext(
		c.Request.Context(),
		http.MethodPost,
		h.aiRuntimeURL+"/api/v1/chat",
		bytes.NewReader(payload),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create AI request"})
		return
	}
	proxyRequest.Header.Set("Content-Type", "application/json")

	response, err := h.httpClient.Do(proxyRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to connect to AI Runtime"})
		return
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to read AI response"})
		return
	}
	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "AI Runtime returned an error"})
		return
	}

	var envelope agentEnvelope
	if err := json.Unmarshal(responseBody, &envelope); err != nil || !validEnvelope(envelope) {
		c.JSON(http.StatusBadGateway, gin.H{"error": "AI Runtime returned an invalid response"})
		return
	}
	if err := h.persistAssistantMessage(c, sessionID, envelope.Message, responseBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save AI response"})
		return
	}

	c.Data(http.StatusOK, "application/json", responseBody)
}

func (h *Handler) prepareChatRequest(c *gin.Context) (uuid.UUID, []byte, bool) {
	sessionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return uuid.Nil, nil, false
	}

	_, ok := h.getOwnedSession(c, sessionID)
	if !ok {
		return uuid.Nil, nil, false
	}

	var request ChatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return uuid.Nil, nil, false
	}
	payload, ok := h.prepareAIChatPayload(c, sessionID, request.Message)
	return sessionID, payload, ok
}

func (h *Handler) prepareAIChatPayload(
	c *gin.Context,
	sessionID uuid.UUID,
	message string,
) ([]byte, bool) {
	userMessage := &domain.SessionMessage{
		ID:        uuid.New(),
		SessionID: sessionID,
		Role:      "user",
		Content:   message,
		CreatedAt: time.Now().UTC(),
	}
	if err := h.service.AddMessage(c.Request.Context(), userMessage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save message"})
		return nil, false
	}

	session, err := h.service.GetSession(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load session history"})
		return nil, false
	}
	messages := make([]aiChatMessage, 0, len(session.Messages))
	for _, message := range session.Messages {
		if message.Role == "user" || message.Role == "assistant" {
			messages = append(messages, aiChatMessage{Role: message.Role, Content: message.Content})
		}
	}

	payload, err := json.Marshal(aiChatRequest{
		SessionID:            sessionID.String(),
		Messages:             messages,
		AllowedComparisonIDs: sessionProductIDs(session.Messages),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to encode AI request"})
		return nil, false
	}
	return payload, true
}

func sessionProductIDs(messages []domain.SessionMessage) []string {
	seen := make(map[string]bool)
	productIDs := make([]string, 0)
	for _, message := range messages {
		if message.Role != "assistant" || message.ReasoningGraph == "" {
			continue
		}
		var stored struct {
			Decision struct {
				Products []struct {
					ID string `json:"id"`
				} `json:"products"`
			} `json:"decision"`
		}
		if json.Unmarshal([]byte(message.ReasoningGraph), &stored) != nil {
			continue
		}
		for _, product := range stored.Decision.Products {
			if product.ID != "" && !seen[product.ID] {
				seen[product.ID] = true
				productIDs = append(productIDs, product.ID)
			}
		}
	}
	return productIDs
}

func (h *Handler) persistAssistantMessage(
	c *gin.Context,
	sessionID uuid.UUID,
	message string,
	envelope []byte,
) error {
	assistantMessage := &domain.SessionMessage{
		ID:             uuid.New(),
		SessionID:      sessionID,
		Role:           "assistant",
		Content:        message,
		ReasoningGraph: string(envelope),
		CreatedAt:      time.Now().UTC(),
	}
	return h.service.AddMessage(c.Request.Context(), assistantMessage)
}

func (h *Handler) ChatStream(c *gin.Context) {
	sessionID, payload, ok := h.prepareChatRequest(c)
	if !ok {
		return
	}
	proxyRequest, err := http.NewRequestWithContext(
		c.Request.Context(),
		http.MethodPost,
		h.aiRuntimeURL+"/api/v1/chat/stream",
		bytes.NewReader(payload),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create AI request"})
		return
	}
	proxyRequest.Header.Set("Content-Type", "application/json")
	response, err := h.httpClient.Do(proxyRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to connect to AI Runtime"})
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"error": "AI Runtime returned an error"})
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Status(http.StatusOK)
	flusher, _ := c.Writer.(http.Flusher)
	writeBlock := func(lines []string) {
		for _, line := range lines {
			_, _ = fmt.Fprintf(c.Writer, "%s\n", line)
		}
		_, _ = fmt.Fprint(c.Writer, "\n")
		if flusher != nil {
			flusher.Flush()
		}
	}

	var message strings.Builder
	var decision json.RawMessage
	var fullEnvelope json.RawMessage
	var decisionEvent string
	failed := false
	scanner := bufio.NewScanner(response.Body)
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)
	block := make([]string, 0, 2)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			block = append(block, line)
			continue
		}
		event, data := parseSSEBlock(block)
		switch event {
		case "token":
			var token struct {
				Text string `json:"text"`
			}
			if json.Unmarshal([]byte(data), &token) == nil {
				message.WriteString(token.Text)
			}
		case "recommendation", "comparison", "offer_comparison", "checkout_ready":
			decision = json.RawMessage(data)
			decisionEvent = event
		case "envelope":
			fullEnvelope = json.RawMessage(data)
		case "error":
			failed = true
		case "done":
			if !failed && message.Len() > 0 {
				if len(fullEnvelope) > 0 {
					var stored agentEnvelope
					if json.Unmarshal(fullEnvelope, &stored) == nil && validEnvelope(stored) {
						if h.persistAssistantMessage(
							c, sessionID, stored.Message, fullEnvelope,
						) != nil {
							writeBlock([]string{"event: error", `data: {"message":"failed to save AI response"}`})
						}
						break
					}
				}
				envelope := agentEnvelope{Type: "question", Message: message.String()}
				if len(decision) > 0 {
					envelope.Type = decisionEvent
					envelope.Decision = decision
				}
				envelopeJSON, marshalErr := json.Marshal(envelope)
				if marshalErr != nil || h.persistAssistantMessage(
					c, sessionID, envelope.Message, envelopeJSON,
				) != nil {
					writeBlock([]string{"event: error", `data: {"message":"failed to save AI response"}`})
				}
			}
		}
		writeBlock(block)
		block = block[:0]
	}
	if len(block) > 0 {
		writeBlock(block)
	}
}

func parseSSEBlock(lines []string) (string, string) {
	var event string
	var data string
	for _, line := range lines {
		if value, ok := strings.CutPrefix(line, "event: "); ok {
			event = value
		}
		if value, ok := strings.CutPrefix(line, "data: "); ok {
			data = value
		}
	}
	return event, data
}

func ownsSession(session *domain.DecisionSession, userID *uuid.UUID, anonymousID *string) bool {
	if userID != nil {
		return session.UserID != nil && *session.UserID == *userID
	}
	return anonymousID != nil && session.UserID == nil && session.AnonymousID != nil &&
		*session.AnonymousID == *anonymousID
}

func validEnvelope(envelope agentEnvelope) bool {
	if envelope.Message == "" {
		return false
	}
	if envelope.Type == "question" {
		return true
	}
	return (envelope.Type == "recommendation" || envelope.Type == "comparison" || envelope.Type == "offer_comparison" ||
		envelope.Type == "checkout_ready") &&
		len(envelope.Decision) > 0
}
