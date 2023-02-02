package todo

import (
	"context"
	"errors"
	"hex-base/internal/core/domain"
)

func (storage *todoStorage) Create(context.Context, *domain.Todo) (*domain.Todo, error){
	return nil,errors.New("not implemented")
}