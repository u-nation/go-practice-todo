package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/u-nation/go-practice-todo/pkg/infrastructure/record"
	"golang.org/x/xerrors"
)

type TodoRepositoryImpl struct {
	TransactionRepositoryImpl
}

func (r *TodoRepositoryImpl) FindAll(ctx context.Context) ([]*record.TodoRecord, error) {
	var todos []*record.TodoRecord
	if err := r.GetTx(ctx).Find(&todos).Error; err != nil {
		return nil, xerrors.Errorf("failed to FindAll :%w", err)
	}
	return todos, nil
}

func (r *TodoRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (*record.TodoRecord, error) {
	var todo record.TodoRecord
	var pk = id.String()

	if err := r.GetTx(ctx).Where("id = ?", pk).First(&todo).Error; err != nil {
		return nil, xerrors.Errorf("failed to Get id=%s :%w", pk, err)
	}
	return &todo, nil
}

func (r *TodoRepositoryImpl) Save(ctx context.Context, todo *record.TodoRecord) (*record.TodoRecord, error) {
	if err := r.GetTx(ctx).Create(todo).Error; err != nil {
		return nil, xerrors.Errorf("failed to Save :%w", err)
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) Update(ctx context.Context, todo *record.TodoRecord) (*record.TodoRecord, error) {
	if err := r.GetTx(ctx).Save(todo).Error; err != nil {
		return nil, xerrors.Errorf("failed to Update :%w", err)
	}
	return todo, nil
}

func (r *TodoRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	var pk = id.String()
	if err := r.GetTx(ctx).Delete(&record.TodoRecord{}, pk).Error; err != nil {
		return xerrors.Errorf("failed to Delete :%w", err)
	}
	return nil
}
