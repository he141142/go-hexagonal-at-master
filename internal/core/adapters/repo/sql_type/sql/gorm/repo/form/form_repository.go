package form

import (
	"context"
	"gorm.io/gorm"
	"hex-base/internal/core/domain"
)

type FormStorage interface {
	Create(ctx context.Context, form domain.Form) (*domain.Form, error)
}

type formStorage struct {
	db *gorm.DB
}

func NewFormStorage(db *gorm.DB) FormStorage {
	return &formStorage{
		db: db,
	}
}
