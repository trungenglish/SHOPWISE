package job

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"shopwise/retail/internal/orders/domain"

	"github.com/hibiken/asynq"
)

const (
	TypeSendWelcomeEmail     = "notification:send_welcome_email"
	TypePriceCheck           = "notification:price_check"
	TypePhongVuAccessorySync = "retailer:phongvu_accessory_sync"
	TypeOrderConfirmation    = "notification:order_confirmation"
)

type WelcomeEmailPayload struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

type OrderConfirmationPayload struct {
	OrderID        string    `json:"order_id"`
	CustomerEmail  string    `json:"customer_email"`
	CustomerName   string    `json:"customer_name"`
	TotalAmount    int64     `json:"total_amount"`
	SubtotalAmount int64     `json:"subtotal_amount"`
	DiscountAmount int64     `json:"discount_amount"`
	ShippingAmount int64     `json:"shipping_amount"`
	TaxAmount      int64     `json:"tax_amount"`
	Items          []string  `json:"items"`
	DeliveryFrom   time.Time `json:"delivery_from"`
	DeliveryTo     time.Time `json:"delivery_to"`
}

type Client struct {
	client *asynq.Client
}

func NewClient(redisURL string) (*Client, error) {
	opts, err := asynq.ParseRedisURI(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}
	return &Client{client: asynq.NewClient(opts)}, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) EnqueueWelcomeEmail(ctx context.Context, userID, email string) error {
	payload, err := json.Marshal(WelcomeEmailPayload{UserID: userID, Email: email})
	if err != nil {
		return fmt.Errorf("marshal welcome email payload: %w", err)
	}

	task := asynq.NewTask(TypeSendWelcomeEmail, payload)
	if _, err := c.client.EnqueueContext(ctx, task); err != nil {
		return fmt.Errorf("enqueue welcome email: %w", err)
	}
	return nil
}

func (c *Client) EnqueuePriceCheck(ctx context.Context) error {
	task := asynq.NewTask(TypePriceCheck, nil)
	if _, err := c.client.EnqueueContext(ctx, task); err != nil {
		return fmt.Errorf("enqueue price check: %w", err)
	}
	return nil
}

func (c *Client) EnqueueOrderConfirmation(ctx context.Context, order *domain.Order) error {
	items := make([]string, 0, len(order.Items)+len(order.RetailerItems))
	for _, item := range order.Items {
		items = append(items, fmt.Sprintf("Product %s x%d: %d VND", item.ProductID, item.Quantity, item.UnitPrice))
	}
	for _, item := range order.RetailerItems {
		items = append(items, fmt.Sprintf("%s x%d: %d VND", item.Name, item.Quantity, item.UnitPrice))
	}
	payload, err := json.Marshal(OrderConfirmationPayload{
		OrderID: order.ID.String(), CustomerEmail: order.CustomerEmail, CustomerName: order.CustomerName,
		TotalAmount: order.TotalAmount, SubtotalAmount: order.SubtotalAmount, DiscountAmount: order.DiscountAmount,
		ShippingAmount: order.ShippingAmount, TaxAmount: order.TaxAmount, Items: items,
		DeliveryFrom: order.EstimatedDeliveryFrom, DeliveryTo: order.EstimatedDeliveryTo,
	})
	if err != nil {
		return fmt.Errorf("marshal order confirmation payload: %w", err)
	}
	task := asynq.NewTask(TypeOrderConfirmation, payload)
	_, err = c.client.EnqueueContext(ctx, task, asynq.TaskID("order-confirmation-"+order.ID.String()))
	if err != nil && !errors.Is(err, asynq.ErrTaskIDConflict) {
		return fmt.Errorf("enqueue order confirmation: %w", err)
	}
	return nil
}
