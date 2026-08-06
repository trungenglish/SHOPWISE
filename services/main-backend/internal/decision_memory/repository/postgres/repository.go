package postgres

import (
	"context"
	"fmt"
	"time"

	"shopwise/retail/internal/decision_memory/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	database *gorm.DB
}

func (r *Repository) BeginInteraction(
	ctx context.Context,
	interaction *domain.SessionInteraction,
) (*domain.SessionInteraction, bool, error) {
	var stored SessionInteraction
	acquired := false
	err := r.database.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("session_id = ? AND interaction_id = ?", interaction.SessionID, interaction.InteractionID).
			First(&stored)
		if result.Error == gorm.ErrRecordNotFound {
			candidate := SessionInteraction{
				SessionID: interaction.SessionID, InteractionID: interaction.InteractionID,
				PayloadHash: interaction.PayloadHash, Status: "pending",
				LeaseUntil: interaction.LeaseUntil, CreatedAt: interaction.CreatedAt,
				UpdatedAt: interaction.UpdatedAt,
			}
			created := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&candidate)
			if created.Error != nil {
				return fmt.Errorf("create session interaction: %w", created.Error)
			}
			if created.RowsAffected == 1 {
				stored = candidate
				acquired = true
				return nil
			}
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("session_id = ? AND interaction_id = ?", interaction.SessionID, interaction.InteractionID).
				First(&stored).Error; err != nil {
				return fmt.Errorf("reload session interaction: %w", err)
			}
		}
		if result.Error != nil {
			return fmt.Errorf("get session interaction: %w", result.Error)
		}
		if stored.PayloadHash != interaction.PayloadHash {
			return domain.ErrInteractionConflict
		}
		if stored.Status == "completed" {
			return nil
		}
		if stored.Status == "pending" && stored.LeaseUntil.After(time.Now().UTC()) {
			return domain.ErrInteractionInProgress
		}
		stored.Status = "pending"
		stored.LeaseUntil = interaction.LeaseUntil
		stored.UpdatedAt = interaction.UpdatedAt
		if err := tx.Save(&stored).Error; err != nil {
			return fmt.Errorf("resume session interaction: %w", err)
		}
		acquired = true
		return nil
	})
	if err != nil {
		return nil, false, err
	}
	return &domain.SessionInteraction{
		SessionID: stored.SessionID, InteractionID: stored.InteractionID,
		PayloadHash: stored.PayloadHash, Status: stored.Status,
		LeaseUntil: stored.LeaseUntil, ResponseEnvelope: string(stored.ResponseEnvelope),
		CreatedAt: stored.CreatedAt, UpdatedAt: stored.UpdatedAt,
	}, acquired, nil
}

func (r *Repository) FinishInteraction(
	ctx context.Context,
	sessionID uuid.UUID,
	interactionID uuid.UUID,
	status string,
	response []byte,
) error {
	updates := map[string]interface{}{
		"status": status, "response_envelope": response, "updated_at": time.Now().UTC(),
	}
	result := r.database.WithContext(ctx).Model(&SessionInteraction{}).
		Where("session_id = ? AND interaction_id = ?", sessionID, interactionID).
		Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("finish session interaction: %w", result.Error)
	}
	return nil
}

func NewRepository(database *gorm.DB) *Repository {
	return &Repository{database: database}
}

func (r *Repository) CreateSession(ctx context.Context, session *domain.DecisionSession) error {
	model := DecisionSession{
		ID:              session.ID,
		UserID:          session.UserID,
		AnonymousID:     session.AnonymousID,
		Title:           session.Title,
		Status:          session.Status,
		ParentSessionID: session.ParentSessionID,
		CreatedAt:       session.CreatedAt,
		UpdatedAt:       session.UpdatedAt,
	}
	if err := r.database.WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("create decision session: %w", err)
	}
	return nil
}

