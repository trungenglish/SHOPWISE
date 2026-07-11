package handler_test

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	identityusecase "shopwise/retail/internal/identity/usecase"
	"shopwise/retail/internal/orders/domain"
	"shopwise/retail/internal/orders/handler"
	"shopwise/retail/internal/orders/usecase"
	"shopwise/retail/internal/platform/middleware"
	"shopwise/retail/internal/platform/testutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type checkoutRepositoryStub struct {
	customer  *domain.CustomerSnapshot
	products  map[uuid.UUID]domain.ProductQuote
	createErr error
}

func (stub *checkoutRepositoryStub) Create(context.Context, *domain.Order) error {
	return stub.createErr
}

func (stub *checkoutRepositoryStub) GetCustomer(context.Context, uuid.UUID) (*domain.CustomerSnapshot, error) {
	return stub.customer, nil
}

func (stub *checkoutRepositoryStub) GetProducts(context.Context, []uuid.UUID) (map[uuid.UUID]domain.ProductQuote, error) {
	return stub.products, nil
}

func (*checkoutRepositoryStub) GetPromotion(context.Context, string) (*domain.Promotion, error) {
	return nil, domain.ErrPromotionNotFound
}

func newCheckoutRouter(t *testing.T, repository *checkoutRepositoryStub) (*gin.Engine, string, uuid.UUID) {
	t.Helper()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	jwtService := identityusecase.NewJWTService("checkout-test-secret", 15*time.Minute)
	customerID := uuid.New()
	token, _, err := jwtService.IssueAccessToken(customerID)
	if err != nil {
		t.Fatalf("issue access token: %v", err)
	}
	if repository.customer == nil {
		repository.customer = &domain.CustomerSnapshot{
			ID: customerID, Name: "Nguyen Van A", Email: "customer@example.com", Phone: "0912345678",
		}
	}

	router := testutil.NewTestRouter()
	router.Use(middleware.ErrorHandler(logger, false))
	v1 := router.Group("/api/v1")
	service := usecase.NewService(repository, repository, repository, repository)
	handler.RegisterRoutes(v1.Group("/checkout"), handler.NewHandler(service), jwtService)
	return router, token, customerID
}

func TestCreateCheckoutReturnsServerCalculatedOrder(t *testing.T) {
	t.Parallel()

	productID := uuid.New()
	repository := &checkoutRepositoryStub{
		products: map[uuid.UUID]domain.ProductQuote{
			productID: {ID: productID, UnitPrice: 100_000, Available: true, Stock: 5},
		},
	}
	router, token, customerID := newCheckoutRouter(t, repository)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/checkout", map[string]any{
		"customer_id":        customerID.String(),
		"items":              []map[string]any{{"product_id": productID.String(), "quantity": 2}},
		"fulfillment_method": "DELIVERY",
		"shipping_address":   "1 Nguyen Hue, District 1",
	}, map[string]string{"Authorization": "Bearer " + token})

	testutil.AssertStatus(t, recorder, http.StatusCreated)
	var response handler.OrderResponse
	testutil.AssertJSON(t, recorder, &response)
	if response.CustomerName != "Nguyen Van A" || response.CustomerPhone != "0912345678" {
		t.Fatalf("customer snapshot = %q/%q", response.CustomerName, response.CustomerPhone)
	}
	if response.SubtotalAmount != 200_000 || response.TaxAmount != 16_000 || response.TotalAmount != 216_000 {
		t.Fatalf("subtotal/tax/total = %d/%d/%d", response.SubtotalAmount, response.TaxAmount, response.TotalAmount)
	}
	if len(response.Items) != 1 || response.Items[0].UnitPrice != 100_000 {
		t.Fatalf("items = %#v, want official unit price", response.Items)
	}
}

