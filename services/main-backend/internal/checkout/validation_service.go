package checkout

import (
	"context"
	"encoding/json"
	"time"

	dmDomain "shopwise/retail/internal/decision_memory/domain"
	dmUsecase "shopwise/retail/internal/decision_memory/usecase"

	"github.com/google/uuid"
)

// RetailSDK represents the external retail service we validate against
type RetailSDK interface {
	CheckInventory(ctx context.Context, productID string) (int, error)
	GetPrice(ctx context.Context, productID string) (int, error)
	VerifyWarranty(ctx context.Context, productID string) (bool, error)
}

type ValidationService struct {
	retailSDK RetailSDK
	dmService *dmUsecase.Service
}

func NewValidationService(retailSDK RetailSDK, dmService *dmUsecase.Service) *ValidationService {
	return &ValidationService{
		retailSDK: retailSDK,
		dmService: dmService,
	}
}

func (s *ValidationService) Validate(ctx context.Context, sessionIDStr string, productID string, profile *CustomerProfile) (*CheckoutReadinessState, error) {
	sessionID, err := uuid.Parse(sessionIDStr)
	if err != nil {
		return nil, err
	}

	state := &CheckoutReadinessState{
		SessionID:       sessionIDStr,
		CartVersion:     uuid.New().String(),
		Status:          StatusReady,
		LastValidatedAt: time.Now().UTC(),
		ValidationChecklist: []ValidationItem{},
	}

	// 1. Basic Readiness: Inventory, Price, Warranty
	inventory, err := s.retailSDK.CheckInventory(ctx, productID)
	if err != nil || inventory <= 0 {
		state.Status = StatusOutOfStock
		state.ValidationChecklist = append(state.ValidationChecklist, ValidationItem{
			Type:    ValidationInventory,
			Passed:  false,
			Message: "Product is out of stock.",
		})
	} else {
		state.ValidationChecklist = append(state.ValidationChecklist, ValidationItem{
			Type:    ValidationInventory,
			Passed:  true,
			Message: "Product is in stock.",
		})
	}

	price, err := s.retailSDK.GetPrice(ctx, productID)
	if err != nil {
		state.Status = StatusRequiresUserAction
		state.ValidationChecklist = append(state.ValidationChecklist, ValidationItem{
			Type:    ValidationPrice,
			Passed:  false,
			Message: "Could not confirm price.",
		})
	} else {
		state.ValidationChecklist = append(state.ValidationChecklist, ValidationItem{
			Type:    ValidationPrice,
			Passed:  true,
			Message: "Price confirmed.",
		})
		state.ProductSnapshot = &ProductSnapshot{
			ProductID: productID,
			Name:      "Product Name", // Mock name
			PriceVND:  price,
		}
	}

	warrantyOk, _ := s.retailSDK.VerifyWarranty(ctx, productID)
	if !warrantyOk {
		state.ValidationChecklist = append(state.ValidationChecklist, ValidationItem{
			Type:    ValidationInventory,
			Passed:  false,
			Message: "Warranty information could not be verified.",
		})
		if state.Status == StatusReady {
			state.Status = StatusMissingInformation
		}
	}

	// Persist to Decision Memory
	stateJSON, _ := json.Marshal(state)
	msg := &dmDomain.SessionMessage{
		ID:             uuid.New(),
		SessionID:      sessionID,
		Role:           "system",
		Content:        "Checkout Readiness State Updated",
		ReasoningGraph: string(stateJSON), // Store state in ReasoningGraph for now
		CreatedAt:      time.Now().UTC(),
	}
	// Note: You would normally access the Repo through the service if AddMessage was exposed on the Service struct
	// s.dmService.repo.AddMessage(ctx, msg)
	_ = msg

	return state, nil
}