func (r *Repository) GetSession(ctx context.Context, id uuid.UUID) (*domain.DecisionSession, error) {
	var model DecisionSession
	if err := r.database.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, fmt.Errorf("get decision session: %w", err)
	}

	var messages []SessionMessage
	if err := r.database.WithContext(ctx).Where("session_id = ?", id).Order("created_at asc").Find(&messages).Error; err != nil {
		return nil, fmt.Errorf("get session messages: %w", err)
	}

	domainMessages := make([]domain.SessionMessage, len(messages))
	for i, m := range messages {
		domainMessages[i] = domain.SessionMessage{
			ID:             m.ID,
			SessionID:      m.SessionID,
			Role:           m.Role,
			Content:        m.Content,
			ReasoningGraph: string(m.ReasoningGraph),
			PinnedProducts: string(m.PinnedProducts),
			CreatedAt:      m.CreatedAt,
		}
	}

	return &domain.DecisionSession{
		ID:              model.ID,
		UserID:          model.UserID,
		AnonymousID:     model.AnonymousID,
		Title:           model.Title,
		Status:          model.Status,
		ParentSessionID: model.ParentSessionID,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		Messages:        domainMessages,
	}, nil
}

func (r *Repository) ListSessions(ctx context.Context, userID *uuid.UUID, anonymousID *string, limit, offset int) ([]domain.DecisionSession, error) {
	var models []DecisionSession
	query := r.database.WithContext(ctx)

	if userID != nil {
		query = query.Where("user_id = ?", userID)
	} else if anonymousID != nil {
		query = query.Where("anonymous_id = ?", anonymousID).Where("user_id IS NULL")
	} else {
		return nil, fmt.Errorf("either userID or anonymousID must be provided")
	}

	if err := query.Order("updated_at desc").Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, fmt.Errorf("list decision sessions: %w", err)
	}

	sessions := make([]domain.DecisionSession, len(models))
	for i, m := range models {
		sessions[i] = domain.DecisionSession{
			ID:              m.ID,
			UserID:          m.UserID,
			AnonymousID:     m.AnonymousID,
			Title:           m.Title,
			Status:          m.Status,
			ParentSessionID: m.ParentSessionID,
			CreatedAt:       m.CreatedAt,
			UpdatedAt:       m.UpdatedAt,
		}
	}
	return sessions, nil
}

func (r *Repository) UpdateSession(ctx context.Context, session *domain.DecisionSession, clientTimestamp time.Time) error {
	return r.database.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var current DecisionSession
		if err := tx.Where("id = ?", session.ID).First(&current).Error; err != nil {
			return fmt.Errorf("find session for update: %w", err)
		}

		if clientTimestamp.Before(current.UpdatedAt) {
			return fmt.Errorf("conflict: client timestamp %v is older than server timestamp %v", clientTimestamp, current.UpdatedAt)
		}

		updates := map[string]interface{}{
			"title":      session.Title,
			"status":     session.Status,
			"updated_at": time.Now().UTC(),
		}
		if err := tx.Model(&current).Updates(updates).Error; err != nil {
			return fmt.Errorf("update session: %w", err)
		}

		session.UpdatedAt = updates["updated_at"].(time.Time)
		return nil
	})
}

func (r *Repository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	return r.database.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("session_id = ?", id).Delete(&SessionMessage{}).Error; err != nil {
			return fmt.Errorf("delete session messages: %w", err)
		}
		if err := tx.Where("id = ?", id).Delete(&DecisionSession{}).Error; err != nil {
			return fmt.Errorf("delete decision session: %w", err)
		}
		return nil
	})
}

func (r *Repository) CountAnonymousSessions(ctx context.Context, anonymousID string) (int64, error) {
	var count int64
	if err := r.database.WithContext(ctx).Model(&DecisionSession{}).Where("anonymous_id = ? AND user_id IS NULL", anonymousID).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("count anonymous sessions: %w", err)
	}
	return count, nil
}

func (r *Repository) DeleteOldestAnonymousSession(ctx context.Context, anonymousID string) error {
	var oldest DecisionSession
	if err := r.database.WithContext(ctx).Where("anonymous_id = ? AND user_id IS NULL", anonymousID).Order("created_at asc").First(&oldest).Error; err != nil {
		return fmt.Errorf("find oldest anonymous session: %w", err)
	}
	return r.DeleteSession(ctx, oldest.ID)
}

func (r *Repository) AddMessage(ctx context.Context, msg *domain.SessionMessage) error {
	model := SessionMessage{
		ID:             msg.ID,
		SessionID:      msg.SessionID,
		Role:           msg.Role,
		Content:        msg.Content,
		ReasoningGraph: []byte(msg.ReasoningGraph),
		PinnedProducts: []byte(msg.PinnedProducts),
		CreatedAt:      msg.CreatedAt,
	}
	if err := r.database.WithContext(ctx).Create(&model).Error; err != nil {
		return fmt.Errorf("add session message: %w", err)
	}
	return nil
}

