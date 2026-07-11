package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound  = errors.New("checkout customer not found")
	ErrPromotionNotFound = errors.New("promotion not found")
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
