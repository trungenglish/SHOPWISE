package accessories

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type recommendationServiceStub struct {
	items []Recommendation
	err   error
}

func (stub recommendationServiceStub) Recommend(context.Context, []string) ([]Recommendation, error) {
	return stub.items, stub.err
}

func TestRecommendationHandlerReturnsPhongVuOfferMetadata(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	productID := uuid.New()
	fetchedAt := time.Date(2026, time.August, 5, 12, 0, 0, 0, time.UTC)
	handler := NewHandler(recommendationServiceStub{items: []Recommendation{{
		Accessory:            Accessory{ID: uuid.New(), Kind: KindExternal, Category: "mouse", Name: "Logitech mouse"},
		Offer:                RetailerOffer{ID: uuid.New(), RetailerProductID: "PV-1", SourceURL: "https://phongvu.vn/p/logitech-mouse", Price: 890_000, InStock: true, FetchedAt: fetchedAt},
		CompatibleProductIDs: []uuid.UUID{productID}, CompatibilityReason: "Suitable for gaming",
		CompatibilityStatus: VerificationCategory, OfferState: OfferFresh, CheckoutAvailable: true,
	}}})
	router := gin.New()
	RegisterRoutes(router.Group("/api/v1/accessories"), handler)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/accessories/recommendations?product_ids="+productID.String(), nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body = %s", response.Code, response.Body.String())
	}
	var body struct {
		Items []struct {
			SourceURL            string   `json:"source_url"`
			OfferState           string   `json:"offer_state"`
			CompatibleProductIDs []string `json:"compatible_product_ids"`
		} `json:"items"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(body.Items) != 1 || body.Items[0].SourceURL != "https://phongvu.vn/p/logitech-mouse" || body.Items[0].OfferState != "fresh" {
		t.Fatalf("response = %#v", body)
	}
	if len(body.Items[0].CompatibleProductIDs) != 1 || body.Items[0].CompatibleProductIDs[0] != productID.String() {
		t.Fatalf("compatible ids = %v", body.Items[0].CompatibleProductIDs)
	}
}
