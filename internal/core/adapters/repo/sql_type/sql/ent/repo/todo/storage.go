package todo

import (
	"context"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/domain"
)

type TodoStorage interface {
	Create(context.Context, *domain.Todo) (*domain.Todo, error)
	GetById(context.Context, uint) (*domain.Todo, error)
}

type todoStorage struct {
	client *ent.Client
}

func NewTodoStorage(client *ent.Client) TodoStorage{
	return &todoStorage{
		client: client,
	}
}