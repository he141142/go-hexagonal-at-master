package sql

import (
	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/repo/sql_type/sql"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm"
	"hex-base/internal/infrastructure/database"
	ent_provider "hex-base/internal/infrastructure/database/sql/ent"
	gorm_provider "hex-base/internal/infrastructure/database/sql/gorm"
)

type sqlActor struct {
	appCtx appctx.AppContext
	gorm   gorm.GormAdapter
	ent    ent.EntAdapter
	framework constant.DBFrameWork
}

func (actor *sqlActor) Gorm() gorm.GormAdapter {
	return actor.gorm
}

func (actor *sqlActor) Ent() ent.EntAdapter {
	return actor.ent
}

func (actor *sqlActor) SetGormFramework() sql.SqlAdapter {
	actor.framework = constant.GORM
	actor.gorm = gorm_provider.NewGormProvider(actor.appCtx, database.NewSqlConfig())
	return actor
}

func (actor *sqlActor) SetEntFramework() sql.SqlAdapter {
	actor.framework = constant.ENT
	actor.ent = ent_provider.NewEntProvider(actor.appCtx, database.NewSqlConfig())
	return actor
}

func (actor *sqlActor) GetFramework() constant.DBFrameWork {
	return actor.framework
}

func NewSqlActor(appCtx appctx.AppContext) *sqlActor {
	return &sqlActor{
		appCtx: appCtx,
	}
}
