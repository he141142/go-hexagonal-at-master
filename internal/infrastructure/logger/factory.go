package logger

import (
	"errors"
	logger2 "hex-base/internal/core/adapters/logger"
)

type TypeLogger int32

var loggerTypeName = []string{
	"zap",
	"logrus",
}

const (
	ZapType TypeLogger = iota + 1
	LogrusType
)

var (
	errInvalidLoggerInstance = errors.New("invalid log instance")
)

func NewLoggerFactory(instanceType TypeLogger) (logger2.ILogger, error) {
	switch instanceType {
	case ZapType:
		if loggerTyp, err := NewZapLogger(); err != nil {
			return nil, err
		} else {
			return loggerTyp, nil
		}
	case LogrusType:
		return NewLogrusLogger(), nil
	default:
		return nil, errInvalidLoggerInstance
	}
}
