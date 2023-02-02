package logger

import (
	"errors"
	"github.com/gsabadini/go-bank-transfer/adapter/logger"
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

func NewLoggerFactory(instanceType TypeLogger) (logger.Logger, error) {
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
