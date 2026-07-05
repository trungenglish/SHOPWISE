package local

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

func NewStorage(basePath string) (*Storage, error) {
	if err := os.MkdirAll(basePath, 0o750); err != nil {
		return nil, err
	}
	return &Storage{basePath: basePath}, nil
}

func (s *Storage) Put(_ context.Context, key string, body io.Reader, _ string) error {
	fullPath := filepath.Join(s.basePath, filepath.FromSlash(key))
	if err := os.MkdirAll(filepath.Dir(fullPath), 0o750); err != nil {
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	return err
}

func (s *Storage) Delete(_ context.Context, key string) error {
	fullPath := filepath.Join(s.basePath, filepath.FromSlash(key))
	if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}
