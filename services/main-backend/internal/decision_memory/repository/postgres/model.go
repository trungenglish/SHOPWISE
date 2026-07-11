package postgres

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DecisionSession struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID          *uuid.UUID `gorm:"type:uuid;index"`
	AnonymousID     *string   `gorm:"type:varchar(255);index"`
	Title           string    `gorm:"type:varchar(255);not null"`
	Status          string    `gorm:"type:varchar(50);not null;default:'active'"`
	ParentSessionID *uuid.UUID `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (DecisionSession) TableName() string {
	return "decision_sessions"
}

type SessionMessage struct {
	ID             uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	SessionID      uuid.UUID      `gorm:"type:uuid;not null;index"`
	Role           string         `gorm:"type:varchar(50);not null"`
	Content        string         `gorm:"type:text;not null"`
	ReasoningGraph datatypes.JSON `gorm:"type:jsonb"`
	PinnedProducts datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt      time.Time
}

func (SessionMessage) TableName() string {
	return "session_messages"
}

type ExtractedPreference struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID          uuid.UUID      `gorm:"type:uuid;not null;index"`
	Category        string         `gorm:"type:varchar(100);not null"`
	Value           datatypes.JSON `gorm:"type:jsonb;not null"`
	SourceSessionID *uuid.UUID     `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (ExtractedPreference) TableName() string {
	return "extracted_preferences"
}
