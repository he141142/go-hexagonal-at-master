package create

import (
	"context"
	"hex-base/internal/core/domain"
	"time"
)

type (
	// CreateAccountUseCase input port
	CreateFormUseCase interface {
		Execute(context.Context, CreateFormInput) (CreateFormOutput, error)
	}

	// CreateAccountInput input data
	CreateFormInput struct {
		Category string `json:"category" validate:"required"`
		Tile     string `json:"tile" validate:"required"`
	}

	// CreateAccountOutput output data
	CreateFormOutput struct {
		ID        uint    `json:"id"`
		Category  string  `json:"category"`
		Tile      string  `json:"tile"`
		TodoID    float64 `json:"todo_id"`
		Status    string  `json:"status"`
		IsDeleted bool    `json:"is_deleted"`
	}

	CreateFormInteractor struct {
		repo       domain.FormRepository
		presenter  CreateFormPresenter
		ctxTimeout time.Duration
	}

	// CreateAccountPresenter output port
	CreateFormPresenter interface {
		Output(form domain.Form) CreateFormOutput
	}
)

// NewCreateAccountInteractor creates new createAccountInteractor with its dependencies
func NewCreateFormInteractor(
	repo domain.FormRepository,
	presenter CreateFormPresenter,
	t time.Duration,
) CreateFormUseCase {
	return CreateFormInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a CreateFormInteractor) Execute(ctx context.Context, input CreateFormInput) (CreateFormOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	var _form = domain.Form{
		Category: input.Category,
		Title:    input.Tile,
	}

	form, err := a.repo.Create(ctx, _form)

	if err != nil {
		return a.presenter.Output(domain.Form{}), err
	}

	return a.presenter.Output(*form), nil
}
