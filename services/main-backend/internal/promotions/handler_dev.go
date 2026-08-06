package promotions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type simulatePaymentRequest struct {
	SessionID string `json:"session_id" binding:"required"`
}

// SimulatePayment is a development-only endpoint to simulate a successful payment event
func (h *Handler) SimulatePayment(c *gin.Context) {
	var req simulatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.service.HandlePaymentCompleted(c.Request.Context(), req.SessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment simulation successful"})
}

func (h *Handler) RegisterDevRoutes(router *gin.RouterGroup) {
	group := router.Group("/checkout-incentive/dev")
	{
		group.POST("/simulate-payment", h.SimulatePayment)
	}
}
