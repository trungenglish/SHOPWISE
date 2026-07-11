package database

import (
	"context"

	"shopwise/retail/internal/platform/database/model"

	"gorm.io/gorm"
)

func EnsureStartupKey(ctx context.Context, db *gorm.DB) error {
	var count int64
	if err := db.WithContext(ctx).Model(&model.AppMetadata{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	return db.WithContext(ctx).Create(&model.AppMetadata{
		Key:   "startup",
		Value: "ok",
	}).Error
}
