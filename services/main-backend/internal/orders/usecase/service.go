package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/platform/apperror"

	"github.com/google/uuid"
)

type Service struct {
	orders     OrderRepository
	customers  CustomerReader
	products   ProductReader
	promotions PromotionReader
}

func NewService(
	orders OrderRepository,
	customers CustomerReader,
	products ProductReader,
	promotions PromotionReader,
) *Service {
	return &Service{
		orders:     orders,
		customers:  customers,
		products:   products,
		promotions: promotions,
	}
}

type CreateItemInput struct {
	ProductID string
	Quantity  int
}

type CreateInput struct {
	AuthenticatedCustomerID uuid.UUID
	CustomerID              uuid.UUID
	Items                   []CreateItemInput
	FulfillmentMethod       string
	ShippingAddress         string
	CouponCode              string
}

func (service *Service) Create(ctx context.Context, input CreateInput) (*domain.Order, error) {
	if input.AuthenticatedCustomerID == uuid.Nil || input.CustomerID == uuid.Nil {
		return nil, apperror.Validation("customer id is required", nil)
	}
	if input.AuthenticatedCustomerID != input.CustomerID {
		return nil, apperror.Forbidden("customer id does not match the authenticated user", nil)
	}

	fulfillmentMethod, shippingAddress, err := validateFulfillment(input.FulfillmentMethod, input.ShippingAddress)
	if err != nil {
		return nil, err
	}
	productIDs, quantities, err := validateRequestedItems(input.Items)
	if err != nil {
		return nil, err
	}

	customer, err := service.customers.GetCustomer(ctx, input.CustomerID)
	if err != nil {
		if errors.Is(err, domain.ErrCustomerNotFound) {
			return nil, apperror.Validation("customer profile was not found", err)
		}
		return nil, apperror.Internal("failed to load customer profile", err)
	}
	if customer == nil || strings.TrimSpace(customer.Name) == "" || strings.TrimSpace(customer.Email) == "" || strings.TrimSpace(customer.Phone) == "" {
		return nil, apperror.Validation("customer name, email, and phone must be completed before checkout", nil)
	}

	quotes, err := service.products.GetProducts(ctx, productIDs)
	if err != nil {
		return nil, apperror.Internal("failed to load product pricing", err)
	}
	items, subtotalAmount, err := buildPricedItems(productIDs, quantities, quotes)
	if err != nil {
		return nil, err
	}

	couponCode := strings.ToUpper(strings.TrimSpace(input.CouponCode))
	var promotion *domain.Promotion
	if couponCode != "" {
		promotion, err = service.promotions.GetPromotion(ctx, couponCode)
		if err != nil {
			if errors.Is(err, domain.ErrPromotionNotFound) {
				return nil, apperror.Validation("coupon code is invalid", err)
			}
			return nil, apperror.Internal("failed to validate coupon", err)
		}
	}

	discountAmount, err := calculateDiscount(subtotalAmount, promotion, time.Now().UTC())
	if err != nil {
		return nil, apperror.Validation(err.Error(), err)
	}
	shippingAmount := int64(0)
	taxableAmount := subtotalAmount - discountAmount
	taxAmount, err := percentageOf(taxableAmount, taxPercentage)
	if err != nil {
		return nil, apperror.Validation("tax amount exceeds the supported range", err)
	}
	totalAmount, err := checkedAdd(taxableAmount, shippingAmount, taxAmount)
	if err != nil {
		return nil, apperror.Validation("total amount exceeds the supported range", err)
	}

	order := &domain.Order{
		ID:                uuid.New(),
		CustomerID:        input.CustomerID,
		CustomerName:      strings.TrimSpace(customer.Name),
		CustomerEmail:     strings.TrimSpace(customer.Email),
		CustomerPhone:     strings.TrimSpace(customer.Phone),
		FulfillmentMethod: fulfillmentMethod,
		ShippingAddress:   shippingAddress,
		Items:             items,
		CouponCode:        couponCode,
		SubtotalAmount:    subtotalAmount,
		DiscountAmount:    discountAmount,
		ShippingAmount:    shippingAmount,
		TaxAmount:         taxAmount,
		TotalAmount:       totalAmount,
		Status:            domain.StatusPending,
		CreatedAt:         time.Now().UTC(),
	}
	if err := service.orders.Create(ctx, order); err != nil {
		return nil, apperror.Internal("failed to create order", err)
	}

	return order, nil
}

