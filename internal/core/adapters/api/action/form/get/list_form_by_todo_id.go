package get

import (
	"errors"
	"hex-base/internal/core/adapters/api/logging"
	"hex-base/internal/core/adapters/api/response"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/usecase/form/get"
	"net/http"
	"strconv"
)

type ListFormByTodoIdAction struct {
	uc get.ListFormByTodoIDUseCase
	log logger.ILogger
}

func NewListFormByTodoIdAction(	uc get.ListFormByTodoIDUseCase, log logger.ILogger,) ListFormByTodoIdAction {
	return ListFormByTodoIdAction{
		uc:        uc,
		log:       log,
	}
}

func (a ListFormByTodoIdAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "list_form_by_todo_id"

	todoIdString, ok := r.URL.Query()["todo_id"]
	if !ok || len(todoIdString[0]) < 1 {
		erMsg := "url Param 'todo_id' is missing"
		err := errors.New(erMsg)
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log(erMsg)
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	todoid,err := strconv.Atoi(todoIdString[0])
	if err != nil{
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when convert query param to int")
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), uint(todoid))

	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when list form")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("operation success")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
