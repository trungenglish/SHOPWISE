package api

import (
	"net/http"
	"shopwise/retail/retail"
	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/decision_memory/usecase"

	"github.com/gin-gonic/gin"
)

type CompareAPI struct {
	provider *retail.MockRetailProvider
	uc       usecase.ComparisonUsecase
}

func NewCompareAPI(provider *retail.MockRetailProvider, uc usecase.ComparisonUsecase) *CompareAPI {
	return &CompareAPI{provider: provider, uc: uc}
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

type AddToComparisonRequest struct {
	WorkspaceID string `json:"workspace_id"`
	ProductID   string `json:"product_id"`
}

func (api *CompareAPI) AddToComparison(c *gin.Context) {
	var req AddToComparisonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	products := api.provider.GetProducts([]string{req.ProductID})
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "product not found"})
		return
	}
	
	p := products[0]
	avail := domain.AvailabilityOutOfStock
	if p.InStock {
		avail = domain.AvailabilityInStock
	}

	detail := &domain.ComparisonProductDetail{
		ProductID:      p.ProductID,
		Name:           p.Name,
		Brand:          p.Brand,
		Category:       p.Category,
		PriceVND:       int64(p.Price),
		Availability:   avail,
		Images:         p.Images,
		Specifications: p.Specs,
	}

	ws, err := api.uc.AddProductToWorkspace(c.Request.Context(), req.WorkspaceID, detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"success": true, "data": ws})
}

type ReplaceProductRequest struct {
	WorkspaceID  string `json:"workspace_id"`
	OldProductID string `json:"old_product_id"`
	NewProductID string `json:"new_product_id"`
}

func (api *CompareAPI) ReplaceProduct(c *gin.Context) {
	var req ReplaceProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	products := api.provider.GetProducts([]string{req.NewProductID})
	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "new product not found"})
		return
	}

	p := products[0]
	avail := domain.AvailabilityOutOfStock
	if p.InStock {
		avail = domain.AvailabilityInStock
	}

	detail := &domain.ComparisonProductDetail{
		ProductID:      p.ProductID,
		Name:           p.Name,
		Brand:          p.Brand,
		Category:       p.Category,
		PriceVND:       int64(p.Price),
		Availability:   avail,
		Images:         p.Images,
		Specifications: p.Specs,
	}

	ws, err := api.uc.ReplaceProductInWorkspace(c.Request.Context(), req.WorkspaceID, req.OldProductID, detail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"success": true, "data": ws})
}

type RemoveProductRequest struct {
	WorkspaceID string `json:"workspace_id"`
	ProductID   string `json:"product_id"`
}

func (api *CompareAPI) RemoveProduct(c *gin.Context) {
	var req RemoveProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	ws, err := api.uc.RemoveProductFromWorkspace(c.Request.Context(), req.WorkspaceID, req.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"success": true, "data": ws})
}
