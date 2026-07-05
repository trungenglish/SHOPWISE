package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"shopwise/apps/server/internal/identity/domain"
	userpostgres "shopwise/apps/server/internal/users/repository/postgres"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateWithPassword(ctx context.Context, email, name, passwordHash string) (*domain.AuthUser, error) {
	hash := passwordHash
	model := &userpostgres.UserModel{
		Email:        strings.ToLower(strings.TrimSpace(email)),
		Name:         strings.TrimSpace(name),
		PasswordHash: &hash,
	}
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if isDuplicateKey(err) {
			return nil, domain.ErrDuplicateEmail
		}
		return nil, fmt.Errorf("create user: %w", err)
	}
	return toAuthUser(model), nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*domain.AuthUser, error) {
	var model userpostgres.UserModel
	err := r.db.WithContext(ctx).Where("email = ?", strings.ToLower(strings.TrimSpace(email))).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return toAuthUser(&model), nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.AuthUser, error) {
	var model userpostgres.UserModel
	err := r.db.WithContext(ctx).First(&model, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return toAuthUser(&model), nil
}

func (r *Repository) MarkEmailVerified(ctx context.Context, userID uuid.UUID, verifiedAt time.Time) error {
	result := r.db.WithContext(ctx).Model(&userpostgres.UserModel{}).
		Where("id = ?", userID).
		Update("email_verified_at", verifiedAt)
	if result.Error != nil {
		return fmt.Errorf("mark email verified: %w", result.Error)
	}
	return nil
}

func (r *Repository) FindByProvider(ctx context.Context, provider, subject string) (*domain.AuthUser, error) {
	var identity AuthIdentityModel
	err := r.db.WithContext(ctx).
		Where("provider = ? AND provider_subject = ?", provider, subject).
		First(&identity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, fmt.Errorf("find auth identity: %w", err)
	}
	return r.GetByID(ctx, identity.UserID)
}

func (r *Repository) Link(ctx context.Context, userID uuid.UUID, provider, subject string) error {
	model := &AuthIdentityModel{
		UserID:          userID,
		Provider:        provider,
		ProviderSubject: subject,
	}
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		if isDuplicateKey(err) {
			return nil
		}
		return fmt.Errorf("link auth identity: %w", err)
	}
	return nil
}

func (r *Repository) CreateGoogleUser(ctx context.Context, email, name, provider, subject string, verifiedAt *time.Time) (*domain.AuthUser, error) {
	var user *domain.AuthUser
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := &Repository{db: tx}
		existing, err := repo.GetByEmail(ctx, email)
		if err != nil && !errors.Is(err, domain.ErrInvalidCredentials) {
			return err
		}
		if existing != nil {
			if err := repo.Link(ctx, existing.ID, provider, subject); err != nil {
				return err
			}
			if verifiedAt != nil && existing.EmailVerifiedAt == nil {
				if err := repo.MarkEmailVerified(ctx, existing.ID, *verifiedAt); err != nil {
					return err
				}
			}
			user, err = repo.GetByID(ctx, existing.ID)
			return err
		}

		model := &userpostgres.UserModel{
			Email: strings.ToLower(strings.TrimSpace(email)),
			Name:  strings.TrimSpace(name),
		}
		if verifiedAt != nil {
			model.EmailVerifiedAt = verifiedAt
		}
		if err := tx.Create(model).Error; err != nil {
			if isDuplicateKey(err) {
				return domain.ErrDuplicateEmail
			}
			return fmt.Errorf("create google user: %w", err)
		}
		if err := repo.Link(ctx, model.ID, provider, subject); err != nil {
			return err
		}
		user = toAuthUser(model)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) Store(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error {
	model := &RefreshTokenModel{
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
	}
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return fmt.Errorf("store refresh token: %w", err)
	}
	return nil
}

func (r *Repository) FindValid(ctx context.Context, tokenHash string, now time.Time) (uuid.UUID, error) {
	var model RefreshTokenModel
	err := r.db.WithContext(ctx).
		Where("token_hash = ? AND revoked_at IS NULL AND expires_at > ?", tokenHash, now).
		First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uuid.Nil, domain.ErrInvalidRefreshToken
		}
		return uuid.Nil, fmt.Errorf("find refresh token: %w", err)
	}
	return model.UserID, nil
}

func (r *Repository) Revoke(ctx context.Context, tokenHash string, revokedAt time.Time) error {
	result := r.db.WithContext(ctx).Model(&RefreshTokenModel{}).
		Where("token_hash = ?", tokenHash).
		Update("revoked_at", revokedAt)
	if result.Error != nil {
		return fmt.Errorf("revoke refresh token: %w", result.Error)
	}
	return nil
}

func (r *Repository) CreateVerificationToken(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error {
	model := &EmailVerificationTokenModel{
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
	}
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return fmt.Errorf("create verification token: %w", err)
	}
	return nil
}

func (r *Repository) Create(ctx context.Context, userID uuid.UUID, tokenHash string, expiresAt time.Time) error {
	return r.CreateVerificationToken(ctx, userID, tokenHash, expiresAt)
}

func (r *Repository) FindValidByHash(ctx context.Context, tokenHash string, now time.Time) (uuid.UUID, error) {
	var model EmailVerificationTokenModel
	err := r.db.WithContext(ctx).Where("token_hash = ?", tokenHash).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return uuid.Nil, domain.ErrInvalidVerificationToken
		}
		return uuid.Nil, fmt.Errorf("find verification token: %w", err)
	}
	if model.UsedAt != nil {
		return uuid.Nil, domain.ErrVerificationTokenUsed
	}
	if !model.ExpiresAt.After(now) {
		return uuid.Nil, domain.ErrVerificationTokenExpired
	}
	return model.UserID, nil
}

func (r *Repository) MarkUsed(ctx context.Context, tokenHash string, usedAt time.Time) error {
	result := r.db.WithContext(ctx).Model(&EmailVerificationTokenModel{}).
		Where("token_hash = ? AND used_at IS NULL", tokenHash).
		Update("used_at", usedAt)
	if result.Error != nil {
		return fmt.Errorf("mark verification token used: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return domain.ErrInvalidVerificationToken
	}
	return nil
}

func toAuthUser(model *userpostgres.UserModel) *domain.AuthUser {
	return &domain.AuthUser{
		ID:              model.ID,
		Email:           model.Email,
		Name:            model.Name,
		PasswordHash:    model.PasswordHash,
		EmailVerifiedAt: model.EmailVerifiedAt,
	}
}

func isDuplicateKey(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	return strings.Contains(strings.ToLower(err.Error()), "duplicate key")
}
