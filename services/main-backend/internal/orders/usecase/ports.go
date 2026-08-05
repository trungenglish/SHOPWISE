package usecase

import (
	"context"

	"shopwise/retail/internal/orders/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	ListByCustomer(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]domain.Order, error)
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

type RetailerOfferReader interface {
	GetRetailerOffers(context.Context, []uuid.UUID, []uuid.UUID) (map[uuid.UUID]domain.RetailerOfferQuote, error)
}

type RetailerOfferRefresher interface {
	RefreshRetailerOffer(context.Context, domain.RetailerOfferQuote) (domain.RetailerOfferQuote, error)
}

type RetailerOfferUpdater interface {
	UpdateRetailerOffer(context.Context, domain.RetailerOfferQuote) error
}

type IdempotentOrderRepository interface {
	FindIdempotent(context.Context, uuid.UUID, string, string) (*domain.Order, error)
	CreateIdempotent(context.Context, *domain.Order, string, string) (*domain.Order, error)
}

type OrderConfirmationEnqueuer interface {
	EnqueueOrderConfirmation(context.Context, *domain.Order) error
}
