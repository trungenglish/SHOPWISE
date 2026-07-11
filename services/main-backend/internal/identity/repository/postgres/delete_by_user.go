package postgres

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) DeleteUserData(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Delete(&EmailVerificationTokenModel{}).Error; err != nil {
		return err
	}
	if err := db.WithContext(ctx).Where("user_id = ?", userID).Delete(&RefreshTokenModel{}).Error; err != nil {
		return err
	}
	return db.WithContext(ctx).Where("user_id = ?", userID).Delete(&AuthIdentityModel{}).Error
}
