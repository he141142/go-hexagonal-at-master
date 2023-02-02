package gorm

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/todo"
)

type GormAdapter interface {
	Engine()  *gorm.DB
	FormStorage() form.FormStorage
	TodoStorage() todo.TodoStorage
}

type gormAdapter struct {
	db          *gorm.DB
	formStorage form.FormStorage
	todoStorage todo.TodoStorage
}
func (adapter *gormAdapter) Engine()  *gorm.DB{
	return adapter.db
}
func (adapter *gormAdapter) FormStorage() form.FormStorage  {
	return adapter.formStorage
}

func (adapter *gormAdapter) TodoStorage() todo.TodoStorage {
	return adapter.todoStorage
}

func NewGormAdapter(db *gorm.DB) GormAdapter {
	return &gormAdapter{
		db:          db,
		formStorage: form.NewFormStorage(db),
	}
}
