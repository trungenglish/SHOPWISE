package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"

	"shopwise/retail/internal/platform/config"
	"shopwise/retail/internal/platform/database"
	"shopwise/retail/internal/platform/database/model"
	"shopwise/retail/internal/platform/logger"
)

type SeedProduct struct {
	ID             string                 `json:"id"`
	SKU            string                 `json:"sku"`
	Name           string                 `json:"name"`
	Brand          string                 `json:"brand"`
	Category       string                 `json:"category"`
	Price          float64                `json:"price"`
	Specifications map[string]interface{} `json:"specifications"`
	Metadata       map[string]interface{} `json:"metadata"`
}

const demoUSDToVNDRate = 25_000

func priceInVND(priceUSD float64) int64 {
	return int64(priceUSD * demoUSDToVNDRate)
}

func seedConflictClause() clause.OnConflict {
	return clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"sku", "name", "brand", "category", "price", "specifications", "metadata",
		}),
	}
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	log := logger.New(cfg.AppName, cfg.LogLevel)

	db, err := database.Connect(cfg.DatabaseURL, cfg.GinMode)
	if err != nil {
		log.Error("failed to connect database", slog.Any("error", err))
		os.Exit(1)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Error("database close failed", slog.Any("error", err))
		}
	}()

	seedFilePath := filepath.Join("cmd", "seed", "products", "seed_products_demo.json")
	data, err := os.ReadFile(seedFilePath)
	if err != nil {
		log.Error("failed to read seed file", slog.Any("error", err), slog.String("path", seedFilePath))
		os.Exit(1)
	}

	var seedProducts []SeedProduct
	if err := json.Unmarshal(data, &seedProducts); err != nil {
		log.Error("failed to unmarshal seed data", slog.Any("error", err))
		os.Exit(1)
	}

	var products []model.Product
	for _, sp := range seedProducts {
		specBytes, _ := json.Marshal(sp.Specifications)
		metaBytes, _ := json.Marshal(sp.Metadata)

		id, err := uuid.Parse(sp.ID)
		if err != nil {
			log.Error("invalid uuid", slog.String("id", sp.ID), slog.Any("error", err))
			continue
		}

		products = append(products, model.Product{
			ID:             id,
			SKU:            sp.SKU,
			Name:           sp.Name,
			Brand:          sp.Brand,
			Category:       sp.Category,
			Price:          priceInVND(sp.Price),
			Specifications: datatypes.JSON(specBytes),
			Metadata:       datatypes.JSON(metaBytes),
		})
	}

	if len(products) > 0 {
		result := db.Clauses(seedConflictClause()).Create(&products)
		if result.Error != nil {
			log.Error("failed to seed products", slog.Any("error", result.Error))
			os.Exit(1)
		}
		log.Info("successfully seeded products", slog.Int("inserted", int(result.RowsAffected)), slog.Int("total", len(products)))
	} else {
		log.Info("no products to seed")
	}
}
