package model

import (
	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"gorm.io/datatypes"
)

type KnowledgeBase struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Type      string          `gorm:"type:varchar(50);not null"`
	Content   string          `gorm:"type:text;not null"`
	Metadata  datatypes.JSON  `gorm:"type:jsonb"`
	Embedding pgvector.Vector `gorm:"type:vector(1536)"`
}

type ProductInsight struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ProductID uuid.UUID       `gorm:"type:uuid;index;not null"`
	Insight   string          `gorm:"type:text;not null"`
	Embedding pgvector.Vector `gorm:"type:vector(1536)"`
}
