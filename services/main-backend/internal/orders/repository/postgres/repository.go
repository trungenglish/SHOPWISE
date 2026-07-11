package postgres

import (
	"context"
	"fmt"

	"shopwise/retail/internal/orders/domain"

	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) *Repository {
	return &Repository{database: database}
}

func (repository *Repository) Create(ctx context.Context, order *domain.Order) error {
	return repository.database.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		orderModel := OrderModel{
			ID:                order.ID,
			CustomerID:        order.CustomerID,
			CustomerName:      order.CustomerName,
			CustomerEmail:     order.CustomerEmail,
			CustomerPhone:     order.CustomerPhone,
			FulfillmentMethod: string(order.FulfillmentMethod),
			ShippingAddress:   order.ShippingAddress,
			CouponCode:        order.CouponCode,
			SubtotalAmount:    order.SubtotalAmount,
			DiscountAmount:    order.DiscountAmount,
			ShippingAmount:    order.ShippingAmount,
			TaxAmount:         order.TaxAmount,
			TotalAmount:       order.TotalAmount,
			Status:            string(order.Status),
			CreatedAt:         order.CreatedAt,
		}
		if err := transaction.Create(&orderModel).Error; err != nil {
			return fmt.Errorf("create order: %w", err)
		}

		itemModels := make([]OrderItemModel, 0, len(order.Items))
		for _, item := range order.Items {
			itemModels = append(itemModels, OrderItemModel{
				OrderID:   order.ID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				UnitPrice: item.UnitPrice,
			})
		}
		if err := transaction.Create(&itemModels).Error; err != nil {
			return fmt.Errorf("create order items: %w", err)
		}

		return nil
	})
}
