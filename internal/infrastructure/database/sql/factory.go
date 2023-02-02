package sql

import (
	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	dbAdapter "hex-base/internal/core/adapters/repo/sql_type/sql"
)

func NewFactorySql(lib constant.DBFrameWork,appCtx appctx.AppContext ) (dbAdapter.SqlAdapter, error) {
	_sqlActor := NewSqlActor(appCtx)
	switch lib {
	case constant.GORM:
		return _sqlActor.SetGormFramework(), nil
	case constant.ENT:
		return _sqlActor.SetEntFramework(), nil
	}
	return nil,nil
}
