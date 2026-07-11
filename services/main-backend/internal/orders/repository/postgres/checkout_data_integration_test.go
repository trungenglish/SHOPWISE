//go:build integration

package postgres_test

import (
	"context"
	"testing"
	"time"

	"shopwise/retail/internal/orders/domain"
	orderpostgres "shopwise/retail/internal/orders/repository/postgres"
	"shopwise/retail/internal/platform/database/model"
	integration "shopwise/retail/internal/platform/testutil/integration"
	userpostgres "shopwise/retail/internal/users/repository/postgres"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

func TestCheckoutDataReadsCustomerProductAndPromotion(t *testing.T) {
	database, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	customerID := uuid.New()
	if err := database.Create(&userpostgres.UserModel{
		ID: customerID, Email: "customer@example.com", Name: "Nguyen Van A", Phone: "0912345678",
	}).Error; err != nil {
		t.Fatalf("seed customer: %v", err)
	}
	productID := uuid.New()
	if err := database.Create(&model.Product{
		ID: productID, SKU: "LAPTOP-1", Name: "Laptop", Price: int64(87_475_000),
	}).Error; err != nil {
		t.Fatalf("seed product: %v", err)
	}
	if err := database.Create(&model.Inventory{
		ProductID: productID, Stock: 3, Availability: true,
	}).Error; err != nil {
		t.Fatalf("seed inventory: %v", err)
	}
	if err := database.Create(&model.Promotion{
		ID:              uuid.New(),
		Campaign:        "ShopWise 5",
		CouponCode:      "SHOPWISE5",
		Discount:        datatypes.JSON([]byte(`{}`)),
		DiscountType:    string(domain.DiscountPercentage),
		DiscountValue:   5,
		MinimumSubtotal: 0,
		Active:          true,
	}).Error; err != nil {
		t.Fatalf("seed promotion: %v", err)
	}

	repository := orderpostgres.NewRepository(database)
	customer, err := repository.GetCustomer(context.Background(), customerID)
	if err != nil {
		t.Fatalf("GetCustomer() error = %v", err)
	}
	if customer.Phone != "0912345678" {
		t.Fatalf("customer phone = %q", customer.Phone)
	}

	products, err := repository.GetProducts(context.Background(), []uuid.UUID{productID})
	if err != nil {
		t.Fatalf("GetProducts() error = %v", err)
	}
	if products[productID].UnitPrice != 87_475_000 || products[productID].Stock != 3 || !products[productID].Available {
		t.Fatalf("product quote = %#v", products[productID])
	}

	promotion, err := repository.GetPromotion(context.Background(), "shopwise5")
	if err != nil {
		t.Fatalf("GetPromotion() error = %v", err)
	}
	if promotion.DiscountType != domain.DiscountPercentage || promotion.DiscountValue != 5 {
		t.Fatalf("promotion = %#v", promotion)
	}
	if promotion.StartsAt != nil && promotion.StartsAt.After(time.Now().UTC()) {
		t.Fatal("seeded promotion unexpectedly starts in the future")
	}
}
