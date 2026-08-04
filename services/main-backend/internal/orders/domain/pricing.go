package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound           = errors.New("checkout customer not found")
	ErrPromotionNotFound          = errors.New("promotion not found")
	ErrIdempotencyPayloadConflict = errors.New("idempotency key was reused with a different payload")
)

type CustomerSnapshot struct {
	ID    uuid.UUID
	Name  string
	Email string
	Phone string
}

type ProductQuote struct {
	ID        uuid.UUID
	UnitPrice int64
	Available bool
	Stock     int
}

type RetailerOfferQuote struct {
	ID                uuid.UUID
	RetailerProductID string
	AccessoryName     string
	Category          string
	UnitPrice         int64
	OriginalPrice     int64
	InStock           bool
	SourceURL         string
	FetchedAt         time.Time
}

type DiscountType string

const (
	DiscountPercentage  DiscountType = "PERCENTAGE"
	DiscountFixedAmount DiscountType = "FIXED_AMOUNT"
)

type Promotion struct {
	CouponCode            string
	DiscountType          DiscountType
	DiscountValue         int64
	MinimumSubtotal       int64
	MaximumDiscountAmount *int64
	Active                bool
	StartsAt              *time.Time
	EndsAt                *time.Time
}
