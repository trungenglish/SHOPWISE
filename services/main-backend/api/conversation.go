package api

import (
	"net/http"
	"shopwise/retail/session"

	"github.com/gin-gonic/gin"
)

type ConversationAPI struct {
	manager *session.Manager
}

func NewConversationAPI(manager *session.Manager) *ConversationAPI {
	return &ConversationAPI{manager: manager}
}

type MessageRequest struct {
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}

func (api *ConversationAPI) SendMessage(c *gin.Context) {
	var req MessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Fetch or create session
	sess, err := api.manager.GetSession(req.SessionID)
	if err != nil {
		sess = api.manager.CreateSession(req.SessionID)
	}

	api.manager.UpdateSession(req.SessionID, func(s *session.CustomerSession) {
		s.Messages = append(s.Messages, map[string]interface{}{
			"role":    "user",
			"content": req.Message,
		})
	})

	// TODO: Actually call AI Runtime and stream back the response.
	// For now, return a placeholder response.
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": map[string]interface{}{
			"response": "Acknowledged. This is a placeholder from the Go backend.",
		},
	})
}

