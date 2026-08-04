package postgres

import (
	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/platform/database/model"

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
	return database.Where(model.Promotion{CouponCode: "SHOPWISE5"}).
		Attrs(defaultPromotion).
		FirstOrCreate(&model.Promotion{}).Error
}
