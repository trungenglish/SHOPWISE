package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNotFound       = errors.New("user not found")
	ErrDuplicateEmail = errors.New("email already registered")
)

type User struct {
	ID              uuid.UUID
	Email           string
	Name            string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type UserPreferences struct {
	UserID              uuid.UUID
	BudgetSensitivity   string
	PreferredCategories []string
	BrandOpenness       string
	DefaultCurrency     string
}

type UserProfile struct {
	User
	Preferences UserPreferences
}

const (
	BudgetLow    = "low"
	BudgetMedium = "medium"
	BudgetHigh   = "high"

	BrandLoyal    = "loyal"
	BrandFlexible = "flexible"
	BrandAgnostic = "agnostic"
)

func DefaultPreferences(userID uuid.UUID) UserPreferences {
	return UserPreferences{
		UserID:              userID,
		BudgetSensitivity:   BudgetMedium,
		PreferredCategories: []string{"laptop", "monitor"},
		BrandOpenness:       BrandFlexible,
		DefaultCurrency:     "USD",
	}
}
