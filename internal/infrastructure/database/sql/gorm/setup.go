package gorm

import (
	"gorm.io/gorm"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/todo"
)

func (provider *gormProvider) Engine() *gorm.DB {
	return provider.db
}

func (provider *gormProvider) FormStorage() form.FormStorage {
	return provider.formStore
}

func (provider *gormProvider) TodoStorage() todo.TodoStorage {
	return provider.todoStore
}
