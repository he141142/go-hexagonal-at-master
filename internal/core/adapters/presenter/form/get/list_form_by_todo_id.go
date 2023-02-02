package get

import (
	"hex-base/internal/core/domain"
	"hex-base/internal/core/usecase/form/get"
)

type listFormByTodoIdPresenter struct {
}

func (presenter *listFormByTodoIdPresenter) Output(form *domain.FormList) *get.ListFormByTodoIdOutput {
	out := &get.ListFormByTodoIdOutput{
		Data: make([]*get.FormObject, 0),
	}
	for k, v := range form.Data {
		formData := out.Data[k]
		formData.Id = v.Id
		formData.Todo = &get.TodoObject{
			Id:   v.Todo.ID(),
			Name: v.Todo.Name(),
			Task: v.Todo.Task(),
		}
		out.Data = append(out.Data, formData)
	}

	return out
}

func NewListFormByTodoIdPresenter() *listFormByTodoIdPresenter {
	return &listFormByTodoIdPresenter{}
}
