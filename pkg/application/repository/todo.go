package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/u-nation/go-practice-todo/pkg/infrastructure/record"
)

type TodoRepository interface {
	FindAll(ctx context.Context) ([]*record.TodoRecord, error)
	Get(ctx context.Context, id uuid.UUID) (*record.TodoRecord, error)
	Save(ctx context.Context, todo *record.TodoRecord) (*record.TodoRecord, error)
	Update(ctx context.Context, todo *record.TodoRecord) (*record.TodoRecord, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
