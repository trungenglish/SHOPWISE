package accessories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	platformmodel "shopwise/retail/internal/platform/database/model"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct{ database *gorm.DB }

func NewRepository(database *gorm.DB) *Repository { return &Repository{database: database} }

func Migrate(database *gorm.DB) error {
	if err := database.AutoMigrate(
		&AccessoryModel{}, &RetailerOfferModel{}, &AccessoryCompatibilityModel{},
		&LaptopCompatibilityProfileModel{}, &RetailerSyncStateModel{},
	); err != nil {
		return err
	}

	productID := uuid.MustParse("7cc26fb0-ace4-4e35-b6c5-fc3e13623546")
	fetchedAt := time.Date(2026, time.August, 5, 0, 0, 0, 0, time.UTC)
	seeds := []struct {
		accessory AccessoryModel
		offer     RetailerOfferModel
	}{
		{AccessoryModel{ID: uuid.MustParse("30000000-0000-4000-8000-000000000001"), Kind: string(KindExternal), Category: "mouse", Name: "Gaming Mouse", Specifications: datatypes.JSON([]byte(`{}`))}, RetailerOfferModel{ID: uuid.MustParse("10000000-0000-4000-8000-000000000001"), Retailer: "shopwise", RetailerProductID: "ROG-BUNDLE-MOUSE", SourceURL: "https://shopwise.local/offers/rog-mouse", Price: 0, OriginalPrice: 1_500_000, InStock: true, FetchedAt: fetchedAt}},
		{AccessoryModel{ID: uuid.MustParse("30000000-0000-4000-8000-000000000002"), Kind: string(KindExternal), Category: "bag", Name: "Anti-shock Bag", Specifications: datatypes.JSON([]byte(`{}`))}, RetailerOfferModel{ID: uuid.MustParse("10000000-0000-4000-8000-000000000002"), Retailer: "shopwise", RetailerProductID: "ROG-BUNDLE-BAG", SourceURL: "https://shopwise.local/offers/rog-bag", Price: 0, OriginalPrice: 800_000, InStock: true, FetchedAt: fetchedAt}},
		{AccessoryModel{ID: uuid.MustParse("30000000-0000-4000-8000-000000000003"), Kind: string(KindExternal), Category: "warranty", Name: "2-Year Accidental Damage Warranty", Specifications: datatypes.JSON([]byte(`{}`))}, RetailerOfferModel{ID: uuid.MustParse("10000000-0000-4000-8000-000000000003"), Retailer: "shopwise", RetailerProductID: "ROG-BUNDLE-WARRANTY", SourceURL: "https://shopwise.local/offers/rog-warranty", Price: 1_000_000, OriginalPrice: 2_000_000, InStock: true, FetchedAt: fetchedAt}},
	}
	for _, seed := range seeds {
		if err := database.Where(AccessoryModel{ID: seed.accessory.ID}).Attrs(seed.accessory).FirstOrCreate(&AccessoryModel{}).Error; err != nil {
			return err
		}
		seed.offer.AccessoryID = seed.accessory.ID
		if err := database.Where(RetailerOfferModel{ID: seed.offer.ID}).Attrs(seed.offer).FirstOrCreate(&RetailerOfferModel{}).Error; err != nil {
			return err
		}
		compatibility := AccessoryCompatibilityModel{ID: uuid.New(), AccessoryID: seed.accessory.ID, ProductID: &productID, Reason: "Included in the ROG RTX 5070 bundle", VerificationStatus: string(VerificationModel)}
		if err := database.Where("accessory_id = ? AND product_id = ?", seed.accessory.ID, productID).Attrs(compatibility).FirstOrCreate(&AccessoryCompatibilityModel{}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (repository *Repository) UpsertPhongVuCatalog(ctx context.Context, products []SyncedProduct, fetchedAt time.Time) error {
	return repository.database.WithContext(ctx).Transaction(func(transaction *gorm.DB) error {
		for _, product := range products {
			var existingOffer RetailerOfferModel
			err := transaction.Where("retailer = ? AND retailer_product_id = ?", "phongvu", product.RetailerProductID).First(&existingOffer).Error
			accessoryID := existingOffer.AccessoryID
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return fmt.Errorf("find retailer offer: %w", err)
				}
				accessoryID = uuid.New()
			}
			accessory := AccessoryModel{
				ID: accessoryID, Kind: string(KindExternal), Category: product.Category,
				Name: product.Name, Brand: product.Brand, ImageURL: product.ImageURL,
				Specifications: datatypes.JSON([]byte("{}")),
			}
			if err := transaction.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoUpdates: clause.AssignmentColumns([]string{"kind", "category", "name", "brand", "image_url", "specifications"})}).Create(&accessory).Error; err != nil {
				return fmt.Errorf("upsert accessory: %w", err)
			}
			offer := RetailerOfferModel{
				ID: uuid.New(), AccessoryID: accessoryID, Retailer: "phongvu",
				RetailerProductID: product.RetailerProductID, SourceURL: product.SourceURL,
				Price: product.Price, OriginalPrice: product.OriginalPrice, InStock: product.InStock, FetchedAt: fetchedAt,
			}
			if existingOffer.ID != uuid.Nil {
				offer.ID = existingOffer.ID
			}
			if err := transaction.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "retailer"}, {Name: "retailer_product_id"}}, DoUpdates: clause.AssignmentColumns([]string{"accessory_id", "source_url", "price", "original_price", "in_stock", "fetched_at"})}).Create(&offer).Error; err != nil {
				return fmt.Errorf("upsert retailer offer: %w", err)
			}
			for _, laptopCategory := range DefaultLaptopCategories(product.Category) {
				compatibility := AccessoryCompatibilityModel{
					ID: uuid.New(), AccessoryID: accessoryID, LaptopCategory: laptopCategory,
					Reason: compatibilityReason(product.Category, laptopCategory), VerificationStatus: string(VerificationCategory),
				}
				var count int64
				if err := transaction.Model(&AccessoryCompatibilityModel{}).Where("accessory_id = ? AND laptop_category = ? AND product_id IS NULL", accessoryID, laptopCategory).Count(&count).Error; err != nil {
					return fmt.Errorf("find category compatibility: %w", err)
				}
				if count == 0 {
					if err := transaction.Create(&compatibility).Error; err != nil {
						return fmt.Errorf("create category compatibility: %w", err)
					}
				}
			}
		}
		state := RetailerSyncStateModel{Retailer: "phongvu", LastRunAt: &fetchedAt, LastSuccessAt: &fetchedAt, LastError: ""}
		return transaction.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "retailer"}}, DoUpdates: clause.AssignmentColumns([]string{"last_run_at", "last_success_at", "last_error"})}).Create(&state).Error
	})
}

