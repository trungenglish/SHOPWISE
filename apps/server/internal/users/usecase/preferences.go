package usecase

import (
	"context"
	"errors"
	"slices"
	"strings"

	"shopwise/apps/server/internal/platform/apperror"
	"shopwise/apps/server/internal/users/domain"

	"github.com/google/uuid"
)

var validBudget = []string{domain.BudgetLow, domain.BudgetMedium, domain.BudgetHigh}
var validBrand = []string{domain.BrandLoyal, domain.BrandFlexible, domain.BrandAgnostic}
var validCategories = []string{"laptop", "monitor"}

type UpdateProfileInput struct {
	Name                *string
	BudgetSensitivity   *string
	PreferredCategories *[]string
	BrandOpenness       *string
	DefaultCurrency     *string
}

func (s *Service) GetProfile(ctx context.Context, userID uuid.UUID) (*domain.UserProfile, error) {
	profile, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, apperror.NotFound("user not found", err)
		}
		return nil, apperror.Internal("failed to get profile", err)
	}
	return profile, nil
}

func (s *Service) UpdateProfile(ctx context.Context, userID uuid.UUID, input UpdateProfileInput) (*domain.UserProfile, error) {
	profile, err := s.repo.GetProfile(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, apperror.NotFound("user not found", err)
		}
		return nil, apperror.Internal("failed to get profile", err)
	}

	if input.Name != nil {
		name := strings.TrimSpace(*input.Name)
		if name == "" {
			return nil, apperror.Validation("name cannot be empty", nil)
		}
		if err := s.repo.UpdateProfileName(ctx, userID, name); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				return nil, apperror.NotFound("user not found", err)
			}
			return nil, apperror.Internal("failed to update profile", err)
		}
		profile.Name = name
	}

	prefs := profile.Preferences
	updated := false

	if input.BudgetSensitivity != nil {
		value := strings.TrimSpace(*input.BudgetSensitivity)
		if !slices.Contains(validBudget, value) {
			return nil, apperror.Validation("invalid budget sensitivity", nil)
		}
		prefs.BudgetSensitivity = value
		updated = true
	}
	if input.PreferredCategories != nil {
		if len(*input.PreferredCategories) == 0 {
			return nil, apperror.Validation("preferred categories cannot be empty", nil)
		}
		for _, category := range *input.PreferredCategories {
			if !slices.Contains(validCategories, category) {
				return nil, apperror.Validation("invalid preferred category", nil)
			}
		}
		prefs.PreferredCategories = *input.PreferredCategories
		updated = true
	}
	if input.BrandOpenness != nil {
		value := strings.TrimSpace(*input.BrandOpenness)
		if !slices.Contains(validBrand, value) {
			return nil, apperror.Validation("invalid brand openness", nil)
		}
		prefs.BrandOpenness = value
		updated = true
	}
	if input.DefaultCurrency != nil {
		currency := strings.ToUpper(strings.TrimSpace(*input.DefaultCurrency))
		if len(currency) != 3 {
			return nil, apperror.Validation("default currency must be a 3-letter code", nil)
		}
		prefs.DefaultCurrency = currency
		updated = true
	}

	if updated {
		if err := s.repo.UpsertPreferences(ctx, prefs); err != nil {
			return nil, apperror.Internal("failed to update preferences", err)
		}
	}

	return s.repo.GetProfile(ctx, userID)
}
