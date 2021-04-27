package repository

import (
	"context"
	"golang.org/x/xerrors"
)

type HealthCheckRepositoryImpl struct {
	TransactionRepositoryImpl
}

func (r HealthCheckRepositoryImpl) Ping(ctx context.Context) error {
	sqlDB, err := r.GetTx(ctx).DB()
	if err != nil {
		return xerrors.Errorf("failed to Healthcheck: %w", err)
	}

	return sqlDB.Ping()
}
