package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInteractionConflict   = errors.New("interaction payload conflict")
	ErrInteractionInProgress = errors.New("interaction is already in progress")
)

type DecisionSession struct {
	ID              uuid.UUID
	UserID          *uuid.UUID
	AnonymousID     *string
	Title           string
	Status          string
	ParentSessionID *uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Messages        []SessionMessage
}

type SessionMessage struct {
	ID             uuid.UUID
	SessionID      uuid.UUID
	Role           string
	Content        string
	ReasoningGraph string // JSON string
	PinnedProducts string // JSON string
	CreatedAt      time.Time
}

type SessionInteraction struct {
	SessionID        uuid.UUID
	InteractionID    uuid.UUID
	PayloadHash      string
	Status           string
	LeaseUntil       time.Time
	ResponseEnvelope string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserPreference struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	Category        string
	Value           string // JSON string
	SourceSessionID *uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
