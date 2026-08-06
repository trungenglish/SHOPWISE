package handler

import (
	"testing"
	"time"

	"shopwise/retail/internal/platform/database/model"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func TestBuildOfferComparisonUsesAuthoritativeBundleAndCampaignPrices(t *testing.T) {
	productID := uuid.MustParse("7cc26fb0-ace4-4e35-b6c5-fc3e13623546")
	product := model.Product{ID: productID, SKU: "ASUS-0001", Name: "ROG Strix G16", Price: 42_000_000}
	current := model.Promotion{
		Campaign:   "ROG Complete Gaming Bundle",
		CouponCode: "ROG5070BUNDLE",
		Discount: datatypes.JSON([]byte(`{
			"eligible_product_ids":["7cc26fb0-ace4-4e35-b6c5-fc3e13623546"],
			"lines":[
				{"retailer_offer_id":"10000000-0000-4000-8000-000000000001","name":"Gaming Mouse","original_price":1500000,"price":0,"required":true,"default_selected":true},
				{"retailer_offer_id":"10000000-0000-4000-8000-000000000002","name":"Anti-shock Bag","original_price":800000,"price":0,"required":true,"default_selected":true},
				{"retailer_offer_id":"10000000-0000-4000-8000-000000000003","name":"2-Year Accidental Damage Warranty","original_price":2000000,"price":1000000,"required":false,"default_selected":true}
			]
		}`)),
		Active: true,
	}
	startsAt := time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)
	endsAt := time.Date(2026, time.October, 1, 0, 0, 0, 0, time.UTC)
	future := model.Promotion{
		Campaign: "Back to School 2026", DiscountType: "PERCENTAGE", DiscountValue: 5,
		StartsAt: &startsAt, EndsAt: &endsAt, Active: true,
	}

	response, err := buildOfferComparison(product, current, future)

	if err != nil {
		t.Fatalf("buildOfferComparison() error = %v", err)
	}
	if response.BaseTotal != 42_000_000 || response.DefaultTotal != 43_000_000 {
		t.Fatalf("base/default total = %d/%d", response.BaseTotal, response.DefaultTotal)
	}
	if response.ScheduledCampaign.SalePrice != 39_900_000 || response.ScheduledCampaign.Savings != 2_100_000 {
		t.Fatalf("sale price/savings = %d/%d", response.ScheduledCampaign.SalePrice, response.ScheduledCampaign.Savings)
	}
}
