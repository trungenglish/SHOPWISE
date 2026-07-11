package handler

import (
	"time"

	"shopwise/retail/internal/orders/domain"
)

type CheckoutItemRequest struct {
	ProductID string `json:"product_id" binding:"required,uuid" format:"uuid"`
	Quantity  int    `json:"quantity" binding:"required,gt=0" minimum:"1"`
}

type CheckoutRequest struct {
	CustomerID        string                `json:"customer_id" binding:"required,uuid" format:"uuid"`
	Items             []CheckoutItemRequest `json:"items" binding:"required,min=1,dive"`
	FulfillmentMethod string                `json:"fulfillment_method" binding:"required,oneof=DELIVERY STORE_PICKUP" enums:"DELIVERY,STORE_PICKUP"`
	ShippingAddress   string                `json:"shipping_address,omitempty" binding:"max=1000"`
	CouponCode        string                `json:"coupon_code,omitempty" binding:"max=50"`
}

type OrderItemResponse struct {
	ProductID string `json:"product_id" format:"uuid"`
	Quantity  int    `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
}

type OrderResponse struct {
	OrderID           string              `json:"order_id" format:"uuid"`
	CustomerID        string              `json:"customer_id" format:"uuid"`
	CustomerName      string              `json:"customer_name"`
	CustomerEmail     string              `json:"customer_email" format:"email"`
	CustomerPhone     string              `json:"customer_phone"`
	FulfillmentMethod string              `json:"fulfillment_method" enums:"DELIVERY,STORE_PICKUP"`
	ShippingAddress   string              `json:"shipping_address,omitempty"`
	Items             []OrderItemResponse `json:"items"`
	CouponCode        string              `json:"coupon_code,omitempty"`
	SubtotalAmount    int64               `json:"subtotal_amount"`
	DiscountAmount    int64               `json:"discount_amount"`
	ShippingAmount    int64               `json:"shipping_amount"`
	TaxAmount         int64               `json:"tax_amount"`
	TotalAmount       int64               `json:"total_amount"`
	Status            string              `json:"status" enums:"PENDING,PROCESSING"`
	CreatedAt         time.Time           `json:"created_at" swaggertype:"string" format:"date-time"`
}

type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance,omitempty"`
	Code     string `json:"code,omitempty"`
}

func toOrderResponse(order *domain.Order) OrderResponse {
	items := make([]OrderItemResponse, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, OrderItemResponse{
			ProductID: item.ProductID.String(),
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}

	return OrderResponse{
		OrderID:           order.ID.String(),
		CustomerID:        order.CustomerID.String(),
		CustomerName:      order.CustomerName,
		CustomerEmail:     order.CustomerEmail,
		CustomerPhone:     order.CustomerPhone,
		FulfillmentMethod: string(order.FulfillmentMethod),
		ShippingAddress:   order.ShippingAddress,
		Items:             items,
		CouponCode:        order.CouponCode,
		SubtotalAmount:    order.SubtotalAmount,
		DiscountAmount:    order.DiscountAmount,
		ShippingAmount:    order.ShippingAmount,
		TaxAmount:         order.TaxAmount,
		TotalAmount:       order.TotalAmount,
		Status:            string(order.Status),
		CreatedAt:         order.CreatedAt,
	}
}
