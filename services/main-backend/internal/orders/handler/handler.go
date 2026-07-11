package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"shopwise/retail/internal/orders/usecase"
	"shopwise/retail/internal/platform/apperror"
	"shopwise/retail/internal/platform/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler(service *usecase.Service) *Handler {
	return &Handler{service: service}
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
	if !exists {
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
	group.Use(middleware.Auth(verifier))
	group.POST("", handler.Create)
}
