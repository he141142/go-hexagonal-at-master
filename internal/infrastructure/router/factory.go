package router

import (
	"errors"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/repo/sql_type/sql"
	"hex-base/internal/core/adapters/validator"
	"hex-base/internal/infrastructure/router/gin"
	"time"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid router server instance")
)

const (
	InstanceGorillaMux int = iota
	InstanceGin
)


func NewWebServerFactory(
	instance int,
	log logger.ILogger,
	dbSQL sql.SqlAdapter,
	dbNoSQL any,
	validator validator.ValidatorAdapter,
	port Port,
	ctxTimeout time.Duration,
) (Server, error) {
	switch instance {
	case InstanceGorillaMux:
		return nil,errors.New("not implement yet")
	case InstanceGin:
		return gin.NewGinServer(log, dbSQL, validator, port, ctxTimeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}