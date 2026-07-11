package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"shopwise/retail/internal/orders/usecase"
	"shopwise/retail/internal/platform/apperror"
	"shopwise/retail/internal/platform/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type Handler struct {
	service               *usecase.Service
	developmentAuthBypass bool
}

func NewHandler(service *usecase.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) WithDevelopmentAuthBypass() *Handler {
	handler.developmentAuthBypass = true
	return handler
}

// Create godoc
//
//	@Summary	Create a checkout order
//	@Tags		checkout
//	@Accept		json
//	@Produce	json
//	@Security	BearerAuth
//	@Param		body	body		CheckoutRequest	true	"Checkout items"
//	@Success	201		{object}	OrderResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	403		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/checkout [post]
func (handler *Handler) Create(ctx *gin.Context) {
	authenticatedCustomerID, exists := middleware.UserID(ctx)
	if !exists && !handler.developmentAuthBypass {
		_ = ctx.Error(apperror.Unauthorized("authenticated user is required", nil))
		return
	}

	var request CheckoutRequest
	if err := decodeCheckoutRequest(ctx, &request); err != nil {
		_ = ctx.Error(apperror.Validation("invalid request body", err))
		return
	}
	customerID, err := uuid.Parse(request.CustomerID)
	if err != nil {
		_ = ctx.Error(apperror.Validation("customer_id must be a valid UUID", err))
		return
	}
	if !exists {
		authenticatedCustomerID = customerID
	}

	items := make([]usecase.CreateItemInput, 0, len(request.Items))
	for _, item := range request.Items {
		items = append(items, usecase.CreateItemInput{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	order, err := handler.service.Create(ctx.Request.Context(), usecase.CreateInput{
		AuthenticatedCustomerID: authenticatedCustomerID,
		CustomerID:              customerID,
		Items:                   items,
		FulfillmentMethod:       request.FulfillmentMethod,
		ShippingAddress:         request.ShippingAddress,
		CouponCode:              request.CouponCode,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, toOrderResponse(order))
}

// List godoc
//
//	@Summary	List customer orders
//	@Tags		orders
//	@Produce	json
//	@Security	BearerAuth
//	@Param		customer_id	query		string	false	"Customer ID (required only in debug auth bypass)"	format(uuid)
//	@Param		limit		query		int		false	"Page size"											default(20)	maximum(100)
//	@Param		offset		query		int		false	"Offset"											default(0)
//	@Success	200			{object}	OrderListResponse
//	@Failure	400			{object}	ErrorResponse
//	@Failure	401			{object}	ErrorResponse
//	@Failure	403			{object}	ErrorResponse
//	@Failure	500			{object}	ErrorResponse
//	@Router		/orders [get]
func (handler *Handler) List(ctx *gin.Context) {
	authenticatedCustomerID, authenticated := middleware.UserID(ctx)
	if !authenticated && !handler.developmentAuthBypass {
		_ = ctx.Error(apperror.Unauthorized("authenticated user is required", nil))
		return
	}

	requestedCustomerID := authenticatedCustomerID
	customerIDQuery := ctx.Query("customer_id")
	if customerIDQuery != "" {
		customerID, err := uuid.Parse(customerIDQuery)
		if err != nil {
			_ = ctx.Error(apperror.Validation("customer_id must be a valid UUID", err))
			return
		}
		requestedCustomerID = customerID
	} else if !authenticated {
		_ = ctx.Error(apperror.Validation("customer_id is required when checkout auth bypass is enabled", nil))
		return
	}
	if !authenticated {
		authenticatedCustomerID = requestedCustomerID
	}

	limit, err := parseIntegerQuery(ctx, "limit", 20)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	offset, err := parseIntegerQuery(ctx, "offset", 0)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	result, err := handler.service.List(ctx.Request.Context(), usecase.ListInput{
		AuthenticatedCustomerID: authenticatedCustomerID,
		CustomerID:              requestedCustomerID,
		Limit:                   limit,
		Offset:                  offset,
	})
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	items := make([]OrderResponse, 0, len(result.Orders))
	for index := range result.Orders {
		items = append(items, toOrderResponse(&result.Orders[index]))
	}
	ctx.JSON(http.StatusOK, OrderListResponse{Items: items, Limit: result.Limit, Offset: result.Offset})
}

func parseIntegerQuery(ctx *gin.Context, name string, defaultValue int) (int, error) {
	value := ctx.Query(name)
	if value == "" {
		return defaultValue, nil
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, apperror.Validation(name+" must be an integer", err)
	}
	return parsed, nil
}

func decodeCheckoutRequest(ctx *gin.Context, request *CheckoutRequest) error {
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(request); err != nil {
		return err
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		if err == nil {
			return fmt.Errorf("request body must contain one JSON object")
		}
		return err
	}
	if binding.Validator == nil {
		return nil
	}
	return binding.Validator.ValidateStruct(request)
}

func RegisterRoutes(group *gin.RouterGroup, handler *Handler, verifier middleware.TokenVerifier) {
	if !handler.developmentAuthBypass {
		group.Use(middleware.Auth(verifier))
	}
	group.POST("", handler.Create)
}

func RegisterListRoutes(group *gin.RouterGroup, handler *Handler, verifier middleware.TokenVerifier) {
	if !handler.developmentAuthBypass {
		group.Use(middleware.Auth(verifier))
	}
	group.GET("", handler.List)
}
