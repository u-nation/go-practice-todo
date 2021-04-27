package repository

import (
	"context"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

const (
	TransactionKey = "transaction_key"
)

// GetTx Contextからtxを取得する
func (t TransactionRepositoryImpl) GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TransactionKey).(*gorm.DB); ok {
		return tx
	}

	return t.DB
}

func (t TransactionRepositoryImpl) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, TransactionKey, tx)
		if err := f(ctx); err != nil {
			return xerrors.Errorf("failed to Transaction: %w", err)
		}
		// return nil will commit the whole transaction
		return nil
	})
}
