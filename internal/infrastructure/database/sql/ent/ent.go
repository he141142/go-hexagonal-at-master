package ent

import (
	"hex-base/internal/appctx"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/todo"
)

type entProvider struct {
	client *ent.Client
	formStore form.FormStorage
	todoStore todo.TodoStorage
}

func (provider *entProvider) Engine() *ent.Client {
	return provider.client
}

func NewEntProvider(appCtx appctx.AppContext) *entProvider {
	return &entProvider{}
}

// get the core
func ConnectDb() any {
	return any(1)
}
