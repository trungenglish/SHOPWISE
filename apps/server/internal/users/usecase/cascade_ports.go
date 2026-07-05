package usecase

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IdentityAccountDeleter interface {
	DeleteUserData(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error
}

type AdvisorAccountDeleter interface {
	DeleteUserData(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error
}

type WishlistAccountDeleter interface {
	DeleteUserData(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error
}

type OwnershipAccountDeleter interface {
	DeleteUserData(ctx context.Context, tx *gorm.DB, userID uuid.UUID) ([]string, error)
}

type NotificationAccountDeleter interface {
	DeleteUserData(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error
}

type UserAccountDeleter interface {
	DeleteAccount(ctx context.Context, tx *gorm.DB, userID uuid.UUID) error
}

type FileStorageDeleter interface {
	Delete(ctx context.Context, key string) error
}

type TransactionRunner interface {
	Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

type gormTransactionRunner struct {
	db *gorm.DB
}

func (r gormTransactionRunner) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(fn)
}

type AccountDeletionDeps struct {
	DB           *gorm.DB
	Transactions TransactionRunner
	Identity     IdentityAccountDeleter
	Advisor      AdvisorAccountDeleter
	Wishlist     WishlistAccountDeleter
	Ownership    OwnershipAccountDeleter
	Notification NotificationAccountDeleter
	Users        UserAccountDeleter
	Storage      FileStorageDeleter
}

func (deps AccountDeletionDeps) runTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	if deps.Transactions != nil {
		return deps.Transactions.Transaction(ctx, fn)
	}
	if deps.DB == nil {
		return fn(nil)
	}
	return gormTransactionRunner{db: deps.DB}.Transaction(ctx, fn)
}
