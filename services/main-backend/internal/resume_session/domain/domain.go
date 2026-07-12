package domain

import (
	"time"

	"github.com/google/uuid"
)

type ResumeToken struct {
	ID                  uuid.UUID
	SessionID           uuid.UUID
	UserID              *uuid.UUID
	TokenHash           string
	ExpiresAt           time.Time
	ConsumedAt          *time.Time
	RevokedAt           *time.Time
	IssuedContextHash   *string
	ConsumedContextHash *string
	CreatedAt           time.Time
}

type NotificationLog struct {
	ID           uuid.UUID
	SessionID    uuid.UUID
	Provider     string
	Status       string
	ErrorDetails *string
	CreatedAt    time.Time
}
