package postgres

import (
	"github.com/google/uuid"
)

type UserPreferencesModel struct {
	UserID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	BudgetSensitivity   string    `gorm:"size:16;not null"`
	PreferredCategories []byte    `gorm:"type:jsonb;not null"`
	BrandOpenness       string    `gorm:"size:16;not null"`
	DefaultCurrency     string    `gorm:"size:3;not null"`
}

func (m *UserPreferencesModel) TableName() string {
	return "user_preferences"
}
