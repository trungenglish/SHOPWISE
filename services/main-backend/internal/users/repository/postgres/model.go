package postgres

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Email           string     `gorm:"size:255;uniqueIndex;not null"`
	Name            string     `gorm:"size:255;not null"`
	Phone           string     `gorm:"size:32;not null;default:''"`
	EmailVerifiedAt *time.Time `gorm:"type:timestamptz"`
	PasswordHash    *string    `gorm:"size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (m *UserModel) TableName() string {
	return "users"
}

func (m *UserModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&UserModel{}, &UserPreferencesModel{})
}
