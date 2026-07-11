package usecase

import (
	"fmt"
	"math"
	"time"

	"shopwise/retail/internal/orders/domain"
)

const taxPercentage int64 = 8

func checkedMultiply(amount int64, quantity int) (int64, error) {
	if amount < 0 || quantity < 0 {
		return 0, fmt.Errorf("amount and quantity must not be negative")
	}
	quantity64 := int64(quantity)
	if quantity64 != 0 && amount > math.MaxInt64/quantity64 {
		return 0, fmt.Errorf("integer multiplication overflow")
	}
	return amount * quantity64, nil
}

func checkedAdd(amounts ...int64) (int64, error) {
	var total int64
	for _, amount := range amounts {
		if amount < 0 || total > math.MaxInt64-amount {
			return 0, fmt.Errorf("integer addition overflow")
		}
		total += amount
	}
	return total, nil
}

func percentageOf(amount, percentage int64) (int64, error) {
	if amount < 0 || percentage < 0 || percentage > 100 {
		return 0, fmt.Errorf("invalid percentage calculation")
	}
	whole := amount / 100
	if percentage != 0 && whole > math.MaxInt64/percentage {
		return 0, fmt.Errorf("percentage calculation overflow")
	}
	base := whole * percentage
	roundedRemainder := ((amount%100)*percentage + 50) / 100
	return checkedAdd(base, roundedRemainder)
}

func calculateDiscount(subtotal int64, promotion *domain.Promotion, now time.Time) (int64, error) {
	if promotion == nil {
		return 0, nil
	}
	if !promotion.Active {
		return 0, fmt.Errorf("coupon is inactive")
	}
	if promotion.StartsAt != nil && now.Before(*promotion.StartsAt) {
		return 0, fmt.Errorf("coupon is not active yet")
	}
	if promotion.EndsAt != nil && !now.Before(*promotion.EndsAt) {
		return 0, fmt.Errorf("coupon has expired")
	}
	if subtotal < promotion.MinimumSubtotal {
		return 0, fmt.Errorf("order subtotal does not meet the coupon minimum")
	}
	if promotion.MinimumSubtotal < 0 {
		return 0, fmt.Errorf("coupon minimum subtotal is invalid")
	}
	if promotion.MaximumDiscountAmount != nil && *promotion.MaximumDiscountAmount < 0 {
		return 0, fmt.Errorf("coupon maximum discount is invalid")
	}

	var discount int64
	var err error
	switch promotion.DiscountType {
	case domain.DiscountPercentage:
		if promotion.DiscountValue <= 0 || promotion.DiscountValue > 100 {
			return 0, fmt.Errorf("coupon percentage is invalid")
		}
		discount, err = percentageOf(subtotal, promotion.DiscountValue)
	case domain.DiscountFixedAmount:
		if promotion.DiscountValue <= 0 {
			return 0, fmt.Errorf("coupon amount is invalid")
		}
		discount = promotion.DiscountValue
	default:
		return 0, fmt.Errorf("coupon discount type is invalid")
	}
	if err != nil {
		return 0, err
	}
	if promotion.MaximumDiscountAmount != nil && discount > *promotion.MaximumDiscountAmount {
		discount = *promotion.MaximumDiscountAmount
	}
	if discount > subtotal {
		discount = subtotal
	}
	return discount, nil
}
