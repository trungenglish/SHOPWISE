package postgres

import (
	"context"
	"errors"
	"fmt"

	"shopwise/retail/internal/orders/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) *Repository {
	return &Repository{database: database}
}

func (repository *Repository) Create(ctx context.Context, order *domain.Order) error {
	return repository.database.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		return createOrder(transaction, order)
	})
}

func createOrder(transaction *gorm.DB, order *domain.Order) error {
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
	if len(itemModels) > 0 {
		if err := transaction.Create(&itemModels).Error; err != nil {
			return fmt.Errorf("create order items: %w", err)
		}
	}
	retailerModels := make([]RetailerOrderItemModel, 0, len(order.RetailerItems))
	for _, item := range order.RetailerItems {
		retailerModels = append(retailerModels, RetailerOrderItemModel{
			ID: uuid.New(), OrderID: order.ID, RetailerOfferID: item.RetailerOfferID,
			Name: item.Name, Quantity: item.Quantity, UnitPrice: item.UnitPrice,
			SourceURL: item.SourceURL, VerifiedAt: item.VerifiedAt,
		})
	}
	if len(retailerModels) > 0 {
		if err := transaction.Create(&retailerModels).Error; err != nil {
			return fmt.Errorf("create retailer order items: %w", err)
		}
	}

	return nil
}

func (repository *Repository) CreateIdempotent(ctx context.Context, order *domain.Order, key, payloadHash string) (*domain.Order, error) {
	created := order
	err := repository.database.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		record := CheckoutIdempotencyModel{CustomerID: order.CustomerID, Key: key, PayloadHash: payloadHash, OrderID: order.ID, CreatedAt: order.CreatedAt}
		result := transaction.Clauses(clause.OnConflict{DoNothing: true}).Create(&record)
		if result.Error != nil {
			return fmt.Errorf("reserve checkout idempotency key: %w", result.Error)
		}
		if result.RowsAffected == 0 {
			var existing CheckoutIdempotencyModel
			if err := transaction.First(&existing, "customer_id = ? AND key = ?", order.CustomerID, key).Error; err != nil {
				return fmt.Errorf("load checkout idempotency key: %w", err)
			}
			if existing.PayloadHash != payloadHash {
				return domain.ErrIdempotencyPayloadConflict
			}
			existingOrder, err := loadOrder(transaction, existing.OrderID)
			if err != nil {
				return err
			}
			created = existingOrder
			return nil
		}
		return createOrder(transaction, order)
	})
	return created, err
}

func (repository *Repository) FindIdempotent(
	ctx context.Context,
	customerID uuid.UUID,
	key, payloadHash string,
) (*domain.Order, error) {
	var existing CheckoutIdempotencyModel
	err := repository.database.WithContext(ctx).
		First(&existing, "customer_id = ? AND key = ?", customerID, key).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("load checkout idempotency key: %w", err)
	}
	if existing.PayloadHash != payloadHash {
		return nil, domain.ErrIdempotencyPayloadConflict
	}
	return loadOrder(repository.database.WithContext(ctx), existing.OrderID)
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
		Preload("RetailerItems", func(database *gorm.DB) *gorm.DB { return database.Order("id ASC") }).
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

func loadOrder(database *gorm.DB, orderID uuid.UUID) (*domain.Order, error) {
	var model OrderModel
	if err := database.Preload("Items").Preload("RetailerItems").First(&model, "id = ?", orderID).Error; err != nil {
		return nil, fmt.Errorf("load idempotent order: %w", err)
	}
	order := toDomainOrder(&model)
	return &order, nil
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
	retailerItems := make([]domain.RetailerItem, 0, len(model.RetailerItems))
	for _, item := range model.RetailerItems {
		retailerItems = append(retailerItems, domain.RetailerItem{
			RetailerOfferID: item.RetailerOfferID, Name: item.Name, Quantity: item.Quantity,
			UnitPrice: item.UnitPrice, SourceURL: item.SourceURL, VerifiedAt: item.VerifiedAt,
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
		RetailerItems:     retailerItems,
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
