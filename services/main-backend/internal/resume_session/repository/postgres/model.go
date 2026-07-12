package postgres

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type ResumeToken struct {
	ID                  uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	SessionID           uuid.UUID  `gorm:"type:uuid;not null;index"`
	UserID              *uuid.UUID `gorm:"type:uuid"`
	TokenHash           string     `gorm:"type:varchar(255);not null;index"`
	ExpiresAt           time.Time  `gorm:"not null"`
	ConsumedAt          *time.Time
	RevokedAt           *time.Time
	IssuedContextHash   *string
	ConsumedContextHash *string
	CreatedAt           time.Time
}

func (ResumeToken) TableName() string {
	return "resume_tokens"
}

type NotificationLog struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	SessionID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	Provider     string         `gorm:"type:varchar(50);not null"`
	Status       string         `gorm:"type:varchar(50);not null"`
	ErrorDetails datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt    time.Time
}

func (NotificationLog) TableName() string {
	return "notification_logs"
}
