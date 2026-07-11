package usecase_test

import (
	"context"
	"errors"
	"math"
	"testing"
	"time"

	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/orders/usecase"
	"shopwise/retail/internal/platform/apperror"

	"github.com/google/uuid"
)

type checkoutRepositoryStub struct {
	customer         *domain.CustomerSnapshot
	products         map[uuid.UUID]domain.ProductQuote
	promotion        *domain.Promotion
	customerErr      error
	productsErr      error
	promotionErr     error
	createErr        error
	persisted        *domain.Order
	listed           []domain.Order
	listErr          error
	listedCustomerID uuid.UUID
	listedLimit      int
	listedOffset     int
}

func (stub *checkoutRepositoryStub) Create(_ context.Context, order *domain.Order) error {
	if stub.createErr != nil {
		return stub.createErr
	}
	copyOrder := *order
	copyOrder.Items = append([]domain.Item(nil), order.Items...)
	stub.persisted = &copyOrder
	return nil
}

func (stub *checkoutRepositoryStub) ListByCustomer(
	_ context.Context,
	customerID uuid.UUID,
	limit, offset int,
) ([]domain.Order, error) {
	stub.listedCustomerID = customerID
	stub.listedLimit = limit
	stub.listedOffset = offset
	return stub.listed, stub.listErr
}

func (stub *checkoutRepositoryStub) GetCustomer(context.Context, uuid.UUID) (*domain.CustomerSnapshot, error) {
	return stub.customer, stub.customerErr
}

func (stub *checkoutRepositoryStub) GetProducts(context.Context, []uuid.UUID) (map[uuid.UUID]domain.ProductQuote, error) {
	return stub.products, stub.productsErr
}

func (stub *checkoutRepositoryStub) GetPromotion(context.Context, string) (*domain.Promotion, error) {
	return stub.promotion, stub.promotionErr
}

func newService(stub *checkoutRepositoryStub) *usecase.Service {
	return usecase.NewService(stub, stub, stub, stub)
}

func validInput(customerID, productID uuid.UUID) usecase.CreateInput {
	return usecase.CreateInput{
		AuthenticatedCustomerID: customerID,
		CustomerID:              customerID,
		Items: []usecase.CreateItemInput{
			{ProductID: productID.String(), Quantity: 1},
		},
		FulfillmentMethod: string(domain.FulfillmentDelivery),
		ShippingAddress:   "1 Nguyen Hue, District 1, Ho Chi Minh City",
	}
}

func validRepository(customerID, productID uuid.UUID, unitPrice int64) *checkoutRepositoryStub {
	return &checkoutRepositoryStub{
		customer: &domain.CustomerSnapshot{
			ID:    customerID,
			Name:  "Nguyen Van A",
			Email: "customer@example.com",
			Phone: "0912345678",
		},
		products: map[uuid.UUID]domain.ProductQuote{
			productID: {
				ID:        productID,
				UnitPrice: unitPrice,
				Available: true,
				Stock:     10,
			},
		},
	}
}

func TestCreateUsesOfficialPriceAndPersistsCheckoutSnapshot(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	repository := validRepository(customerID, productID, 87_475_000)
	service := newService(repository)
	input := validInput(customerID, productID)

	order, err := service.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if order.CustomerName != "Nguyen Van A" || order.CustomerEmail != "customer@example.com" || order.CustomerPhone != "0912345678" {
		t.Fatalf("customer snapshot = %q/%q/%q", order.CustomerName, order.CustomerEmail, order.CustomerPhone)
	}
	if order.Items[0].UnitPrice != 87_475_000 {
		t.Fatalf("UnitPrice = %d, want official price 87475000", order.Items[0].UnitPrice)
	}
	if order.SubtotalAmount != 87_475_000 {
		t.Fatalf("SubtotalAmount = %d, want 87475000", order.SubtotalAmount)
	}
	if order.DiscountAmount != 0 || order.ShippingAmount != 0 {
		t.Fatalf("discount/shipping = %d/%d, want 0/0", order.DiscountAmount, order.ShippingAmount)
	}
	if order.TaxAmount != 6_998_000 {
		t.Fatalf("TaxAmount = %d, want 6998000", order.TaxAmount)
	}
	if order.TotalAmount != 94_473_000 {
		t.Fatalf("TotalAmount = %d, want 94473000", order.TotalAmount)
	}
	if repository.persisted == nil || repository.persisted.ID != order.ID {
		t.Fatal("repository did not receive the completed order")
	}
}

