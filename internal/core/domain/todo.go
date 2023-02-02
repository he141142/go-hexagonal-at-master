package domain

import (
	"context"
	"hex-base/internal/constant"
)

type (
	Todo struct {
		id   uint
		form []*Form
		name string
		task string
	}

	TodoRepository interface {
		Create(context.Context, *Todo) (*Todo, error)
		GetById(context.Context, uint) (*Todo, error)
	}
)

func (Todo) TableName() string {
	return constant.Todo.String()
}

func (t *Todo) ID() uint {
	return t.id
}

func (t *Todo) Form() []*Form {
	return t.form
}

func (t *Todo) Name() string {
	return t.name
}

func (t *Todo) Task() string {
	return t.task
}

func NewTodo(id uint, name string, task string) *Todo {
	return &Todo{
		id: id,
		name: name,
		task: task,
	}
}
