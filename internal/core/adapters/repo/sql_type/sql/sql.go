package sql

import (
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm"
)

type SqlAdapter interface {
	GetFramework() constant.DBFrameWork

	SetGormFramework() SqlAdapter
	SetEntFramework() SqlAdapter

	Gorm() gorm.GormAdapter
	Ent() ent.EntAdapter
}
