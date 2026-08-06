package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/platform/apperror"

	"github.com/google/uuid"
)

const (
	defaultListLimit = 20
	maximumListLimit = 100
)

type Service struct {
	orders            OrderRepository
	customers         CustomerReader
	products          ProductReader
	promotions        PromotionReader
	retailerOffers    RetailerOfferReader
	retailerRefresher RetailerOfferRefresher
	idempotentOrders  IdempotentOrderRepository
	confirmationQueue OrderConfirmationEnqueuer
}

func (service *Service) WithOrderConfirmation(queue OrderConfirmationEnqueuer) *Service {
	service.confirmationQueue = queue
	return service
}

func (service *Service) WithRetailerOffers(
	offers RetailerOfferReader,
	refresher RetailerOfferRefresher,
	idempotentOrders IdempotentOrderRepository,
) *Service {
	service.retailerOffers = offers
	service.retailerRefresher = refresher
	service.idempotentOrders = idempotentOrders
	return service
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
	ProductID       string
	RetailerOfferID string
	Quantity        int
}

type CreateInput struct {
	AuthenticatedCustomerID uuid.UUID
	CustomerID              uuid.UUID
	Items                   []CreateItemInput
	FulfillmentMethod       string
	ShippingAddress         string
	CouponCode              string
	IdempotencyKey          string
}

type OfferChangedError struct{ Offer domain.RetailerOfferQuote }

func (err *OfferChangedError) Error() string { return "retailer offer price changed" }

type OfferUnavailableError struct{ OfferID uuid.UUID }

func (err *OfferUnavailableError) Error() string { return "retailer offer is unavailable" }

type ListInput struct {
	AuthenticatedCustomerID uuid.UUID
	CustomerID              uuid.UUID
	Limit                   int
	Offset                  int
}

type ListResult struct {
	Orders []domain.Order
	Limit  int
	Offset int
}

