package usecase

import (
	"context"
	"errors"
	"fmt"

	"shopwise/retail/internal/decision_memory/domain"
	"shopwise/retail/internal/decision_memory/repository"
)

var (
	ErrCategoryMismatch  = errors.New("category mismatch")
	ErrWorkspaceNotFound = errors.New("workspace not found")
)

type ComparisonUsecase interface {
	AddProductToWorkspace(ctx context.Context, workspaceID string, product *domain.ComparisonProductDetail) (*domain.ComparisonWorkspace, error)
	RemoveProductFromWorkspace(ctx context.Context, workspaceID string, productID string) (*domain.ComparisonWorkspace, error)
	ReplaceProductInWorkspace(ctx context.Context, workspaceID string, oldProductID string, product *domain.ComparisonProductDetail) (*domain.ComparisonWorkspace, error)
	ComputeHighlights(ctx context.Context, products []*domain.ComparisonProductDetail) []domain.ComparisonDifferenceHighlight
	MarkCheckoutReady(ctx context.Context, workspaceID string, productID string) (*domain.ComparisonWorkspace, error)
}

type comparisonUsecase struct {
	repo repository.ComparisonRepository
}

func NewComparisonUsecase(repo repository.ComparisonRepository) ComparisonUsecase {
	return &comparisonUsecase{
		repo: repo,
	}
}

// AddProductToWorkspace adds a product to the workspace, enforcing category limits and capacity.
func (u *comparisonUsecase) AddProductToWorkspace(ctx context.Context, workspaceID string, product *domain.ComparisonProductDetail) (*domain.ComparisonWorkspace, error) {
	ws, err := u.repo.GetComparisonWorkspace(ctx, workspaceID)
	if err != nil {
		if errors.Is(err, ErrWorkspaceNotFound) {
			// Initialize new workspace
			ws = &domain.ComparisonWorkspace{
				WorkspaceID: workspaceID,
				CategoryID:  product.Category,
				State:       domain.WorkspaceStateAddingProduct,
			}
		} else {
			return nil, err
		}
	}

	// T008: Category Validation
	if ws.CategoryID != "" && ws.CategoryID != product.Category {
		return ws, fmt.Errorf("%w: expected %s but got %s", ErrCategoryMismatch, ws.CategoryID, product.Category)
	}

	// If pending replacement, we cannot just add normally. The user must replace.
	if ws.State == domain.WorkspaceStatePendingReplacement {
		return ws, nil // or return an error indicating they must resolve the pending state
	}

	// T008: 4-product limit and T009: 5th product handling
	if len(ws.ProductIDs) >= 4 {
		// Attempting to add a 5th product
		ws.State = domain.WorkspaceStatePendingReplacement
		err = u.repo.SaveComparisonWorkspace(ctx, ws)
		return ws, err
	}

	// Add product normally
	ws.ProductIDs = append(ws.ProductIDs, product.ProductID)
	ws.State = domain.WorkspaceStateActive

	err = u.repo.SaveComparisonWorkspace(ctx, ws)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (u *comparisonUsecase) RemoveProductFromWorkspace(ctx context.Context, workspaceID string, productID string) (*domain.ComparisonWorkspace, error) {
	ws, err := u.repo.GetComparisonWorkspace(ctx, workspaceID)
	if err != nil {
		return nil, err
	}
	
	newProductIDs := make([]string, 0, len(ws.ProductIDs))
	for _, id := range ws.ProductIDs {
		if id != productID {
			newProductIDs = append(newProductIDs, id)
		}
	}
	ws.ProductIDs = newProductIDs
	
	if len(ws.ProductIDs) == 0 {
		ws.State = domain.WorkspaceStateInit
		ws.CategoryID = ""
	} else if ws.State == domain.WorkspaceStatePendingReplacement {
		if len(ws.ProductIDs) < 4 {
			ws.State = domain.WorkspaceStateActive
		}
	}
	
	err = u.repo.SaveComparisonWorkspace(ctx, ws)
	return ws, err
}

func (u *comparisonUsecase) ReplaceProductInWorkspace(ctx context.Context, workspaceID string, oldProductID string, product *domain.ComparisonProductDetail) (*domain.ComparisonWorkspace, error) {
	ws, err := u.repo.GetComparisonWorkspace(ctx, workspaceID)
	if err != nil {
		return nil, err
	}
	
	if ws.CategoryID != "" && ws.CategoryID != product.Category {
		return ws, fmt.Errorf("%w: expected %s but got %s", ErrCategoryMismatch, ws.CategoryID, product.Category)
	}

	replaced := false
	for i, id := range ws.ProductIDs {
		if id == oldProductID {
			ws.ProductIDs[i] = product.ProductID
			replaced = true
			break
		}
	}
	
	if !replaced {
		return nil, fmt.Errorf("old product not found in workspace")
	}
	
	ws.State = domain.WorkspaceStateActive
	err = u.repo.SaveComparisonWorkspace(ctx, ws)
	return ws, err
}

func (u *comparisonUsecase) ComputeHighlights(ctx context.Context, products []*domain.ComparisonProductDetail) []domain.ComparisonDifferenceHighlight {
	highlights := make([]domain.ComparisonDifferenceHighlight, 0)
	keys := make(map[string]bool)
	for _, p := range products {
		for k := range p.Specifications {
			keys[k] = true
		}
	}
	
	for key := range keys {
		h := domain.ComparisonDifferenceHighlight{
			SpecificationKey: key,
			Highlights:       make(map[string]domain.HighlightType),
		}
		
		var firstVal string
		allSame := true
		for i, p := range products {
			val, ok := p.Specifications[key]
			if !ok {
				h.Highlights[p.ProductID] = domain.HighlightMissing
				allSame = false
				continue
			}
			if i == 0 {
				firstVal = val
			} else if val != firstVal {
				allSame = false
			}
		}
		
		for _, p := range products {
			if h.Highlights[p.ProductID] == domain.HighlightMissing {
				continue
			}
			if allSame {
				h.Highlights[p.ProductID] = domain.HighlightEqual
			} else {
				h.Highlights[p.ProductID] = domain.HighlightCategoryAdvantage
			}
		}
		highlights = append(highlights, h)
	}
	return highlights
}

func (u *comparisonUsecase) MarkCheckoutReady(ctx context.Context, workspaceID string, productID string) (*domain.ComparisonWorkspace, error) {
	ws, err := u.repo.GetComparisonWorkspace(ctx, workspaceID)
	if err != nil {
		return nil, err
	}

	found := false
	for _, id := range ws.ProductIDs {
		if id == productID {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("product %s not found in workspace %s", productID, workspaceID)
	}

	ws.State = domain.WorkspaceStateCompleted
	err = u.repo.SaveComparisonWorkspace(ctx, ws)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
