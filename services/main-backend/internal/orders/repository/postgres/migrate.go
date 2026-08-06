package postgres

import (
	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/platform/database/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	if err := database.AutoMigrate(
		&model.Product{},
		&model.Inventory{},
		&model.Promotion{},
		&OrderModel{},
		&OrderItemModel{},
		&RetailerOrderItemModel{},
		&CheckoutIdempotencyModel{},
	); err != nil {
		return err
	}

	defaultPromotion := model.Promotion{
		ID:              uuid.New(),
		Campaign:        "ShopWise 5 Percent",
		CouponCode:      "SHOPWISE5",
		Discount:        datatypes.JSON([]byte(`{}`)),
		DiscountType:    string(domain.DiscountPercentage),
		DiscountValue:   5,
		MinimumSubtotal: 0,
		Active:          true,
	}
	if err := database.Where(model.Promotion{CouponCode: "SHOPWISE5"}).
		Attrs(defaultPromotion).
		FirstOrCreate(&model.Promotion{}).Error; err != nil {
		return err
	}

	bundle := model.Promotion{
		ID: uuid.MustParse("20000000-0000-4000-8000-000000000001"), Campaign: "ROG Complete Gaming Bundle",
		CouponCode: "ROG5070BUNDLE", Discount: datatypes.JSON([]byte(`{"eligible_product_ids":["7cc26fb0-ace4-4e35-b6c5-fc3e13623546"],"lines":[{"retailer_offer_id":"10000000-0000-4000-8000-000000000001","name":"Gaming Mouse","original_price":1500000,"price":0,"required":true,"default_selected":true},{"retailer_offer_id":"10000000-0000-4000-8000-000000000002","name":"Anti-shock Bag","original_price":800000,"price":0,"required":true,"default_selected":true},{"retailer_offer_id":"10000000-0000-4000-8000-000000000003","name":"2-Year Accidental Damage Warranty","original_price":2000000,"price":1000000,"required":false,"default_selected":true}]}`)),
		DiscountType: string(domain.DiscountFixedAmount), DiscountValue: 0, Active: true,
	}
	if err := database.Where(model.Promotion{CouponCode: bundle.CouponCode}).Attrs(bundle).FirstOrCreate(&model.Promotion{}).Error; err != nil {
		return err
	}

	startsAt := time.Date(2026, time.September, 1, 0, 0, 0, 0, time.UTC)
	endsAt := time.Date(2026, time.October, 1, 0, 0, 0, 0, time.UTC)
	campaign := model.Promotion{
		ID: uuid.MustParse("20000000-0000-4000-8000-000000000002"), Campaign: "Back to School 2026",
		CouponCode: "BACKTOSCHOOL2026", Discount: datatypes.JSON([]byte(`{}`)),
		DiscountType: string(domain.DiscountPercentage), DiscountValue: 5, Active: true,
		StartsAt: &startsAt, EndsAt: &endsAt,
	}
	return database.Where(model.Promotion{CouponCode: campaign.CouponCode}).Attrs(campaign).FirstOrCreate(&model.Promotion{}).Error
}
