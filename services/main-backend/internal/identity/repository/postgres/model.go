package postgres

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthIdentityModel struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID `gorm:"type:uuid;not null;index"`
	Provider        string    `gorm:"size:32;not null"`
	ProviderSubject string    `gorm:"size:255;not null"`
	CreatedAt       time.Time
}

func (m *AuthIdentityModel) TableName() string {
	return "auth_identities"
}

func (m *AuthIdentityModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

type EmailVerificationTokenModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	TokenHash string     `gorm:"size:255;not null"`
	ExpiresAt time.Time  `gorm:"type:timestamptz;not null"`
	UsedAt    *time.Time `gorm:"type:timestamptz"`
	CreatedAt time.Time
}

func (m *EmailVerificationTokenModel) TableName() string {
	return "email_verification_tokens"
}

func (m *EmailVerificationTokenModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

type RefreshTokenModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index"`
	TokenHash string     `gorm:"size:255;not null;uniqueIndex"`
	ExpiresAt time.Time  `gorm:"type:timestamptz;not null"`
	RevokedAt *time.Time `gorm:"type:timestamptz"`
	CreatedAt time.Time
}

func (m *RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}

func (m *RefreshTokenModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