func (repository *Repository) MarkPhongVuSyncFailure(ctx context.Context, runAt time.Time, syncError error) error {
	state := RetailerSyncStateModel{Retailer: "phongvu", LastRunAt: &runAt, LastError: syncError.Error()}
	return repository.database.WithContext(ctx).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "retailer"}}, DoUpdates: clause.AssignmentColumns([]string{"last_run_at", "last_error"})}).Create(&state).Error
}

func compatibilityReason(accessoryCategory, laptopCategory string) string {
	return fmt.Sprintf("Phong Vũ %s phù hợp với nhu cầu %s", accessoryCategory, laptopCategory)
}

func (repository *Repository) LoadRecommendationSnapshot(
	ctx context.Context,
	productIDs []uuid.UUID,
) (RecommendationSnapshot, error) {
	var productModels []platformmodel.Product
	if err := repository.database.WithContext(ctx).Where("id IN ?", productIDs).Find(&productModels).Error; err != nil {
		return RecommendationSnapshot{}, fmt.Errorf("load recommendation products: %w", err)
	}
	var profiles []LaptopCompatibilityProfileModel
	if err := repository.database.WithContext(ctx).Where("product_id IN ? AND verified = ?", productIDs, true).Find(&profiles).Error; err != nil {
		return RecommendationSnapshot{}, fmt.Errorf("load compatibility profiles: %w", err)
	}
	verified := make(map[uuid.UUID]bool, len(profiles))
	for _, profile := range profiles {
		verified[profile.ProductID] = true
	}
	products := make([]LaptopTarget, 0, len(productModels))
	for _, product := range productModels {
		products = append(products, LaptopTarget{ID: product.ID, Name: product.Name, Category: product.Category, UpgradeProfileVerified: verified[product.ID]})
	}

	var accessoryModels []AccessoryModel
	if err := repository.database.WithContext(ctx).Find(&accessoryModels).Error; err != nil {
		return RecommendationSnapshot{}, fmt.Errorf("load accessories: %w", err)
	}
	var offerModels []RetailerOfferModel
	if err := repository.database.WithContext(ctx).Where("retailer = ?", "phongvu").Order("fetched_at DESC").Find(&offerModels).Error; err != nil {
		return RecommendationSnapshot{}, fmt.Errorf("load retailer offers: %w", err)
	}
	var compatibilityModels []AccessoryCompatibilityModel
	if err := repository.database.WithContext(ctx).
		Where("product_id IN ? OR product_id IS NULL", productIDs).
		Find(&compatibilityModels).Error; err != nil {
		return RecommendationSnapshot{}, fmt.Errorf("load accessory compatibility: %w", err)
	}

	latestOffers := make(map[uuid.UUID]RetailerOfferModel, len(offerModels))
	for _, offer := range offerModels {
		if _, exists := latestOffers[offer.AccessoryID]; !exists {
			latestOffers[offer.AccessoryID] = offer
		}
	}
	compatibilities := make(map[uuid.UUID][]AccessoryCompatibility, len(compatibilityModels))
	for _, model := range compatibilityModels {
		compatibilities[model.AccessoryID] = append(compatibilities[model.AccessoryID], AccessoryCompatibility{
			ID: model.ID, AccessoryID: model.AccessoryID, LaptopCategory: model.LaptopCategory,
			ProductID: model.ProductID, Reason: model.Reason, VerificationStatus: VerificationStatus(model.VerificationStatus),
		})
	}
	candidates := make([]Candidate, 0, len(accessoryModels))
	for _, model := range accessoryModels {
		offerModel, exists := latestOffers[model.ID]
		if !exists {
			continue
		}
		specifications := map[string]string{}
		if len(model.Specifications) > 0 {
			_ = json.Unmarshal(model.Specifications, &specifications)
		}
		candidates = append(candidates, Candidate{
			Accessory:       Accessory{ID: model.ID, Kind: Kind(model.Kind), Category: model.Category, Name: model.Name, Brand: model.Brand, ImageURL: model.ImageURL, Specifications: specifications},
			Offer:           RetailerOffer{ID: offerModel.ID, AccessoryID: offerModel.AccessoryID, RetailerProductID: offerModel.RetailerProductID, SourceURL: offerModel.SourceURL, Price: offerModel.Price, OriginalPrice: offerModel.OriginalPrice, InStock: offerModel.InStock, FetchedAt: offerModel.FetchedAt},
			Compatibilities: compatibilities[model.ID],
		})
	}
	return RecommendationSnapshot{Products: products, Candidates: candidates}, nil
}
