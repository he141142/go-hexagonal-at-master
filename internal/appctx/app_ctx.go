package appctx

import (
	config "hex-base/internal/common"
	"hex-base/internal/core/adapters/logger"
)

type AppContext interface {
	Logger() logger.ILogger
	Viper() config.IViper
}

type Port int

type appContext struct {
	logger logger.ILogger
	viper  config.IViper
}

func (appCtx *appContext) Logger() logger.ILogger {
	return appCtx.logger
}

func (appCtx *appContext) Viper() config.IViper {
	return appCtx.viper
}

func NewAppContext(log logger.ILogger) AppContext {
	return &appContext{
		logger: log,
		viper:  config.Viper(),
	}
}
