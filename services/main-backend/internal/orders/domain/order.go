package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string
type FulfillmentMethod string

const (
	StatusPending    Status = "PENDING"
	StatusProcessing Status = "PROCESSING"

	FulfillmentDelivery    FulfillmentMethod = "DELIVERY"
	FulfillmentStorePickup FulfillmentMethod = "STORE_PICKUP"
)

type Item struct {
	ProductID uuid.UUID
	Quantity  int
	UnitPrice int64
}

type Order struct {
	ID                uuid.UUID
	CustomerID        uuid.UUID
	CustomerName      string
	CustomerEmail     string
	CustomerPhone     string
	FulfillmentMethod FulfillmentMethod
	ShippingAddress   string
	Items             []Item
	CouponCode        string
	SubtotalAmount    int64
	DiscountAmount    int64
	ShippingAmount    int64
	TaxAmount         int64
	TotalAmount       int64
	Status            Status
	CreatedAt         time.Time
}
