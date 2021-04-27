package repository

import "context"

type HealthCheckRepository interface {
	Ping(ctx context.Context) error
}
