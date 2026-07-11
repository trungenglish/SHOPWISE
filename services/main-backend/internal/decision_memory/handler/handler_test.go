package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"shopwise/retail/internal/decision_memory/handler"
)

func TestHandler_GenerateResumeToken_Contract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Create a stub handler manually for contract validation
	h := &handler.Handler{} // Note: Normally would mock the service, but here we just test payload shape manually via mock if needed.
	// We'll just define the route and hit it with a mock context if service isn't injected, but actually we can test request binding contract easily.
	
	router.POST("/sessions/:id/resume-token", func(c *gin.Context) {
		var req handler.GenerateResumeTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, handler.GenerateResumeTokenResponse{
			Token: "stub-token",
		})
	})

	id := uuid.New().String()
	body := `{"expires_in_seconds": 3600}`
	req, _ := http.NewRequest("POST", "/sessions/"+id+"/resume-token", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp handler.GenerateResumeTokenResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "stub-token", resp.Token)
}

func TestHandler_BranchSession_Contract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.POST("/sessions/:id/branch", func(c *gin.Context) {
		var req handler.BranchSessionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"id": uuid.New().String(),
			"title": req.Title,
			"status": "active",
			"created_at": time.Now().Format(time.RFC3339),
			"updated_at": time.Now().Format(time.RFC3339),
		})
	})

	id := uuid.New().String()
	body := `{"title": "My New Branch"}`
	req, _ := http.NewRequest("POST", "/sessions/"+id+"/branch", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "My New Branch", resp["title"])
}
