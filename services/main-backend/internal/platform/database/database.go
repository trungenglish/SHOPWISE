package database

import (
	"fmt"
	"time"

	"shopwise/retail/internal/platform/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(databaseURL string, ginMode string) (*gorm.DB, error) {
	logLevel := logger.Warn
	if ginMode == "debug" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get sql db: %w", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Ensure vector extension exists for pgvector
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS vector;").Error; err != nil {
		return nil, fmt.Errorf("create vector extension: %w", err)
	}

	if err := db.AutoMigrate(
		&model.AppMetadata{},
		&model.User{},
		&model.UserPreference{},
		&model.Product{},
		&model.Inventory{},
		&model.Promotion{},
		&model.Order{},
		&model.KnowledgeBase{},
		&model.ProductInsight{},
	); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}

	// Create HNSW indexes for pgvector (GORM doesn't support hnsw index directly via tags easily in all cases)
	// We use raw SQL to ensure they are created correctly
	db.Exec("CREATE INDEX IF NOT EXISTS idx_knowledge_base_embedding ON knowledge_bases USING hnsw (embedding vector_l2_ops);")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_product_insight_embedding ON product_insights USING hnsw (embedding vector_l2_ops);")

	return db, nil
}

func Ping(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
