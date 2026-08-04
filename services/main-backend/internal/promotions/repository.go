package promotions

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, promotion *CheckoutPromotion) error
	FindBySession(ctx context.Context, sessionID string) (*CheckoutPromotion, error)
	FindByEligibilityKey(ctx context.Context, userID *string, deviceID *string) (*CheckoutPromotion, error)
	UpdateStatus(ctx context.Context, promotionID string, status PromotionStatus) error
	SaveVoucherResult(ctx context.Context, promotionID string, status PromotionStatus, voucherID *string, voucherStatus *VoucherStatus) error
}
