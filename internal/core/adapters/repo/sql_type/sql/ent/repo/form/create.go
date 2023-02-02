package form

import (
	"context"
	"hex-base/internal/core/domain"
)

func (store *formStorage) Create(ctx context.Context, form domain.Form) (*domain.Form, error) {
	formEnt, err := store.client.Form.Create().Save(ctx)
	if err != nil {
		return nil, err
	}
	//transform
	return domain.NewForm(
		uint(formEnt.ID),
		formEnt.Category,
		formEnt.IsDeleted,
		formEnt.Status,
		uint(formEnt.TodoID),
		formEnt.Title,
	), nil
}