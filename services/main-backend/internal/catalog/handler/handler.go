package handler

import (
	"net/http"

	"shopwise/retail/internal/platform/database/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) ListProducts(c *gin.Context) {
	var products []model.Product
	if err := h.db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": products,
		"limit": len(products),
		"offset": 0,
	})
}

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("", h.ListProducts)
	rg.GET("/:id/offer-comparison", h.GetOfferComparison)
}
