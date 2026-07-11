package api

import (
	"net/http"
	"shopwise/retail/internal/decision_memory/usecase"

	"github.com/gin-gonic/gin"
)

type CheckoutReadinessAPI struct {
	uc usecase.ComparisonUsecase
}

func NewCheckoutReadinessAPI(uc usecase.ComparisonUsecase) *CheckoutReadinessAPI {
	return &CheckoutReadinessAPI{uc: uc}
}

type CheckoutReadyRequest struct {
	WorkspaceID string `json:"workspace_id"`
	ProductID   string `json:"product_id"`
}

func (api *CheckoutReadinessAPI) MarkCheckoutReady(c *gin.Context) {
	var req CheckoutReadyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	ws, err := api.uc.MarkCheckoutReady(c.Request.Context(), req.WorkspaceID, req.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ws,
	})
}