func (r *Repository) ListPreferences(ctx context.Context, userID uuid.UUID) ([]domain.UserPreference, error) {
	var models []ExtractedPreference
	if err := r.database.WithContext(ctx).Where("user_id = ?", userID).Order("created_at desc").Find(&models).Error; err != nil {
		return nil, fmt.Errorf("list preferences: %w", err)
	}

	prefs := make([]domain.UserPreference, len(models))
	for i, m := range models {
		prefs[i] = domain.UserPreference{
			ID:              m.ID,
			UserID:          m.UserID,
			Category:        m.Category,
			Value:           string(m.Value),
			SourceSessionID: m.SourceSessionID,
			CreatedAt:       m.CreatedAt,
			UpdatedAt:       m.UpdatedAt,
		}
	}
	return prefs, nil
}

func (r *Repository) UpdatePreference(ctx context.Context, pref *domain.UserPreference) error {
	model := ExtractedPreference{
		ID:              pref.ID,
		UserID:          pref.UserID,
		Category:        pref.Category,
		Value:           []byte(pref.Value),
		SourceSessionID: pref.SourceSessionID,
		CreatedAt:       pref.CreatedAt,
		UpdatedAt:       time.Now().UTC(),
	}
	if err := r.database.WithContext(ctx).Save(&model).Error; err != nil {
		return fmt.Errorf("update preference: %w", err)
	}
	pref.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *Repository) DeletePreference(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	if err := r.database.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&ExtractedPreference{}).Error; err != nil {
		return fmt.Errorf("delete preference: %w", err)
	}
	return nil
}

func (r *Repository) BranchSession(ctx context.Context, id uuid.UUID, newTitle string) (*domain.DecisionSession, error) {
	var newSession domain.DecisionSession
	err := r.database.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var oldSession DecisionSession
		if err := tx.Where("id = ?", id).First(&oldSession).Error; err != nil {
			return fmt.Errorf("find session to branch: %w", err)
		}

		newID := uuid.New()
		now := time.Now().UTC()

		newDbSession := DecisionSession{
			ID:              newID,
			UserID:          oldSession.UserID,
			AnonymousID:     oldSession.AnonymousID,
			Title:           newTitle,
			Status:          "active",
			ParentSessionID: &oldSession.ID,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		if err := tx.Create(&newDbSession).Error; err != nil {
			return fmt.Errorf("create branched session: %w", err)
		}

		var messages []SessionMessage
		if err := tx.Where("session_id = ?", id).Order("created_at asc").Find(&messages).Error; err != nil {
			return fmt.Errorf("find messages for branch: %w", err)
		}

		for _, msg := range messages {
			newMsg := SessionMessage{
				ID:             uuid.New(),
				SessionID:      newID,
				Role:           msg.Role,
				Content:        msg.Content,
				ReasoningGraph: msg.ReasoningGraph,
				PinnedProducts: msg.PinnedProducts,
				CreatedAt:      msg.CreatedAt,
			}
			if err := tx.Create(&newMsg).Error; err != nil {
				return fmt.Errorf("copy message for branch: %w", err)
			}
		}

		newSession = domain.DecisionSession{
			ID:              newDbSession.ID,
			UserID:          newDbSession.UserID,
			AnonymousID:     newDbSession.AnonymousID,
			Title:           newDbSession.Title,
			Status:          newDbSession.Status,
			ParentSessionID: newDbSession.ParentSessionID,
			CreatedAt:       newDbSession.CreatedAt,
			UpdatedAt:       newDbSession.UpdatedAt,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &newSession, nil
}

func (r *Repository) ArchiveInactiveSessions(ctx context.Context, before time.Time) (int64, error) {
	result := r.database.WithContext(ctx).Model(&DecisionSession{}).
		Where("status = ? AND updated_at < ?", "active", before).
		Update("status", "archived")

	if result.Error != nil {
		return 0, fmt.Errorf("archive inactive sessions: %w", result.Error)
	}
	return result.RowsAffected, nil
}
