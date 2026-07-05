package domain

import (
	"time"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID              uuid.UUID
	Email           string
	Name            string
	PasswordHash    *string
	EmailVerifiedAt *time.Time
}
