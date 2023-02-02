package sql

import (
	gorm2 "gorm.io/gorm"
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent"
	ent2 "hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm"
)

type BuilderInterface interface {
	SetGormFramework(gormDb *gorm2.DB) SqlAdapter
	SetEntFramework(entClient *ent2.Client) SqlAdapter
}

type SqlAdapter interface {
	Gorm() gorm.GormAdapter
	Ent() ent.EntAdapter
}

type sqlAdapter struct {
	gorm      gorm.GormAdapter
	ent       ent.EntAdapter
	dialect   constant.Driver
}

func (adapter *sqlAdapter) Gorm() gorm.GormAdapter {
	return adapter.gorm
}

func (adapter *sqlAdapter) Ent() ent.EntAdapter {
	return adapter.ent
}

func (adapter *sqlAdapter) Build() SqlAdapter{
	return adapter
}

