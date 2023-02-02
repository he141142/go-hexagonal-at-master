package infrastructure

import (
	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/repo/sql_type/sql"
	"hex-base/internal/core/adapters/validator"
	sql2 "hex-base/internal/infrastructure/database/sql"
	logger_provider "hex-base/internal/infrastructure/logger"
	"hex-base/internal/infrastructure/router"
	"time"
)

type config struct {
	appCtx appctx.AppContext
	appName       string
	logger        logger.ILogger
	validator     validator.ValidatorAdapter
	dbSQL         sql.SqlAdapter
	dbNoSQL       any
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}


func NewConfig() *config {
	return &config{}
}


func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}


func (c *config) Name(name string) *config {
	c.appName = name
	return c
}


func (c *config) Logger(instance logger_provider.TypeLogger) *config {
	log, err := logger_provider.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}
	c.logger = log
	c.logger.Infof("Successfully configured log")
	return c
}

func (c *config) DbSQL(instance int) *config {
	if c.appCtx == nil{
		panic(any("app context is missing"))
		return nil
	}
	db, err := sql2.NewFactorySql(constant.ENT,c.appCtx)
	if err != nil {
		c.logger.Fatalln(err, "Could not make a connection to the database")
	}

	c.logger.Infof("Successfully connected to the SQL database")

	c.dbSQL = db
	return c
}