package accessories

import (
	"context"
	"fmt"
	"time"

	"shopwise/retail/internal/platform/apperror"

	"github.com/google/uuid"
)

const (
	maximumRecommendationProducts = 10
	freshOfferAge                 = 6 * time.Hour
	maximumOfferAge               = 72 * time.Hour
)

type RecommendationRepository interface {
	LoadRecommendationSnapshot(context.Context, []uuid.UUID) (RecommendationSnapshot, error)
}

type Service struct {
	repository RecommendationRepository
	now        func() time.Time
}

func NewService(repository RecommendationRepository, now func() time.Time) *Service {
	if now == nil {
		now = time.Now
	}
	return &Service{repository: repository, now: now}
}

func (service *Service) Recommend(ctx context.Context, rawProductIDs []string) ([]Recommendation, error) {
	productIDs, err := parseProductIDs(rawProductIDs)
	if err != nil {
		return nil, err
	}
	snapshot, err := service.repository.LoadRecommendationSnapshot(ctx, productIDs)
	if err != nil {
		return nil, apperror.Internal("failed to load accessory recommendations", err)
	}
	if len(snapshot.Products) != len(productIDs) {
		return nil, apperror.Validation("one or more product_ids do not exist", nil)
	}

	now := service.now().UTC()
	result := make([]Recommendation, 0, len(snapshot.Candidates))
	for _, candidate := range snapshot.Candidates {
		state := offerState(candidate.Offer, now)
		if state == OfferUnavailable && now.Sub(candidate.Offer.FetchedAt) > maximumOfferAge {
			continue
		}
		compatibleIDs, reason, status := matchCandidate(candidate, snapshot.Products)
		if len(compatibleIDs) == 0 {
			continue
		}
		result = append(result, Recommendation{
			Accessory: candidate.Accessory, Offer: candidate.Offer,
			CompatibleProductIDs: compatibleIDs, CompatibilityReason: reason,
			CompatibilityStatus: status, OfferState: state,
			CheckoutAvailable: state == OfferFresh && candidate.Offer.InStock,
		})
	}
	return result, nil
}

func parseProductIDs(values []string) ([]uuid.UUID, error) {
	if len(values) == 0 {
		return nil, apperror.Validation("at least one product_id is required", nil)
	}
	if len(values) > maximumRecommendationProducts {
		return nil, apperror.Validation("at most 10 product_ids are allowed", nil)
	}
	ids := make([]uuid.UUID, 0, len(values))
	seen := make(map[uuid.UUID]struct{}, len(values))
	for index, value := range values {
		id, err := uuid.Parse(value)
		if err != nil {
			return nil, apperror.Validation(fmt.Sprintf("product_ids[%d] must be a valid UUID", index), err)
		}
		if _, exists := seen[id]; exists {
			continue
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}
	return ids, nil
}

func offerState(offer RetailerOffer, now time.Time) OfferState {
	if !offer.InStock || offer.Price <= 0 {
		return OfferUnavailable
	}
	age := now.Sub(offer.FetchedAt)
	if age <= freshOfferAge {
		return OfferFresh
	}
	if age <= maximumOfferAge {
		return OfferStale
	}
	return OfferUnavailable
}

func matchCandidate(candidate Candidate, products []LaptopTarget) ([]uuid.UUID, string, VerificationStatus) {
	compatible := make([]uuid.UUID, 0, len(products))
	reason := ""
	status := VerificationCategory
	for _, product := range products {
		compatibility, matched := bestCompatibility(candidate.Compatibilities, product)
		if !matched {
			continue
		}
		if candidate.Accessory.Kind == KindUpgrade &&
			(!product.UpgradeProfileVerified || compatibility.ProductID == nil || compatibility.VerificationStatus != VerificationModel) {
			continue
		}
		compatible = append(compatible, product.ID)
		if compatibility.ProductID != nil || reason == "" {
			reason = compatibility.Reason
			status = compatibility.VerificationStatus
		}
	}
	return compatible, reason, status
}

func bestCompatibility(compatibilities []AccessoryCompatibility, product LaptopTarget) (AccessoryCompatibility, bool) {
	var categoryMatch AccessoryCompatibility
	foundCategory := false
	for _, compatibility := range compatibilities {
		if compatibility.ProductID != nil && *compatibility.ProductID == product.ID {
			return compatibility, true
		}
		if compatibility.ProductID == nil && compatibility.LaptopCategory == product.Category {
			categoryMatch = compatibility
			foundCategory = true
		}
	}
	return categoryMatch, foundCategory
}

func DefaultLaptopCategories(accessoryCategory string) []string {
	switch accessoryCategory {
	case "keyboard", "cooling_pad":
		return []string{"gaming"}
	case "monitor":
		return []string{"gaming", "creative"}
	case "card_reader", "external_ssd":
		return []string{"creative"}
	case "dock", "hub", "webcam", "bag":
		return []string{"creative", "business", "work", "student", "study", "general"}
	default:
		return []string{"gaming", "creative", "business", "work", "student", "study", "general"}
	}
}