func validateFulfillment(method, address string) (domain.FulfillmentMethod, string, error) {
	normalizedMethod := domain.FulfillmentMethod(strings.ToUpper(strings.TrimSpace(method)))
	normalizedAddress := strings.TrimSpace(address)
	switch normalizedMethod {
	case domain.FulfillmentDelivery:
		if normalizedAddress == "" {
			return "", "", apperror.Validation("shipping address is required for delivery", nil)
		}
	case domain.FulfillmentStorePickup:
		normalizedAddress = ""
	default:
		return "", "", apperror.Validation("fulfillment method must be DELIVERY or STORE_PICKUP", nil)
	}
	return normalizedMethod, normalizedAddress, nil
}

func validateRequestedItems(inputs []CreateItemInput) ([]uuid.UUID, map[uuid.UUID]int, error) {
	if len(inputs) == 0 {
		return nil, nil, apperror.Validation("at least one order item is required", nil)
	}

	productIDs := make([]uuid.UUID, 0, len(inputs))
	quantities := make(map[uuid.UUID]int, len(inputs))
	for index, input := range inputs {
		productID, err := uuid.Parse(input.ProductID)
		if err != nil {
			return nil, nil, apperror.Validation(fmt.Sprintf("items[%d].product_id must be a valid UUID", index), err)
		}
		if input.Quantity <= 0 {
			return nil, nil, apperror.Validation(fmt.Sprintf("items[%d].quantity must be greater than zero", index), nil)
		}
		if _, exists := quantities[productID]; exists {
			return nil, nil, apperror.Validation(fmt.Sprintf("items[%d].product_id is duplicated", index), nil)
		}
		productIDs = append(productIDs, productID)
		quantities[productID] = input.Quantity
	}
	return productIDs, quantities, nil
}

func buildPricedItems(
	productIDs []uuid.UUID,
	quantities map[uuid.UUID]int,
	quotes map[uuid.UUID]domain.ProductQuote,
) ([]domain.Item, int64, error) {
	items := make([]domain.Item, 0, len(productIDs))
	var subtotalAmount int64
	for index, productID := range productIDs {
		quote, exists := quotes[productID]
		if !exists {
			return nil, 0, apperror.Validation(fmt.Sprintf("items[%d].product_id does not exist", index), nil)
		}
		quantity := quantities[productID]
		if !quote.Available {
			return nil, 0, apperror.Validation(fmt.Sprintf("items[%d] is unavailable", index), nil)
		}
		if quote.Stock < quantity {
			return nil, 0, apperror.Validation(fmt.Sprintf("items[%d] has insufficient stock", index), nil)
		}
		if quote.UnitPrice <= 0 {
			return nil, 0, apperror.Validation(fmt.Sprintf("items[%d] has invalid catalog pricing", index), nil)
		}
		lineAmount, err := checkedMultiply(quote.UnitPrice, quantity)
		if err != nil {
			return nil, 0, apperror.Validation(fmt.Sprintf("items[%d] amount exceeds the supported range", index), err)
		}
		subtotalAmount, err = checkedAdd(subtotalAmount, lineAmount)
		if err != nil {
			return nil, 0, apperror.Validation("subtotal amount exceeds the supported range", err)
		}
		items = append(items, domain.Item{ProductID: productID, Quantity: quantity, UnitPrice: quote.UnitPrice})
	}
	return items, subtotalAmount, nil
}
