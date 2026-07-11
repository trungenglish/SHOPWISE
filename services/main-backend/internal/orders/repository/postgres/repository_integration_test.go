//go:build integration

package postgres_test

import (
	"context"
	"testing"
	"time"

	"shopwise/retail/internal/orders/domain"
	orderpostgres "shopwise/retail/internal/orders/repository/postgres"
	integration "shopwise/retail/internal/platform/testutil/integration"

	"github.com/google/uuid"
)

func TestRepositoryCreatePersistsOrderAndItems(t *testing.T) {
	database, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	repository := orderpostgres.NewRepository(database)
	order := &domain.Order{
		ID:                uuid.New(),
		CustomerID:        uuid.New(),
		CustomerName:      "Nguyen Van A",
		CustomerEmail:     "customer@example.com",
		CustomerPhone:     "0912345678",
		FulfillmentMethod: domain.FulfillmentDelivery,
		ShippingAddress:   "1 Nguyen Hue",
		CreatedAt:         time.Now().UTC().Truncate(time.Microsecond),
		SubtotalAmount:    3250,
		DiscountAmount:    100,
		ShippingAmount:    0,
		TaxAmount:         252,
		TotalAmount:       3402,
		CouponCode:        "SHOPWISE5",
		Status:            domain.StatusPending,
		Items: []domain.Item{
			{ProductID: uuid.New(), Quantity: 2, UnitPrice: 1250},
			{ProductID: uuid.New(), Quantity: 1, UnitPrice: 750},
		},
	}

	if err := repository.Create(context.Background(), order); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	var persistedOrder orderpostgres.OrderModel
	if err := database.First(&persistedOrder, "id = ?", order.ID).Error; err != nil {
		t.Fatalf("find persisted order: %v", err)
	}
	if persistedOrder.CustomerID != order.CustomerID {
		t.Fatalf("CustomerID = %s, want %s", persistedOrder.CustomerID, order.CustomerID)
	}
	if persistedOrder.TotalAmount != order.TotalAmount {
		t.Fatalf("TotalAmount = %d, want %d", persistedOrder.TotalAmount, order.TotalAmount)
	}
	if persistedOrder.CustomerName != order.CustomerName || persistedOrder.CustomerPhone != order.CustomerPhone {
		t.Fatalf("customer snapshot = %q/%q", persistedOrder.CustomerName, persistedOrder.CustomerPhone)
	}
	if persistedOrder.SubtotalAmount != order.SubtotalAmount || persistedOrder.TaxAmount != order.TaxAmount {
		t.Fatalf("subtotal/tax = %d/%d", persistedOrder.SubtotalAmount, persistedOrder.TaxAmount)
	}
	if persistedOrder.Status != string(domain.StatusPending) {
		t.Fatalf("Status = %q, want %q", persistedOrder.Status, domain.StatusPending)
	}

	var itemCount int64
	if err := database.Model(&orderpostgres.OrderItemModel{}).
		Where("order_id = ?", order.ID).
		Count(&itemCount).Error; err != nil {
		t.Fatalf("count persisted items: %v", err)
	}
	if itemCount != int64(len(order.Items)) {
		t.Fatalf("item count = %d, want %d", itemCount, len(order.Items))
	}
}

func TestRepositoryCreateRollsBackOrderWhenAnItemFails(t *testing.T) {
	database, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	productID := uuid.New()
	order := &domain.Order{
		ID:                uuid.New(),
		CustomerID:        uuid.New(),
		CustomerName:      "Nguyen Van A",
		CustomerEmail:     "customer@example.com",
		CustomerPhone:     "0912345678",
		FulfillmentMethod: domain.FulfillmentStorePickup,
		CreatedAt:         time.Now().UTC(),
		SubtotalAmount:    200,
		TaxAmount:         16,
		TotalAmount:       216,
		Status:            domain.StatusPending,
		Items: []domain.Item{
			{ProductID: productID, Quantity: 1, UnitPrice: 100},
			{ProductID: productID, Quantity: 1, UnitPrice: 100},
		},
	}

	repository := orderpostgres.NewRepository(database)
	if err := repository.Create(context.Background(), order); err == nil {
		t.Fatal("Create() error = nil, want duplicate item error")
	}

	var orderCount int64
	if err := database.Model(&orderpostgres.OrderModel{}).
		Where("id = ?", order.ID).
		Count(&orderCount).Error; err != nil {
		t.Fatalf("count orders after rollback: %v", err)
	}
	if orderCount != 0 {
		t.Fatalf("order count after rollback = %d, want 0", orderCount)
	}
}

func TestRepositoryListByCustomerReturnsNewestOrdersWithItems(t *testing.T) {
	database, cleanup := integration.SetupIntegrationDB(t)
	defer cleanup()

	repository := orderpostgres.NewRepository(database)
	customerID := uuid.New()
	older := integrationOrder(customerID, time.Now().UTC().Add(-time.Hour))
	newer := integrationOrder(customerID, time.Now().UTC())
	otherCustomer := integrationOrder(uuid.New(), time.Now().UTC().Add(time.Hour))
	for _, order := range []*domain.Order{older, newer, otherCustomer} {
		if err := repository.Create(context.Background(), order); err != nil {
			t.Fatalf("Create() error = %v", err)
		}
	}

	orders, err := repository.ListByCustomer(context.Background(), customerID, 20, 0)
	if err != nil {
		t.Fatalf("ListByCustomer() error = %v", err)
	}
	if len(orders) != 2 {
		t.Fatalf("len(orders) = %d, want 2", len(orders))
	}
	if orders[0].ID != newer.ID || orders[1].ID != older.ID {
		t.Fatalf("order IDs = %s/%s, want newest first", orders[0].ID, orders[1].ID)
	}
	if len(orders[0].Items) != 1 {
		t.Fatalf("len(newest.Items) = %d, want 1", len(orders[0].Items))
	}
}

func integrationOrder(customerID uuid.UUID, createdAt time.Time) *domain.Order {
	return &domain.Order{
		ID:                uuid.New(),
		CustomerID:        customerID,
		CustomerName:      "Nguyen Van A",
		CustomerEmail:     "customer@example.com",
		CustomerPhone:     "0912345678",
		FulfillmentMethod: domain.FulfillmentStorePickup,
		SubtotalAmount:    100,
		TaxAmount:         8,
		TotalAmount:       108,
		Status:            domain.StatusPending,
		CreatedAt:         createdAt,
		Items:             []domain.Item{{ProductID: uuid.New(), Quantity: 1, UnitPrice: 100}},
	}
}
