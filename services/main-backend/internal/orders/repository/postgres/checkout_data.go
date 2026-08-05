package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"shopwise/retail/internal/accessories"
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

func (repository *Repository) GetRetailerOffers(
	ctx context.Context,
	offerIDs []uuid.UUID,
	productIDs []uuid.UUID,
) (map[uuid.UUID]domain.RetailerOfferQuote, error) {
	var offers []accessories.RetailerOfferModel
	if err := repository.database.WithContext(ctx).Where("id IN ?", offerIDs).Find(&offers).Error; err != nil {
		return nil, fmt.Errorf("get retailer offers: %w", err)
	}
	accessoryIDs := make([]uuid.UUID, 0, len(offers))
	for _, offer := range offers {
		accessoryIDs = append(accessoryIDs, offer.AccessoryID)
	}
	var accessoryModels []accessories.AccessoryModel
	if len(accessoryIDs) > 0 {
		if err := repository.database.WithContext(ctx).Where("id IN ?", accessoryIDs).Find(&accessoryModels).Error; err != nil {
			return nil, fmt.Errorf("get retailer accessories: %w", err)
		}
	}
	names := make(map[uuid.UUID]accessories.AccessoryModel, len(accessoryModels))
	for _, accessory := range accessoryModels {
		names[accessory.ID] = accessory
	}
	compatible := make(map[uuid.UUID]bool)
	if len(productIDs) > 0 {
		var rows []accessories.AccessoryCompatibilityModel
		if err := repository.database.WithContext(ctx).Where("accessory_id IN ? AND product_id IN ?", accessoryIDs, productIDs).Find(&rows).Error; err != nil {
			return nil, fmt.Errorf("get bundle compatibility: %w", err)
		}
		for _, row := range rows {
			compatible[row.AccessoryID] = true
		}
	}
	quotes := make(map[uuid.UUID]domain.RetailerOfferQuote, len(offers))
	for _, offer := range offers {
		if offer.Retailer == "shopwise" && !compatible[offer.AccessoryID] {
			continue
		}
		accessory := names[offer.AccessoryID]
		quotes[offer.ID] = domain.RetailerOfferQuote{
			ID: offer.ID, Retailer: offer.Retailer, RetailerProductID: offer.RetailerProductID, AccessoryName: accessory.Name,
			Category: accessory.Category, UnitPrice: offer.Price, OriginalPrice: offer.OriginalPrice,
			InStock: offer.InStock, SourceURL: offer.SourceURL, FetchedAt: offer.FetchedAt,
		}
	}
	return quotes, nil
}

func (repository *Repository) UpdateRetailerOffer(ctx context.Context, quote domain.RetailerOfferQuote) error {
	updates := map[string]any{
		"price": quote.UnitPrice, "original_price": quote.OriginalPrice,
		"in_stock": quote.InStock, "source_url": quote.SourceURL, "fetched_at": quote.FetchedAt,
	}
	result := repository.database.WithContext(ctx).
		Model(&accessories.RetailerOfferModel{}).
		Where("id = ? AND retailer = ?", quote.ID, "phongvu").
		Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("update retailer offer: %w", result.Error)
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf("update retailer offer: offer not found")
	}
	return nil
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
