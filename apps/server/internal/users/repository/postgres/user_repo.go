package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"shopwise/apps/server/internal/users/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user *domain.User) error {
	model := toModel(user)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if isDuplicateKey(err) {
			return domain.ErrDuplicateEmail
		}
		return fmt.Errorf("create user: %w", err)
	}
	*user = *toDomain(model)
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var model UserModel
	if err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return toDomain(&model), nil
}

func (r *Repository) List(ctx context.Context, limit, offset int) ([]domain.User, error) {
	var models []UserModel
	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	users := make([]domain.User, 0, len(models))
	for i := range models {
		users = append(users, *toDomain(&models[i]))
	}
	return users, nil
}

func (r *Repository) Update(ctx context.Context, user *domain.User) error {
	model := toModel(user)
	result := r.db.WithContext(ctx).Model(&UserModel{}).
		Where("id = ?", user.ID).
		Updates(map[string]any{
			"email": model.Email,
			"name":  model.Name,
		})
	if result.Error != nil {
		if isDuplicateKey(result.Error) {
			return domain.ErrDuplicateEmail
		}
		return fmt.Errorf("update user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&UserModel{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("delete user: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func toModel(user *domain.User) *UserModel {
	return &UserModel{
		ID:              user.ID,
		Email:           user.Email,
		Name:            user.Name,
		EmailVerifiedAt: user.EmailVerifiedAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}

func toDomain(model *UserModel) *domain.User {
	return &domain.User{
		ID:              model.ID,
		Email:           model.Email,
		Name:            model.Name,
		EmailVerifiedAt: model.EmailVerifiedAt,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
	}
}

func isDuplicateKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	return strings.Contains(strings.ToLower(err.Error()), "duplicate key")
}
