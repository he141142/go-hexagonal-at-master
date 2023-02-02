package form

import (
	"context"
	"gorm.io/gorm"
	"hex-base/internal/core/domain"
)

type FormStorage interface {
	Create(ctx context.Context, form domain.Form) (*domain.Form, error)
	GetById(context.Context, uint) (*domain.Form, error)
	ListByTodoId(context.Context, uint) (*domain.FormList,error)
}

type formStorage struct {
	db *gorm.DB
}


func NewFormStorage(db *gorm.DB) FormStorage {
	return &formStorage{
		db: db,
	}
}
