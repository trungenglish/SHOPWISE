package repository

import (
	"context"

	"shopwise/retail/internal/decision_memory/domain"
)

type ComparisonRepository interface {
	SaveComparisonWorkspace(ctx context.Context, ws *domain.ComparisonWorkspace) error
	GetComparisonWorkspace(ctx context.Context, workspaceID string) (*domain.ComparisonWorkspace, error)
}
