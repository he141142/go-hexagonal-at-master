package ent

import (
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	ent_form "hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/form"
	ent_todo "hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/todo"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/todo"
)

type EntAdapter interface {
	Engine() *ent.Client
	FormRepository() ent_form.FormStorage
	TodoRepository() ent_todo.TodoStorage
}

type entAdapter struct {
	client      *ent.Client
	formStorage form.FormStorage
	todoStorage todo.TodoStorage
}

func (adapter *entAdapter) Engine() *ent.Client {
	return adapter.client
}

func (adapter *entAdapter) FormRepository() ent_form.FormStorage {
	return adapter.formStorage
}

func (adapter *entAdapter) TodoRepository() ent_todo.TodoStorage {
	return adapter.todoStorage
}

func NewEntAdapter(client *ent.Client) EntAdapter {
	return &entAdapter{
		client:      nil,
		formStorage: ent_form.NewFormStorage(client),
		todoStorage: ent_todo.NewTodoStorage(client),
	}
}
