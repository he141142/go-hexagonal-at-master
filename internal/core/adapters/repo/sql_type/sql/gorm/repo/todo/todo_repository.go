package todo

import (
	"context"
	"gorm.io/gorm"
	"hex-base/internal/core/domain"
)

type TodoStorage interface {
	Create(context.Context, *domain.Todo) (*domain.Todo, error)
	GetById(context.Context, uint) (*domain.Todo, error)
}

type todoStorage struct {
	db *gorm.DB
}

func NewTodoStorage(db *gorm.DB) TodoStorage {
	return &todoStorage{
		db: db,
	}
}
