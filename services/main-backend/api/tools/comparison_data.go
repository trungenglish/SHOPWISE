package tools

import (
	"net/http"
	"shopwise/retail/retail"
	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/decision_memory/usecase"

	"github.com/gin-gonic/gin"
)

type ComparisonDataAPI struct {
	provider *retail.MockRetailProvider
	uc       usecase.ComparisonUsecase
}

func NewComparisonDataAPI(provider *retail.MockRetailProvider, uc usecase.ComparisonUsecase) *ComparisonDataAPI {
	return &ComparisonDataAPI{provider: provider, uc: uc}
}

type FetchComparisonDataRequest struct {
	ProductIDs []string `json:"product_ids"`
}

func (api *ComparisonDataAPI) FetchComparisonData(c *gin.Context) {
	var req FetchComparisonDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	products := api.provider.GetProducts(req.ProductIDs)
	
	details := make([]*domain.ComparisonProductDetail, 0, len(products))
	for _, p := range products {
		avail := domain.AvailabilityOutOfStock
		if p.InStock {
			avail = domain.AvailabilityInStock
		}
		details = append(details, &domain.ComparisonProductDetail{
			ProductID:      p.ProductID,
			Name:           p.Name,
			Brand:          p.Brand,
			Category:       p.Category,
			PriceVND:       int64(p.Price),
			Availability:   avail,
			Images:         p.Images,
			Specifications: p.Specs,
		})
	}

	highlights := api.uc.ComputeHighlights(c.Request.Context(), details)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"products":   details,
			"highlights": highlights,
		},
	})
}
