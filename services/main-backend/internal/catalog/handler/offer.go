package handler

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"time"

	"shopwise/retail/internal/platform/database/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	bundleCouponCode   = "ROG5070BUNDLE"
	campaignCouponCode = "BACKTOSCHOOL2026"
)

type bundleLine struct {
	RetailerOfferID string `json:"retailer_offer_id"`
	Name            string `json:"name"`
	OriginalPrice   int64  `json:"original_price"`
	Price           int64  `json:"price"`
	Required        bool   `json:"required"`
	DefaultSelected bool   `json:"default_selected"`
}

type bundleDefinition struct {
	EligibleProductIDs []string     `json:"eligible_product_ids"`
	Lines              []bundleLine `json:"lines"`
}

type scheduledCampaignResponse struct {
	Name      string    `json:"name"`
	StartsAt  time.Time `json:"starts_at"`
	EndsAt    time.Time `json:"ends_at"`
	Discount  int64     `json:"discount_percentage"`
	SalePrice int64     `json:"sale_price"`
	Savings   int64     `json:"savings"`
}

type offerComparisonResponse struct {
	OfferID           string                    `json:"offer_id"`
	Product           model.Product             `json:"product"`
	BundleName        string                    `json:"bundle_name"`
	Lines             []bundleLine              `json:"lines"`
	BaseTotal         int64                     `json:"base_total"`
	DefaultTotal      int64                     `json:"default_total"`
	ScheduledCampaign scheduledCampaignResponse `json:"scheduled_campaign"`
}

func buildOfferComparison(product model.Product, current model.Promotion, future model.Promotion) (offerComparisonResponse, error) {
	var bundle bundleDefinition
	if err := json.Unmarshal(current.Discount, &bundle); err != nil {
		return offerComparisonResponse{}, errors.New("invalid bundle definition")
	}

	eligible := false
	for _, productID := range bundle.EligibleProductIDs {
		if productID == product.ID.String() {
			eligible = true
			break
		}
	}
	if !eligible || future.StartsAt == nil || future.EndsAt == nil {
		return offerComparisonResponse{}, errors.New("offer is not available for this product")
	}

	defaultTotal := product.Price
	for _, line := range bundle.Lines {
		if line.DefaultSelected {
			defaultTotal += line.Price
		}
	}
	savings := int64(math.Round(float64(product.Price) * float64(future.DiscountValue) / 100))

	return offerComparisonResponse{
		OfferID:      current.ID.String(),
		Product:      product,
		BundleName:   current.Campaign,
		Lines:        bundle.Lines,
		BaseTotal:    product.Price,
		DefaultTotal: defaultTotal,
		ScheduledCampaign: scheduledCampaignResponse{
			Name:      future.Campaign,
			StartsAt:  *future.StartsAt,
			EndsAt:    future.EndsAt.Add(-time.Nanosecond),
			Discount:  future.DiscountValue,
			SalePrice: product.Price - savings,
			Savings:   savings,
		},
	}, nil
}

func (h *Handler) GetOfferComparison(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var product model.Product
	if err = h.db.First(&product, "id = ?", productID).Error; err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": "product offer not found"})
		return
	}

	var promotions []model.Promotion
	if err = h.db.Where("coupon_code IN ?", []string{bundleCouponCode, campaignCouponCode}).Find(&promotions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load offers"})
		return
	}
	byCode := make(map[string]model.Promotion, len(promotions))
	for _, promotion := range promotions {
		byCode[promotion.CouponCode] = promotion
	}
	response, err := buildOfferComparison(product, byCode[bundleCouponCode], byCode[campaignCouponCode])
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
