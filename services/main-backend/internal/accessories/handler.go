package accessories

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RecommendationService interface {
	Recommend(context.Context, []string) ([]Recommendation, error)
}

type Handler struct{ service RecommendationService }

func NewHandler(service RecommendationService) *Handler { return &Handler{service: service} }

type recommendationResponse struct {
	Items []recommendationItemResponse `json:"items"`
}

type recommendationItemResponse struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Kind                 string    `json:"kind"`
	Category             string    `json:"category"`
	Brand                string    `json:"brand"`
	ImageURL             string    `json:"image_url"`
	RetailerOfferID      string    `json:"retailer_offer_id"`
	RetailerProductID    string    `json:"retailer_product_id"`
	SourceURL            string    `json:"source_url"`
	Price                int64     `json:"price"`
	OriginalPrice        int64     `json:"original_price"`
	InStock              bool      `json:"in_stock"`
	FetchedAt            time.Time `json:"fetched_at"`
	OfferState           string    `json:"offer_state"`
	CheckoutAvailable    bool      `json:"checkout_available"`
	CompatibleProductIDs []string  `json:"compatible_product_ids"`
	CompatibilityReason  string    `json:"compatibility_reason"`
	CompatibilityStatus  string    `json:"compatibility_status"`
}

func (handler *Handler) Recommendations(ctx *gin.Context) {
	items, err := handler.service.Recommend(ctx.Request.Context(), ctx.QueryArray("product_ids"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	responseItems := make([]recommendationItemResponse, 0, len(items))
	for _, item := range items {
		productIDs := make([]string, 0, len(item.CompatibleProductIDs))
		for _, productID := range item.CompatibleProductIDs {
			productIDs = append(productIDs, productID.String())
		}
		responseItems = append(responseItems, recommendationItemResponse{
			ID: item.Accessory.ID.String(), Name: item.Accessory.Name, Kind: string(item.Accessory.Kind), Category: item.Accessory.Category,
			Brand: item.Accessory.Brand, ImageURL: item.Accessory.ImageURL, RetailerOfferID: item.Offer.ID.String(),
			RetailerProductID: item.Offer.RetailerProductID, SourceURL: item.Offer.SourceURL,
			Price: item.Offer.Price, OriginalPrice: item.Offer.OriginalPrice, InStock: item.Offer.InStock, FetchedAt: item.Offer.FetchedAt,
			OfferState: string(item.OfferState), CheckoutAvailable: item.CheckoutAvailable,
			CompatibleProductIDs: productIDs, CompatibilityReason: item.CompatibilityReason,
			CompatibilityStatus: string(item.CompatibilityStatus),
		})
	}
	ctx.JSON(http.StatusOK, recommendationResponse{Items: responseItems})
}

func RegisterRoutes(group *gin.RouterGroup, handler *Handler) {
	group.GET("/recommendations", handler.Recommendations)
}
