package sql

import (
	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	dbAdapter "hex-base/internal/core/adapters/repo/sql_type/sql"
)

func NewFactorySql(lib constant.DBFrameWork,appCtx appctx.AppContext ) (dbAdapter.SqlAdapter, error) {
	switch lib {
	case constant.GORM:
		return NewSqlActor(appCtx).SetGormFramework(), nil
	case constant.ENT:
		return NewSqlActor(appCtx).SetEntFramework(), nil
	}
	return nil,nil
}
