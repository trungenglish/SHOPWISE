package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"shopwise/apps/server/internal/users/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *Repository) GetProfile(ctx context.Context, userID uuid.UUID) (*domain.UserProfile, error) {
	var model UserModel
	if err := r.db.WithContext(ctx).First(&model, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get profile user: %w", err)
	}

	prefs, err := r.getOrCreatePreferences(ctx, userID)
	if err != nil {
		return nil, err
	}

	user := toDomain(&model)
	return &domain.UserProfile{
		User:        *user,
		Preferences: *prefs,
	}, nil
}

func (r *Repository) UpdateProfileName(ctx context.Context, userID uuid.UUID, name string) error {
	result := r.db.WithContext(ctx).Model(&UserModel{}).
		Where("id = ?", userID).
		Update("name", name)
	if result.Error != nil {
		return fmt.Errorf("update profile name: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *Repository) UpsertPreferences(ctx context.Context, prefs domain.UserPreferences) error {
	categoriesJSON, err := json.Marshal(prefs.PreferredCategories)
	if err != nil {
		return fmt.Errorf("marshal preferred categories: %w", err)
	}
	model := &UserPreferencesModel{
		UserID:              prefs.UserID,
		BudgetSensitivity:   prefs.BudgetSensitivity,
		PreferredCategories: categoriesJSON,
		BrandOpenness:       prefs.BrandOpenness,
		DefaultCurrency:     prefs.DefaultCurrency,
	}
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *Repository) getOrCreatePreferences(ctx context.Context, userID uuid.UUID) (*domain.UserPreferences, error) {
	var model UserPreferencesModel
	err := r.db.WithContext(ctx).First(&model, "user_id = ?", userID).Error
	if err == nil {
		return preferencesToDomain(&model)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("get preferences: %w", err)
	}

	defaults := domain.DefaultPreferences(userID)
	if err := r.UpsertPreferences(ctx, defaults); err != nil {
		return nil, fmt.Errorf("create default preferences: %w", err)
	}
	return &defaults, nil
}

func preferencesToDomain(model *UserPreferencesModel) (*domain.UserPreferences, error) {
	var categories []string
	if len(model.PreferredCategories) > 0 {
		if err := json.Unmarshal(model.PreferredCategories, &categories); err != nil {
			return nil, fmt.Errorf("unmarshal preferred categories: %w", err)
		}
	}
	return &domain.UserPreferences{
		UserID:              model.UserID,
		BudgetSensitivity:   model.BudgetSensitivity,
		PreferredCategories: categories,
		BrandOpenness:       model.BrandOpenness,
		DefaultCurrency:     model.DefaultCurrency,
	}, nil
}
