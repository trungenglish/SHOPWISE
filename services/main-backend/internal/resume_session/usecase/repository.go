package usecase

import (
	"context"

	"shopwise/retail/internal/resume_session/domain"
	"github.com/google/uuid"
)

type Repository interface {
	CreateResumeToken(ctx context.Context, token *domain.ResumeToken) error
	GetResumeToken(ctx context.Context, hash string) (*domain.ResumeToken, error)
	MarkTokenConsumed(ctx context.Context, id uuid.UUID, contextHash string) error
	RevokeUnconsumedTokens(ctx context.Context, sessionID uuid.UUID) error

	CreateNotificationLog(ctx context.Context, log *domain.NotificationLog) error
}
