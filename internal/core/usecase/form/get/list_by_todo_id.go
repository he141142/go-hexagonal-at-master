package get

import (
	"context"
	"hex-base/internal/core/domain"
	"time"
)

type (
	ListFormByTodoIDUseCase interface {
		Execute(context.Context, uint) (*ListFormByTodoIdOutput, error)
	}

	ListFormByTodoIdOutput struct {
		Data []*FormObject
	}

	FormObject struct {
		Id   uint
		Todo *TodoObject
	}

	TodoObject struct {
		Id   uint
		Name string
		Task string
	}

	ListFormByTodoIDInteractor struct {
		repo       domain.FormRepository
		presenter  ListFormByTodoIDPresenter
		ctxTimeout time.Duration
	}

	ListFormByTodoIDPresenter interface {
		Output(form *domain.FormList) *ListFormByTodoIdOutput
	}
)


func NewListFormByTodoIDInteractor(
	repo domain.FormRepository,
	presenter ListFormByTodoIDPresenter,
	t time.Duration,
) ListFormByTodoIDUseCase {
	return ListFormByTodoIDInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}


func (a ListFormByTodoIDInteractor) Execute(ctx context.Context,todoId uint) (*ListFormByTodoIdOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	form, err := a.repo.ListByTodoId(ctx, todoId)

	if err != nil {
		return a.presenter.Output(form), err
	}

	return a.presenter.Output(form), nil
}