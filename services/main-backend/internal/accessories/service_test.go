package accessories

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

type recommendationRepositoryStub struct {
	snapshot RecommendationSnapshot
}

func (stub recommendationRepositoryStub) LoadRecommendationSnapshot(
	context.Context,
	[]uuid.UUID,
) (RecommendationSnapshot, error) {
	return stub.snapshot, nil
}

func TestRecommendAppliesCategoryDefaultsAndFreshness(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, time.August, 5, 12, 0, 0, 0, time.UTC)
	productID := uuid.New()
	accessoryID := uuid.New()
	service := NewService(recommendationRepositoryStub{snapshot: RecommendationSnapshot{
		Products: []LaptopTarget{{ID: productID, Category: "gaming"}},
		Candidates: []Candidate{{
			Accessory:       Accessory{ID: accessoryID, Kind: KindExternal, Category: "mouse", Name: "Gaming mouse"},
			Offer:           RetailerOffer{ID: uuid.New(), AccessoryID: accessoryID, Price: 990_000, InStock: true, FetchedAt: now.Add(-time.Hour)},
			Compatibilities: []AccessoryCompatibility{{AccessoryID: accessoryID, LaptopCategory: "gaming", Reason: "Suitable for gaming", VerificationStatus: VerificationCategory}},
		}},
	}}, func() time.Time { return now })

	items, err := service.Recommend(context.Background(), []string{productID.String()})
	if err != nil {
		t.Fatalf("Recommend() error = %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("len(items) = %d, want 1", len(items))
	}
	if items[0].OfferState != OfferFresh || !items[0].CheckoutAvailable {
		t.Fatalf("offer state/checkout = %s/%v, want fresh/true", items[0].OfferState, items[0].CheckoutAvailable)
	}
	if len(items[0].CompatibleProductIDs) != 1 || items[0].CompatibleProductIDs[0] != productID {
		t.Fatalf("compatible product ids = %v", items[0].CompatibleProductIDs)
	}
}

func TestRecommendHidesUnverifiedUpgradeAndExpiredOffer(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, time.August, 5, 12, 0, 0, 0, time.UTC)
	productID := uuid.New()
	upgradeID := uuid.New()
	expiredID := uuid.New()
	service := NewService(recommendationRepositoryStub{snapshot: RecommendationSnapshot{
		Products: []LaptopTarget{{ID: productID, Category: "gaming", UpgradeProfileVerified: false}},
		Candidates: []Candidate{
			{
				Accessory:       Accessory{ID: upgradeID, Kind: KindUpgrade, Category: "ram", Name: "Laptop RAM"},
				Offer:           RetailerOffer{ID: uuid.New(), AccessoryID: upgradeID, InStock: true, Price: 1_200_000, FetchedAt: now.Add(-time.Hour)},
				Compatibilities: []AccessoryCompatibility{{AccessoryID: upgradeID, LaptopCategory: "gaming", VerificationStatus: VerificationCategory}},
			},
			{
				Accessory:       Accessory{ID: expiredID, Kind: KindExternal, Category: "mouse", Name: "Old mouse"},
				Offer:           RetailerOffer{ID: uuid.New(), AccessoryID: expiredID, InStock: true, Price: 500_000, FetchedAt: now.Add(-73 * time.Hour)},
				Compatibilities: []AccessoryCompatibility{{AccessoryID: expiredID, LaptopCategory: "gaming", VerificationStatus: VerificationCategory}},
			},
		},
	}}, func() time.Time { return now })

	items, err := service.Recommend(context.Background(), []string{productID.String()})
	if err != nil {
		t.Fatalf("Recommend() error = %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("items = %#v, want no unverified or expired items", items)
	}
}

func TestDefaultLaptopCategoriesCoverAllSupportedGroups(t *testing.T) {
	t.Parallel()

	want := map[string]bool{
		"gaming": true, "creative": true, "business": true, "work": true,
		"student": true, "study": true, "general": true,
	}
	got := DefaultLaptopCategories("mouse")
	if len(got) != len(want) {
		t.Fatalf("categories = %v", got)
	}
	for _, category := range got {
		delete(want, category)
	}
	if len(want) != 0 {
		t.Fatalf("missing categories = %v", want)
	}
}

func TestRecommendPrefersModelOverrideToCategoryDefault(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	productID := uuid.New()
	accessoryID := uuid.New()
	service := NewService(recommendationRepositoryStub{snapshot: RecommendationSnapshot{
		Products: []LaptopTarget{{ID: productID, Category: "gaming"}},
		Candidates: []Candidate{{
			Accessory: Accessory{ID: accessoryID, Kind: KindExternal, Category: "dock"},
			Offer:     RetailerOffer{ID: uuid.New(), AccessoryID: accessoryID, Price: 1, InStock: true, FetchedAt: now},
			Compatibilities: []AccessoryCompatibility{
				{AccessoryID: accessoryID, LaptopCategory: "gaming", Reason: "category default", VerificationStatus: VerificationCategory},
				{AccessoryID: accessoryID, ProductID: &productID, Reason: "verified model override", VerificationStatus: VerificationModel},
			},
		}},
	}}, func() time.Time { return now })

	items, err := service.Recommend(context.Background(), []string{productID.String()})
	if err != nil {
		t.Fatalf("Recommend() error = %v", err)
	}
	if len(items) != 1 || items[0].CompatibilityReason != "verified model override" || items[0].CompatibilityStatus != VerificationModel {
		t.Fatalf("recommendation = %#v", items)
	}
}

func TestRecommendRejectsInvalidProductList(t *testing.T) {
	t.Parallel()

	service := NewService(recommendationRepositoryStub{}, time.Now)
	if _, err := service.Recommend(context.Background(), []string{"not-a-uuid"}); err == nil {
		t.Fatal("invalid UUID error = nil")
	}
	tooMany := make([]string, 11)
	for index := range tooMany {
		tooMany[index] = uuid.NewString()
	}
	if _, err := service.Recommend(context.Background(), tooMany); err == nil {
		t.Fatal("more than 10 product IDs error = nil")
	}
}