func TestCreateCheckoutRejectsClientFinancialFields(t *testing.T) {
	t.Parallel()

	productID := uuid.New()
	repository := &checkoutRepositoryStub{products: map[uuid.UUID]domain.ProductQuote{
		productID: {ID: productID, UnitPrice: 100_000, Available: true, Stock: 5},
	}}
	router, token, customerID := newCheckoutRouter(t, repository)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/checkout", map[string]any{
		"customer_id":        customerID.String(),
		"items":              []map[string]any{{"product_id": productID.String(), "quantity": 1, "unit_price": 1}},
		"fulfillment_method": "STORE_PICKUP",
		"total_amount":       1,
	}, map[string]string{"Authorization": "Bearer " + token})

	testutil.AssertStatus(t, recorder, http.StatusBadRequest)
}

func TestCreateCheckoutRejectsCustomerMismatch(t *testing.T) {
	t.Parallel()

	productID := uuid.New()
	repository := &checkoutRepositoryStub{products: map[uuid.UUID]domain.ProductQuote{
		productID: {ID: productID, UnitPrice: 100, Available: true, Stock: 1},
	}}
	router, token, _ := newCheckoutRouter(t, repository)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/checkout", map[string]any{
		"customer_id":        uuid.NewString(),
		"items":              []map[string]any{{"product_id": productID.String(), "quantity": 1}},
		"fulfillment_method": "STORE_PICKUP",
	}, map[string]string{"Authorization": "Bearer " + token})

	testutil.AssertStatus(t, recorder, http.StatusForbidden)
}

func TestCreateCheckoutRequiresAuthentication(t *testing.T) {
	t.Parallel()

	router, _, customerID := newCheckoutRouter(t, &checkoutRepositoryStub{})
	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/checkout", map[string]any{
		"customer_id": customerID.String(),
	}, nil)

	testutil.AssertStatus(t, recorder, http.StatusUnauthorized)
}

func TestCreateCheckoutAllowsDebugAuthBypass(t *testing.T) {
	t.Parallel()

	productID := uuid.New()
	customerID := uuid.New()
	repository := &checkoutRepositoryStub{
		customer: &domain.CustomerSnapshot{
			ID: customerID, Name: "Nguyen Van A", Email: "customer@example.com", Phone: "0912345678",
		},
		products: map[uuid.UUID]domain.ProductQuote{
			productID: {ID: productID, UnitPrice: 100_000, Available: true, Stock: 5},
		},
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	jwtService := identityusecase.NewJWTService("checkout-test-secret", 15*time.Minute)
	router := testutil.NewTestRouter()
	router.Use(middleware.ErrorHandler(logger, false))
	service := usecase.NewService(repository, repository, repository, repository)
	checkoutHandler := handler.NewHandler(service).WithDevelopmentAuthBypass()
	handler.RegisterRoutes(router.Group("/api/v1/checkout"), checkoutHandler, jwtService)

	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/checkout", map[string]any{
		"customer_id":        customerID.String(),
		"items":              []map[string]any{{"product_id": productID.String(), "quantity": 1}},
		"fulfillment_method": "STORE_PICKUP",
	}, nil)

	testutil.AssertStatus(t, recorder, http.StatusCreated)
}

func TestCreateCheckoutMapsRepositoryFailure(t *testing.T) {
	t.Parallel()

	productID := uuid.New()
	repository := &checkoutRepositoryStub{
		products: map[uuid.UUID]domain.ProductQuote{
			productID: {ID: productID, UnitPrice: 100, Available: true, Stock: 1},
		},
		createErr: errors.New("database unavailable"),
	}
	router, token, customerID := newCheckoutRouter(t, repository)
	recorder := testutil.PerformRequest(t, router, http.MethodPost, "/api/v1/checkout", map[string]any{
		"customer_id":        customerID.String(),
		"items":              []map[string]any{{"product_id": productID.String(), "quantity": 1}},
		"fulfillment_method": "STORE_PICKUP",
	}, map[string]string{"Authorization": "Bearer " + token})

	testutil.AssertStatus(t, recorder, http.StatusInternalServerError)
}
