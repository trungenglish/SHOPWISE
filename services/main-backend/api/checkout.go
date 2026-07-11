package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckoutAPI struct{}

func NewCheckoutAPI() *CheckoutAPI {
	return &CheckoutAPI{}
}

type CheckoutPrepareRequest struct {
	Items []string `json:"items"`
}

func (api *CheckoutAPI) Prepare(c *gin.Context) {
	var req CheckoutPrepareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Mock checkout response
	checkoutID := fmt.Sprintf("chk_%d", len(req.Items))

	data := map[string]interface{}{
		"checkoutId": checkoutID,
		"items":      req.Items,
		"totalPrice": 25000000.0, // Mock price in VND
		"status":     "PENDING",
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
