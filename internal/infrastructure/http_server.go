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
	validator_provider "hex-base/internal/infrastructure/validator"
	"strconv"
	"time"
)

type config struct {
	appCtx        appctx.AppContext
	appName       string
	logger        logger.ILogger
	validator     validator.ValidatorAdapter
	dbSQL         sql.SqlAdapter
	dbNoSQL       any
	ctxTimeout    time.Duration
	webServerPort appctx.Port
	webServer     router.Server
}

func NewConfig(appCtx appctx.AppContext) *config {
	return &config{
		appCtx: appCtx,
	}
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

func (c *config) DbSQL(framework constant.DBFrameWork) *config {
	if c.appCtx == nil {
		panic(any("app context is missing"))
		return nil
	}
	db, err := sql2.NewFactorySql(framework, c.appCtx)

	if err != nil {
		c.logger.Fatalln(err, "Could not make a connection to the database")
	}
	c.logger.Infof("Successfully connected to the SQL database")
	c.dbSQL = db
	return c
}

func (c *config) DbNoSql() *config{
	c.appCtx.Logger().Fatalln("No Sql not implemented yet")
	return c
}

func (c *config) Validator(instance validator_provider.ValidatorType) *config {
	v, err := validator_provider.NewValidatorFactory(instance)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured validator")
	c.validator = v
	return c
}

func (c *config) WebServer(instance router.RouterFramework) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.logger,
		c.dbSQL,
		c.dbNoSQL,
		c.validator,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPort = appctx.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
