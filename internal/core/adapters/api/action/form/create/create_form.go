package create

import (
	"encoding/json"
	"hex-base/internal/core/adapters/api/logging"
	"hex-base/internal/core/adapters/api/response"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/validator"
	"hex-base/internal/core/usecase/form/create"
	"net/http"
)

type CreateFormAction struct {
	uc create.CreateFormUseCase
	log logger.ILogger
	validator validator.ValidatorAdapter
}

func NewCreateFormAction(uc create.CreateFormUseCase, log logger.ILogger, v validator.ValidatorAdapter) CreateFormAction {
	return CreateFormAction{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

func (a CreateFormAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_form"

	var input create.CreateFormInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if errs := a.validateInput(input); len(errs) > 0 {
		logging.NewError(
			a.log,
			response.ErrInvalidInput,
			logKey,
			http.StatusBadRequest,
		).Log("invalid input")

		response.NewErrorMessage(errs, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input)

	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when creating a new account")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success creating account")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}


func (a CreateFormAction) validateInput(input create.CreateFormInput) []string {
	var msgs []string

	err := a.validator.Validate(input)
	if err != nil {
		for _, msg := range a.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}
