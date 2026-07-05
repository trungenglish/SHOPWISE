package usecase

import (
	"context"
	"errors"
	"log/slog"

	"shopwise/apps/server/internal/platform/apperror"
	"shopwise/apps/server/internal/users/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Service) WithAccountDeletion(deps AccountDeletionDeps) *Service {
	s.accountDeletion = &deps
	return s
}

func (s *Service) DeleteAccount(ctx context.Context, userID uuid.UUID) error {
	if s.accountDeletion == nil {
		return apperror.Internal("account deletion is not configured", nil)
	}

	deps := s.accountDeletion
	var storageKeys []string

	err := deps.runTransaction(ctx, func(tx *gorm.DB) error {
		if err := deps.Notification.DeleteUserData(ctx, tx, userID); err != nil {
			return err
		}
		if err := deps.Wishlist.DeleteUserData(ctx, tx, userID); err != nil {
			return err
		}
		keys, err := deps.Ownership.DeleteUserData(ctx, tx, userID)
		if err != nil {
			return err
		}
		storageKeys = keys
		if err := deps.Advisor.DeleteUserData(ctx, tx, userID); err != nil {
			return err
		}
		if err := deps.Identity.DeleteUserData(ctx, tx, userID); err != nil {
			return err
		}
		return deps.Users.DeleteAccount(ctx, tx, userID)
	})
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return apperror.NotFound("user not found", err)
		}
		return apperror.Internal("failed to delete account", err)
	}

	if deps.Storage != nil {
		for _, key := range storageKeys {
			if delErr := deps.Storage.Delete(ctx, key); delErr != nil {
				s.log.Error("failed to delete stored document",
					slog.String("storage_key", key),
					slog.Any("error", delErr),
				)
			}
		}
	}

	return nil
}