func TestCreateAppliesActivePercentageCoupon(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	repository := validRepository(customerID, productID, 100_000)
	repository.promotion = &domain.Promotion{
		CouponCode:    "SHOPWISE5",
		DiscountType:  domain.DiscountPercentage,
		DiscountValue: 5,
		Active:        true,
	}
	service := newService(repository)
	input := validInput(customerID, productID)
	input.CouponCode = " shopwise5 "

	order, err := service.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if order.CouponCode != "SHOPWISE5" {
		t.Fatalf("CouponCode = %q, want SHOPWISE5", order.CouponCode)
	}
	if order.DiscountAmount != 5_000 {
		t.Fatalf("DiscountAmount = %d, want 5000", order.DiscountAmount)
	}
	if order.TaxAmount != 7_600 || order.TotalAmount != 102_600 {
		t.Fatalf("tax/total = %d/%d, want 7600/102600", order.TaxAmount, order.TotalAmount)
	}
}

func TestCreateAppliesFixedCouponMaximum(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	maximumDiscount := int64(20_000)
	repository := validRepository(customerID, productID, 100_000)
	repository.promotion = &domain.Promotion{
		CouponCode:            "FIXED",
		DiscountType:          domain.DiscountFixedAmount,
		DiscountValue:         50_000,
		MaximumDiscountAmount: &maximumDiscount,
		Active:                true,
	}
	input := validInput(customerID, productID)
	input.CouponCode = "FIXED"

	order, err := newService(repository).Create(context.Background(), input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if order.DiscountAmount != 20_000 || order.TaxAmount != 6_400 || order.TotalAmount != 86_400 {
		t.Fatalf("discount/tax/total = %d/%d/%d", order.DiscountAmount, order.TaxAmount, order.TotalAmount)
	}
}

func TestCreateRoundsTaxToNearestWholeVND(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	repository := validRepository(customerID, productID, 107)

	order, err := newService(repository).Create(context.Background(), validInput(customerID, productID))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if order.TaxAmount != 9 {
		t.Fatalf("TaxAmount = %d, want rounded value 9", order.TaxAmount)
	}
}

func TestCreateRejectsCustomerIDThatDoesNotMatchAuthentication(t *testing.T) {
	t.Parallel()

	productID := uuid.New()
	authenticatedCustomerID := uuid.New()
	repository := validRepository(authenticatedCustomerID, productID, 100)
	input := validInput(uuid.New(), productID)
	input.AuthenticatedCustomerID = authenticatedCustomerID

	_, err := newService(repository).Create(context.Background(), input)
	assertAppErrorCode(t, err, "FORBIDDEN")
	if repository.persisted != nil {
		t.Fatal("order was persisted for mismatched customer identity")
	}
}

func TestCreateValidatesFulfillment(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	tests := []struct {
		name     string
		method   string
		address  string
		wantCode string
	}{
		{name: "delivery requires address", method: string(domain.FulfillmentDelivery), wantCode: "VALIDATION_FAILED"},
		{name: "unknown method", method: "DRONE", address: "somewhere", wantCode: "VALIDATION_FAILED"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			repository := validRepository(customerID, productID, 100)
			input := validInput(customerID, productID)
			input.FulfillmentMethod = test.method
			input.ShippingAddress = test.address
			_, err := newService(repository).Create(context.Background(), input)
			assertAppErrorCode(t, err, test.wantCode)
		})
	}
}

func TestCreateRejectsIncompleteCustomerAndUnavailableProducts(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	tests := []struct {
		name   string
		mutate func(*checkoutRepositoryStub)
	}{
		{name: "missing phone", mutate: func(repository *checkoutRepositoryStub) { repository.customer.Phone = "" }},
		{name: "missing product", mutate: func(repository *checkoutRepositoryStub) { repository.products = map[uuid.UUID]domain.ProductQuote{} }},
		{name: "unavailable product", mutate: func(repository *checkoutRepositoryStub) {
			quote := repository.products[productID]
			quote.Available = false
			repository.products[productID] = quote
		}},
		{name: "insufficient stock", mutate: func(repository *checkoutRepositoryStub) {
			quote := repository.products[productID]
			quote.Stock = 0
			repository.products[productID] = quote
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			repository := validRepository(customerID, productID, 100)
			test.mutate(repository)
			_, err := newService(repository).Create(context.Background(), validInput(customerID, productID))
			assertAppErrorCode(t, err, "VALIDATION_FAILED")
		})
	}
}

