package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type greetingRequest struct {
	Locale      string `json:"locale" binding:"required,oneof=en vi"`
	DisplayName string `json:"display_name" binding:"max=100"`
}

func genericGreeting(locale, name string) string {
	if name == "" {
		name = "there"
	}
	if locale == "vi" {
		return "Chào " + name + "! 👋 Hôm nay mình có thể giúp gì cho bạn?"
	}
	return "Hey " + name + "! 👋 What can I help you with today?"
}

func (h *Handler) Greeting(c *gin.Context) {
	var request greetingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	name := strings.TrimSpace(request.DisplayName)
	uid, anonymousID := getIdentity(c)
	facts := []string{}
	if uid != nil || anonymousID != nil {
		if sessions, err := h.service.ListSessions(c.Request.Context(), uid, anonymousID, 1, 0); err == nil && len(sessions) > 0 {
			for _, message := range sessions[0].Messages {
				if message.Role == "user" && strings.TrimSpace(message.Content) != "" {
					facts = append(facts, message.Content)
				}
			}
		}
	}
	payload, _ := json.Marshal(gin.H{"locale": request.Locale, "display_name": name, "facts": facts})
	proxy, err := http.NewRequestWithContext(c.Request.Context(), http.MethodPost, h.aiRuntimeURL+"/api/v1/greeting", bytes.NewReader(payload))
	if err == nil {
		proxy.Header.Set("Content-Type", "application/json")
		if response, callErr := h.httpClient.Do(proxy); callErr == nil {
			defer response.Body.Close()
			body, _ := io.ReadAll(response.Body)
			var result struct{ Message string `json:"message"` }
			if response.StatusCode == http.StatusOK && json.Unmarshal(body, &result) == nil && strings.TrimSpace(result.Message) != "" {
				c.JSON(http.StatusOK, gin.H{"message": result.Message})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": genericGreeting(request.Locale, name)})
}
