package postgres

import (
	"context"
	"time"

	"shopwise/retail/internal/resume_session/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	database *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{database: db}
}

func (r *Repository) CreateResumeToken(ctx context.Context, token *domain.ResumeToken) error {
	model := ResumeToken{
		ID:                  token.ID,
		SessionID:           token.SessionID,
		UserID:              token.UserID,
		TokenHash:           token.TokenHash,
		ExpiresAt:           token.ExpiresAt,
		ConsumedAt:          token.ConsumedAt,
		RevokedAt:           token.RevokedAt,
		IssuedContextHash:   token.IssuedContextHash,
		ConsumedContextHash: token.ConsumedContextHash,
		CreatedAt:           token.CreatedAt,
	}
	return r.database.WithContext(ctx).Create(&model).Error
}

func (r *Repository) GetResumeToken(ctx context.Context, hash string) (*domain.ResumeToken, error) {
	var model ResumeToken
	if err := r.database.WithContext(ctx).Where("token_hash = ?", hash).First(&model).Error; err != nil {
		return nil, err
	}
	return &domain.ResumeToken{
		ID:                  model.ID,
		SessionID:           model.SessionID,
		UserID:              model.UserID,
		TokenHash:           model.TokenHash,
		ExpiresAt:           model.ExpiresAt,
		ConsumedAt:          model.ConsumedAt,
		RevokedAt:           model.RevokedAt,
		IssuedContextHash:   model.IssuedContextHash,
		ConsumedContextHash: model.ConsumedContextHash,
		CreatedAt:           model.CreatedAt,
	}, nil
}

func (r *Repository) MarkTokenConsumed(ctx context.Context, id uuid.UUID, contextHash string) error {
	now := time.Now()
	updates := map[string]interface{}{
		"consumed_at": now,
	}
	if contextHash != "" {
		updates["consumed_context_hash"] = contextHash
	}
	return r.database.WithContext(ctx).Model(&ResumeToken{}).Where("id = ?", id).Updates(updates).Error
}

func (r *Repository) RevokeUnconsumedTokens(ctx context.Context, sessionID uuid.UUID) error {
	now := time.Now()
	return r.database.WithContext(ctx).
		Model(&ResumeToken{}).
		Where("session_id = ? AND consumed_at IS NULL AND revoked_at IS NULL", sessionID).
		Update("revoked_at", now).Error
}

func (r *Repository) CreateNotificationLog(ctx context.Context, log *domain.NotificationLog) error {
	model := NotificationLog{
		ID:           log.ID,
		SessionID:    log.SessionID,
		Provider:     log.Provider,
		Status:       log.Status,
		CreatedAt:    log.CreatedAt,
	}
	
	if log.ErrorDetails != nil {
		model.ErrorDetails = []byte(*log.ErrorDetails)
	}

	return r.database.WithContext(ctx).Create(&model).Error
}