func (service *Service) List(ctx context.Context, input ListInput) (*ListResult, error) {
	if input.AuthenticatedCustomerID == uuid.Nil || input.CustomerID == uuid.Nil {
		return nil, apperror.Validation("customer id is required", nil)
	}
	if input.AuthenticatedCustomerID != input.CustomerID {
		return nil, apperror.Forbidden("customer id does not match the authenticated user", nil)
	}

	limit := input.Limit
	if limit <= 0 {
		limit = defaultListLimit
	}
	if limit > maximumListLimit {
		limit = maximumListLimit
	}
	offset := input.Offset
	if offset < 0 {
		offset = 0
	}

	orders, err := service.orders.ListByCustomer(ctx, input.CustomerID, limit, offset)
	if err != nil {
		return nil, apperror.Internal("failed to list orders", err)
	}
	return &ListResult{Orders: orders, Limit: limit, Offset: offset}, nil
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
	validated, err := validateRequestedItems(input.Items)
	if err != nil {
		return nil, err
	}
	if len(validated.retailerOfferIDs) > 0 && strings.TrimSpace(input.IdempotencyKey) == "" {
		return nil, apperror.Validation("Idempotency-Key is required for Phong Vu items", nil)
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
	payloadHash := ""
	if len(validated.retailerOfferIDs) > 0 {
		if service.idempotentOrders == nil {
			return nil, apperror.Internal("retailer checkout is not configured", nil)
		}
		payloadHash = hashCreateInput(input)
		existing, lookupErr := service.idempotentOrders.FindIdempotent(
			ctx, input.CustomerID, strings.TrimSpace(input.IdempotencyKey), payloadHash,
		)
		if lookupErr != nil {
			return nil, mapIdempotencyError(lookupErr)
		}
		if existing != nil {
			return existing, nil
		}
	}

	items := []domain.Item{}
	internalSubtotal := int64(0)
	if len(validated.productIDs) > 0 {
		quotes, loadErr := service.products.GetProducts(ctx, validated.productIDs)
		if loadErr != nil {
			return nil, apperror.Internal("failed to load product pricing", loadErr)
		}
		items, internalSubtotal, err = buildPricedItems(validated.productIDs, validated.productQuantities, quotes)
		if err != nil {
			return nil, err
		}
	}
	retailerItems, retailerSubtotal, err := service.buildRetailerItems(ctx, validated)
	if err != nil {
		return nil, err
	}
	subtotalAmount, err := checkedAdd(internalSubtotal, retailerSubtotal)
	if err != nil {
		return nil, apperror.Validation("subtotal amount exceeds the supported range", err)
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

	discountAmount, err := calculateDiscount(internalSubtotal, promotion, time.Now().UTC())
	if err != nil {
		return nil, apperror.Validation(err.Error(), err)
	}
	shippingAmount := int64(0)
	taxableAmount := internalSubtotal - discountAmount
	taxAmount, err := percentageOf(taxableAmount, taxPercentage)
	if err != nil {
		return nil, apperror.Validation("tax amount exceeds the supported range", err)
	}
	internalTotal, err := checkedAdd(taxableAmount, shippingAmount, taxAmount)
	if err != nil {
		return nil, apperror.Validation("total amount exceeds the supported range", err)
	}
	totalAmount, err := checkedAdd(internalTotal, retailerSubtotal)
	if err != nil {
		return nil, apperror.Validation("total amount exceeds the supported range", err)
	}
	status := domain.StatusPending
	if len(retailerItems) > 0 {
		status = domain.StatusPendingSupplierConfirmation
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
		RetailerItems:     retailerItems,
		CouponCode:        couponCode,
		SubtotalAmount:    subtotalAmount,
		DiscountAmount:    discountAmount,
		ShippingAmount:    shippingAmount,
		TaxAmount:         taxAmount,
		TotalAmount:       totalAmount,
		Status:            status,
		CreatedAt:         time.Now().UTC(),
	}
	order.EstimatedDeliveryFrom = order.CreatedAt.AddDate(0, 0, 3)
	order.EstimatedDeliveryTo = order.CreatedAt.AddDate(0, 0, 5)
	order.ConfirmationEmailStatus = "failed"
	if len(retailerItems) > 0 {
		created, createErr := service.idempotentOrders.CreateIdempotent(ctx, order, strings.TrimSpace(input.IdempotencyKey), payloadHash)
		if createErr != nil {
			return nil, mapIdempotencyError(createErr)
		}
		service.enqueueConfirmation(ctx, created)
		return created, nil
	}
	if err := service.orders.Create(ctx, order); err != nil {
		return nil, apperror.Internal("failed to create order", err)
	}

	service.enqueueConfirmation(ctx, order)
	return order, nil
}

func (service *Service) enqueueConfirmation(ctx context.Context, order *domain.Order) {
	if service.confirmationQueue != nil && service.confirmationQueue.EnqueueOrderConfirmation(ctx, order) == nil {
		order.ConfirmationEmailStatus = "queued"
	}
}

func mapIdempotencyError(err error) error {
	if errors.Is(err, domain.ErrIdempotencyPayloadConflict) {
		return &apperror.AppError{Code: "IDEMPOTENCY_CONFLICT", Message: "Idempotency-Key was reused with a different payload", HTTPStatus: 409, Err: err}
	}
	return err
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

type validatedItems struct {
	productIDs         []uuid.UUID
	productQuantities  map[uuid.UUID]int
	retailerOfferIDs   []uuid.UUID
	retailerQuantities map[uuid.UUID]int
}

func validateRequestedItems(inputs []CreateItemInput) (validatedItems, error) {
	if len(inputs) == 0 {
		return validatedItems{}, apperror.Validation("at least one order item is required", nil)
	}

	result := validatedItems{productQuantities: make(map[uuid.UUID]int), retailerQuantities: make(map[uuid.UUID]int)}
	for index, input := range inputs {
		hasProduct := strings.TrimSpace(input.ProductID) != ""
		hasOffer := strings.TrimSpace(input.RetailerOfferID) != ""
		if hasProduct == hasOffer {
			return validatedItems{}, apperror.Validation(fmt.Sprintf("items[%d] must contain exactly one of product_id or retailer_offer_id", index), nil)
		}
		if input.Quantity <= 0 {
			return validatedItems{}, apperror.Validation(fmt.Sprintf("items[%d].quantity must be greater than zero", index), nil)
		}
		rawID := input.ProductID
		if hasOffer {
			rawID = input.RetailerOfferID
		}
		id, err := uuid.Parse(rawID)
		if err != nil {
			return validatedItems{}, apperror.Validation(fmt.Sprintf("items[%d] contains an invalid UUID", index), err)
		}
		if hasProduct {
			if _, exists := result.productQuantities[id]; exists {
				return validatedItems{}, apperror.Validation(fmt.Sprintf("items[%d].product_id is duplicated", index), nil)
			}
			result.productIDs = append(result.productIDs, id)
			result.productQuantities[id] = input.Quantity
			continue
		}
		if _, exists := result.retailerQuantities[id]; exists {
			return validatedItems{}, apperror.Validation(fmt.Sprintf("items[%d].retailer_offer_id is duplicated", index), nil)
		}
		result.retailerOfferIDs = append(result.retailerOfferIDs, id)
		result.retailerQuantities[id] = input.Quantity
	}
	return result, nil
}

func (service *Service) buildRetailerItems(ctx context.Context, validated validatedItems) ([]domain.RetailerItem, int64, error) {
	if len(validated.retailerOfferIDs) == 0 {
		return nil, 0, nil
	}
	if service.retailerOffers == nil || service.retailerRefresher == nil {
		return nil, 0, apperror.Internal("retailer checkout is not configured", nil)
	}
	offers, err := service.retailerOffers.GetRetailerOffers(ctx, validated.retailerOfferIDs, validated.productIDs)
	if err != nil {
		return nil, 0, apperror.Internal("failed to load retailer offers", err)
	}
	items := make([]domain.RetailerItem, 0, len(validated.retailerOfferIDs))
	var subtotal int64
	for index, offerID := range validated.retailerOfferIDs {
		stored, exists := offers[offerID]
		if !exists {
			return nil, 0, apperror.Validation(fmt.Sprintf("items[%d].retailer_offer_id does not exist", index), nil)
		}
		refreshed := stored
		var refreshErr error
		if stored.Retailer != "shopwise" {
			refreshed, refreshErr = service.retailerRefresher.RefreshRetailerOffer(ctx, stored)
		}
		invalidPrice := refreshed.UnitPrice < 0 || (refreshed.UnitPrice == 0 && stored.Retailer != "shopwise")
		if refreshErr != nil || !refreshed.InStock || invalidPrice {
			return nil, 0, &OfferUnavailableError{OfferID: offerID}
		}
		if refreshed.UnitPrice != stored.UnitPrice {
			if updater, ok := service.retailerOffers.(RetailerOfferUpdater); ok {
				if updateErr := updater.UpdateRetailerOffer(ctx, refreshed); updateErr != nil {
					return nil, 0, apperror.Internal("failed to save refreshed retailer offer", updateErr)
				}
			}
			return nil, 0, &OfferChangedError{Offer: refreshed}
		}
		quantity := validated.retailerQuantities[offerID]
		lineAmount, multiplyErr := checkedMultiply(refreshed.UnitPrice, quantity)
		if multiplyErr != nil {
			return nil, 0, apperror.Validation(fmt.Sprintf("retailer item %d amount exceeds the supported range", index), multiplyErr)
		}
		subtotal, err = checkedAdd(subtotal, lineAmount)
		if err != nil {
			return nil, 0, apperror.Validation("retailer subtotal exceeds the supported range", err)
		}
		items = append(items, domain.RetailerItem{RetailerOfferID: offerID, Name: refreshed.AccessoryName, Quantity: quantity, UnitPrice: refreshed.UnitPrice, SourceURL: refreshed.SourceURL, VerifiedAt: refreshed.FetchedAt})
	}
	return items, subtotal, nil
}

func hashCreateInput(input CreateInput) string {
	payload, _ := json.Marshal(input)
	digest := sha256.Sum256(payload)
	return hex.EncodeToString(digest[:])
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
