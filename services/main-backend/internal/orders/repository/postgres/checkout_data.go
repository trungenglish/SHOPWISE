package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/platform/database/model"
	userpostgres "shopwise/retail/internal/users/repository/postgres"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repository *Repository) GetCustomer(ctx context.Context, customerID uuid.UUID) (*domain.CustomerSnapshot, error) {
	var customer userpostgres.UserModel
	if err := repository.database.WithContext(ctx).First(&customer, "id = ?", customerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrCustomerNotFound
		}
		return nil, fmt.Errorf("get checkout customer: %w", err)
	}
	return &domain.CustomerSnapshot{
		ID: customer.ID, Name: customer.Name, Email: customer.Email, Phone: customer.Phone,
	}, nil
}

func (repository *Repository) GetProducts(
	ctx context.Context,
	productIDs []uuid.UUID,
) (map[uuid.UUID]domain.ProductQuote, error) {
	var products []model.Product
	if err := repository.database.WithContext(ctx).
		Preload("Inventory").
		Where("id IN ?", productIDs).
		Find(&products).Error; err != nil {
		return nil, fmt.Errorf("get checkout products: %w", err)
	}

	quotes := make(map[uuid.UUID]domain.ProductQuote, len(products))
	for _, product := range products {
		quote := domain.ProductQuote{ID: product.ID, UnitPrice: product.Price}
		if product.Inventory != nil {
			quote.Available = product.Inventory.Availability
			quote.Stock = product.Inventory.Stock
		}
		quotes[product.ID] = quote
	}
	return quotes, nil
}

func (repository *Repository) GetPromotion(ctx context.Context, couponCode string) (*domain.Promotion, error) {
	var promotion model.Promotion
	if err := repository.database.WithContext(ctx).
		Where("UPPER(coupon_code) = ?", strings.ToUpper(strings.TrimSpace(couponCode))).
		First(&promotion).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrPromotionNotFound
		}
		return nil, fmt.Errorf("get checkout promotion: %w", err)
	}
	return &domain.Promotion{
		CouponCode:            strings.ToUpper(promotion.CouponCode),
		DiscountType:          domain.DiscountType(promotion.DiscountType),
		DiscountValue:         promotion.DiscountValue,
		MinimumSubtotal:       promotion.MinimumSubtotal,
		MaximumDiscountAmount: promotion.MaximumDiscountAmount,
		Active:                promotion.Active,
		StartsAt:              promotion.StartsAt,
		EndsAt:                promotion.EndsAt,
	}, nil
}
