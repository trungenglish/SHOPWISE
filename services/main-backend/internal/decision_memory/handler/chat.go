package handler

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
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

func (h *Handler) ChatStream(c *gin.Context) {
	sessionIDStr := c.Param("id")
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	// Verify session exists
	_, err = h.service.GetSession(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add user message to DB
	userMsg := &domain.SessionMessage{
		ID:        uuid.New(),
		SessionID: sessionID,
		Role:      "user",
		Content:   req.Message,
		CreatedAt: time.Now().UTC(),
	}
	_ = h.service.AddMessage(c.Request.Context(), userMsg)

	// Call AI Runtime
	aiRuntimeURL := "http://localhost:8001/chat/stream" // hardcoded for now
	payload, _ := json.Marshal(map[string]interface{}{
		"session_id": sessionIDStr,
		"message":    req.Message,
	})

	proxyReq, err := http.NewRequestWithContext(c.Request.Context(), "POST", aiRuntimeURL, bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create proxy request"})
		return
	}
	proxyReq.Header.Set("Content-Type", "application/json")
	proxyReq.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to connect to ai-runtime"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "ai-runtime returned an error"})
		return
	}

	// Stream the response to the client
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	var fullResponseBuilder strings.Builder
	reader := bufio.NewReader(resp.Body)

	c.Stream(func(w io.Writer) bool {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return false
		}
		
		lineStr := string(line)
		
		// Parse SSE event and accumulate text
		if strings.HasPrefix(lineStr, "data: ") {
			dataStr := strings.TrimSpace(strings.TrimPrefix(lineStr, "data: "))
			if dataStr != "[DONE]" && dataStr != "" {
				var dataMap map[string]interface{}
				if err := json.Unmarshal([]byte(dataStr), &dataMap); err == nil {
					if text, ok := dataMap["text"].(string); ok {
						fullResponseBuilder.WriteString(text)
					}
				}
			}
		}

		_, writeErr := w.Write(line)
		return writeErr == nil
	})

	// Save the accumulated response to DB as assistant message
	fullResponse := fullResponseBuilder.String()
	if fullResponse != "" {
		assistantMsg := &domain.SessionMessage{
			ID:        uuid.New(),
			SessionID: sessionID,
			Role:      "assistant",
			Content:   fullResponse,
			CreatedAt: time.Now().UTC(),
		}
		_ = h.service.AddMessage(context.Background(), assistantMsg)
	}
}