package postgres

import (
	"context"
	"fmt"

	"shopwise/retail/internal/users/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) DeleteAccount(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	if err := db.WithContext(ctx).Where("user_id = ?", userID).Delete(&UserPreferencesModel{}).Error; err != nil {
		return fmt.Errorf("delete user preferences: %w", err)
	}

	result := db.WithContext(ctx).Delete(&UserModel{}, "id = ?", userID)
	if result.Error != nil {
		return fmt.Errorf("delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
