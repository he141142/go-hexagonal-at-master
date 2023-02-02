package form

import (
	"context"
	"errors"
	"hex-base/internal/core/domain"
)

func (store *formStorage) GetById(ctx context.Context, u uint) (*domain.Form, error) {
	return nil,errors.New("Not implemented yet")
}

func (store *formStorage) ListByTodoId(ctx context.Context, u uint) (*domain.FormList, error) {
	return nil,errors.New("Not implemented yet")
}
