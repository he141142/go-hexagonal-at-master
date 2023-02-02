package sql

import (
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm"
)

type SqlAdapter interface {
	SetGormFramework() SqlAdapter
	SetEntFramework() SqlAdapter

	Gorm() gorm.GormAdapter
	Ent() ent.EntAdapter
}
