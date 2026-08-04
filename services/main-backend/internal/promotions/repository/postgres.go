package repository

import (
	"context"
	"errors"

	"shopwise/retail/internal/promotions"

	"gorm.io/gorm"
)

// Migrate executes database migrations for the promotions domain.
// NOTE: This project uses AutoMigrate as the primary mechanism for schema management.
// The SQL migration in internal/db/migrations/000004_create_checkout_promotions.up.sql
// serves as documentation and a fallback if raw SQL execution is ever required.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&promotions.CheckoutPromotion{})
}

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) promotions.Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r *postgresRepository) Create(ctx context.Context, promotion *promotions.CheckoutPromotion) error {
	return r.db.WithContext(ctx).Create(promotion).Error
}

func (r *postgresRepository) FindBySession(ctx context.Context, sessionID string) (*promotions.CheckoutPromotion, error) {
	var promo promotions.CheckoutPromotion
	if err := r.db.WithContext(ctx).Where("session_id = ?", sessionID).First(&promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &promo, nil
}

func (r *postgresRepository) FindByEligibilityKey(ctx context.Context, userID *string, deviceID *string) (*promotions.CheckoutPromotion, error) {
	var promo promotions.CheckoutPromotion
	
	query := r.db.WithContext(ctx)
	if userID != nil && *userID != "" {
		query = query.Where("user_id = ?", *userID)
	} else if deviceID != nil && *deviceID != "" {
		query = query.Where("device_id = ?", *deviceID)
	} else {
		return nil, errors.New("must provide either userID or deviceID")
	}

	// Get the most recent one to check cooldowns
	if err := query.Order("created_at desc").First(&promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &promo, nil
}

func (r *postgresRepository) UpdateStatus(ctx context.Context, promotionID string, status promotions.PromotionStatus) error {
	res := r.db.WithContext(ctx).Model(&promotions.CheckoutPromotion{}).Where("promotion_id = ?", promotionID).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("promotion not found")
	}
	return nil
}

func (r *postgresRepository) SaveVoucherResult(ctx context.Context, promotionID string, status promotions.PromotionStatus, voucherID *string, voucherStatus *promotions.VoucherStatus) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if voucherID != nil {
		updates["voucher_id"] = *voucherID
	}
	if voucherStatus != nil {
		updates["voucher_status"] = *voucherStatus
	}
	
	res := r.db.WithContext(ctx).Model(&promotions.CheckoutPromotion{}).Where("promotion_id = ?", promotionID).Updates(updates)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("promotion not found")
	}
	return nil
}
