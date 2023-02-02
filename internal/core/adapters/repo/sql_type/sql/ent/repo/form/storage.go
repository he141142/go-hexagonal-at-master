package form

import (
	"context"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/domain"
)

type FormStorage interface {
	Create(ctx context.Context, form domain.Form) (*domain.Form, error)
}

type formStorage struct {
	client *ent.Client
}

func NewFormStorage(client *ent.Client) FormStorage{
	return &formStorage{
		client: client,
	}
}