func TestCreateRejectsInvalidPromotionAndAmountOverflow(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()

	t.Run("invalid promotion", func(t *testing.T) {
		repository := validRepository(customerID, productID, 100)
		repository.promotionErr = domain.ErrPromotionNotFound
		input := validInput(customerID, productID)
		input.CouponCode = "INVALID"
		_, err := newService(repository).Create(context.Background(), input)
		assertAppErrorCode(t, err, "VALIDATION_FAILED")
	})

	t.Run("line amount overflow", func(t *testing.T) {
		repository := validRepository(customerID, productID, math.MaxInt64)
		input := validInput(customerID, productID)
		input.Items[0].Quantity = 2
		_, err := newService(repository).Create(context.Background(), input)
		assertAppErrorCode(t, err, "VALIDATION_FAILED")
	})
}

func TestCreateRejectsExpiredAndBelowMinimumPromotion(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	now := time.Now().UTC()
	tests := []domain.Promotion{
		{CouponCode: "EXPIRED", DiscountType: domain.DiscountPercentage, DiscountValue: 5, Active: true, EndsAt: timePointer(now.Add(-time.Hour))},
		{CouponCode: "MINIMUM", DiscountType: domain.DiscountFixedAmount, DiscountValue: 10, MinimumSubtotal: 1_000, Active: true},
	}

	for _, promotion := range tests {
		promotion := promotion
		t.Run(promotion.CouponCode, func(t *testing.T) {
			t.Parallel()
			repository := validRepository(customerID, productID, 100)
			repository.promotion = &promotion
			input := validInput(customerID, productID)
			input.CouponCode = promotion.CouponCode
			_, err := newService(repository).Create(context.Background(), input)
			assertAppErrorCode(t, err, "VALIDATION_FAILED")
		})
	}
}

func TestCreateRejectsNegativePromotionCap(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	negativeCap := int64(-1)
	repository := validRepository(customerID, productID, 100)
	repository.promotion = &domain.Promotion{
		CouponCode:            "BROKEN",
		DiscountType:          domain.DiscountPercentage,
		DiscountValue:         5,
		MaximumDiscountAmount: &negativeCap,
		Active:                true,
	}
	input := validInput(customerID, productID)
	input.CouponCode = "BROKEN"

	_, err := newService(repository).Create(context.Background(), input)
	assertAppErrorCode(t, err, "VALIDATION_FAILED")
}

func TestCreateMapsRepositoryFailureToInternalError(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	productID := uuid.New()
	repository := validRepository(customerID, productID, 100)
	repository.createErr = errors.New("database unavailable")

	_, err := newService(repository).Create(context.Background(), validInput(customerID, productID))
	assertAppErrorCode(t, err, "INTERNAL_ERROR")
}

func TestListReturnsCustomerOrdersWithNormalizedPagination(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	repository := &checkoutRepositoryStub{listed: []domain.Order{{ID: uuid.New(), CustomerID: customerID}}}
	service := newService(repository)

	result, err := service.List(context.Background(), usecase.ListInput{
		AuthenticatedCustomerID: customerID,
		CustomerID:              customerID,
		Limit:                   500,
		Offset:                  -10,
	})
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if len(result.Orders) != 1 || result.Limit != 100 || result.Offset != 0 {
		t.Fatalf("result orders/limit/offset = %d/%d/%d", len(result.Orders), result.Limit, result.Offset)
	}
	if repository.listedCustomerID != customerID || repository.listedLimit != 100 || repository.listedOffset != 0 {
		t.Fatalf("repository args = %s/%d/%d", repository.listedCustomerID, repository.listedLimit, repository.listedOffset)
	}
}

func TestListRejectsCustomerMismatch(t *testing.T) {
	t.Parallel()

	repository := &checkoutRepositoryStub{}
	_, err := newService(repository).List(context.Background(), usecase.ListInput{
		AuthenticatedCustomerID: uuid.New(),
		CustomerID:              uuid.New(),
		Limit:                   20,
	})
	assertAppErrorCode(t, err, "FORBIDDEN")
}

func TestListMapsRepositoryFailure(t *testing.T) {
	t.Parallel()

	customerID := uuid.New()
	repository := &checkoutRepositoryStub{listErr: errors.New("database unavailable")}
	_, err := newService(repository).List(context.Background(), usecase.ListInput{
		AuthenticatedCustomerID: customerID,
		CustomerID:              customerID,
		Limit:                   20,
	})
	assertAppErrorCode(t, err, "INTERNAL_ERROR")
}

func assertAppErrorCode(t *testing.T, err error, code string) {
	t.Helper()
	var appError *apperror.AppError
	if !errors.As(err, &appError) || appError.Code != code {
		t.Fatalf("error = %v, want app error code %s", err, code)
	}
}

func timePointer(value time.Time) *time.Time {
	return &value
}
