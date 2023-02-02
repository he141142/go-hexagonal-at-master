package create

import (
	"hex-base/internal/core/domain"
	"hex-base/internal/core/usecase/form/create"
)

type createFormPresenter struct {

}

func (a createFormPresenter) Output(form domain.Form) create.CreateFormOutput {
	return create.CreateFormOutput{
		ID:        form.Id,
		Category:  form.Category,
		Tile:      form.Title,
		TodoID: float64(form.TodoID),
		Status:    form.Status,
		IsDeleted: form.IsDeleted,
	}
}

func NewCreateFormPresenter() create.CreateFormPresenter{
	return createFormPresenter{}
}