package postgres

import (
	"time"

	"github.com/google/uuid"
)

type OrderModel struct {
	ID                uuid.UUID                `gorm:"type:uuid;primaryKey"`
	CustomerID        uuid.UUID                `gorm:"type:uuid;index;not null"`
	CustomerName      string                   `gorm:"size:255;not null;default:''"`
	CustomerEmail     string                   `gorm:"size:255;not null;default:''"`
	CustomerPhone     string                   `gorm:"size:32;not null;default:''"`
	FulfillmentMethod string                   `gorm:"size:32;not null;default:STORE_PICKUP"`
	ShippingAddress   string                   `gorm:"type:text;not null;default:''"`
	CouponCode        string                   `gorm:"size:50;not null;default:''"`
	SubtotalAmount    int64                    `gorm:"type:bigint;not null;default:0;check:subtotal_amount >= 0"`
	DiscountAmount    int64                    `gorm:"type:bigint;not null;default:0;check:discount_amount >= 0"`
	ShippingAmount    int64                    `gorm:"type:bigint;not null;default:0;check:shipping_amount >= 0"`
	TaxAmount         int64                    `gorm:"type:bigint;not null;default:0;check:tax_amount >= 0"`
	TotalAmount       int64                    `gorm:"type:bigint;not null;check:total_amount > 0"`
	Status            string                   `gorm:"size:32;not null;default:PENDING"`
	CreatedAt         time.Time                `gorm:"type:timestamptz;not null"`
	Items             []OrderItemModel         `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	RetailerItems     []RetailerOrderItemModel `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

type RetailerOrderItemModel struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID         uuid.UUID `gorm:"type:uuid;not null;index"`
	RetailerOfferID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name            string    `gorm:"size:255;not null"`
	Quantity        int       `gorm:"not null;check:quantity > 0"`
	UnitPrice       int64     `gorm:"type:bigint;not null;check:unit_price > 0"`
	SourceURL       string    `gorm:"type:text;not null"`
	VerifiedAt      time.Time `gorm:"type:timestamptz;not null"`
}

func (RetailerOrderItemModel) TableName() string { return "retailer_order_items" }

type CheckoutIdempotencyModel struct {
	CustomerID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	Key         string    `gorm:"size:255;primaryKey"`
	PayloadHash string    `gorm:"size:64;not null"`
	OrderID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null"`
}

func (CheckoutIdempotencyModel) TableName() string { return "checkout_idempotency" }

func (OrderModel) TableName() string {
	return "orders"
}

type OrderItemModel struct {
	OrderID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Quantity  int       `gorm:"not null;check:quantity > 0"`
	UnitPrice int64     `gorm:"type:bigint;not null;check:unit_price > 0"`
}

func (OrderItemModel) TableName() string {
	return "order_items"
}
