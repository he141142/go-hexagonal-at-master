package sql

import (
	gorm2 "gorm.io/gorm"
	"hex-base/internal/constant"
	ent3 "hex-base/internal/core/adapters/repo/sql_type/sql/ent"
	ent2 "hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm"
)

func (adapter *sqlAdapter) SetGormFramework(gormDb *gorm2.DB)SqlAdapter {
	adapter.gorm = gorm.NewGormAdapter(gormDb)
	return adapter
}

func (adapter *sqlAdapter) SetEntFramework(entClient *ent2.Client) SqlAdapter {
	adapter.ent = ent3.NewEntAdapter(entClient)
	return adapter
}



func SqlAdapterBuilder(dialect  constant.Driver) BuilderInterface {
	return &sqlAdapter{
		gorm:    nil,
		ent:     nil,
		dialect: dialect,
	}
}