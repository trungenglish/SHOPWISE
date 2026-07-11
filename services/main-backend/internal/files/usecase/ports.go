package usecase

import (
	"context"
	"io"
)

// StoragePort abstracts object storage for the files module.
type StoragePort interface {
	Put(ctx context.Context, key string, body io.Reader, contentType string) error
	URL(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}
