package ent

import (
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/todo"
)


func (provider *entProvider) FormRepository() form.FormStorage  {
		return provider.formStore
}

func (provider *entProvider) TodoRepository() todo.TodoStorage {
	return provider.todoStore
}