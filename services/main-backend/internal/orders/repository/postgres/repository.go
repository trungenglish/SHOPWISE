package postgres

import (
	"context"
	"fmt"

	"shopwise/retail/internal/orders/domain"

	"github.com/google/uuid"
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

func (repository *Repository) ListByCustomer(
	ctx context.Context,
	customerID uuid.UUID,
	limit, offset int,
) ([]domain.Order, error) {
	var models []OrderModel
	if err := repository.database.WithContext(ctx).
		Preload("Items", func(database *gorm.DB) *gorm.DB {
			return database.Order("product_id ASC")
		}).
		Where("customer_id = ?", customerID).
		Order("created_at DESC").
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("list customer orders: %w", err)
	}

	orders := make([]domain.Order, 0, len(models))
	for index := range models {
		orders = append(orders, toDomainOrder(&models[index]))
	}
	return orders, nil
}

func toDomainOrder(model *OrderModel) domain.Order {
	items := make([]domain.Item, 0, len(model.Items))
	for _, item := range model.Items {
		items = append(items, domain.Item{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
		})
	}
	return domain.Order{
		ID:                model.ID,
		CustomerID:        model.CustomerID,
		CustomerName:      model.CustomerName,
		CustomerEmail:     model.CustomerEmail,
		CustomerPhone:     model.CustomerPhone,
		FulfillmentMethod: domain.FulfillmentMethod(model.FulfillmentMethod),
		ShippingAddress:   model.ShippingAddress,
		Items:             items,
		CouponCode:        model.CouponCode,
		SubtotalAmount:    model.SubtotalAmount,
		DiscountAmount:    model.DiscountAmount,
		ShippingAmount:    model.ShippingAmount,
		TaxAmount:         model.TaxAmount,
		TotalAmount:       model.TotalAmount,
		Status:            domain.Status(model.Status),
		CreatedAt:         model.CreatedAt,
	}
}
