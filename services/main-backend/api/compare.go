package api

import (
	"net/http"
	"shopwise/retail/retail"

	"github.com/gin-gonic/gin"
)

type CompareAPI struct {
	provider *retail.MockRetailProvider
}

func NewCompareAPI(provider *retail.MockRetailProvider) *CompareAPI {
	return &CompareAPI{provider: provider}
}

type CompareRequest struct {
	ProductIDs []string `json:"product_ids"`
}

func (api *CompareAPI) Compare(c *gin.Context) {
	var req CompareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	results := api.provider.GetProducts(req.ProductIDs)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": results})
}
