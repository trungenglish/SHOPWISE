package promotions

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type VoucherResult struct {
	VoucherID string
	Code      string
}

type VoucherGenerator interface {
	GenerateVoucher(ctx context.Context, promotion *CheckoutPromotion) (*VoucherResult, error)
}

type localVoucherGenerator struct {
	// In a real application, this would interact with a database repository 
	// for a `vouchers` table to persist the generated voucher.
}

func NewLocalVoucherGenerator() VoucherGenerator {
	return &localVoucherGenerator{}
}

func (g *localVoucherGenerator) GenerateVoucher(ctx context.Context, promotion *CheckoutPromotion) (*VoucherResult, error) {
	// Simulate local generation and exactly-once enforcement logic
	// In reality, this would insert into a `vouchers` table with a unique constraint 
	// on `promotion_id` to enforce exactly-once issuance.

	// For MVP, we just generate a random UUID and Code
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	code := "PROMO-" + string(b)

	return &VoucherResult{
		VoucherID: uuid.New().String(),
		Code:      code,
	}, nil
}
