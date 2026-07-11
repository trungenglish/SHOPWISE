package usecase

import (
	"context"

	"shopwise/retail/internal/orders/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
}

type CustomerReader interface {
	GetCustomer(ctx context.Context, customerID uuid.UUID) (*domain.CustomerSnapshot, error)
}

type ProductReader interface {
	GetProducts(ctx context.Context, productIDs []uuid.UUID) (map[uuid.UUID]domain.ProductQuote, error)
}

type PromotionReader interface {
	GetPromotion(ctx context.Context, couponCode string) (*domain.Promotion, error)
}
