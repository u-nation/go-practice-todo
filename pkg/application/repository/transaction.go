package repository

import (
	"context"
)

type TransactionRepository interface {
	Transaction(ctx context.Context, f func(ctx context.Context) error) error
}
