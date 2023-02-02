package logging

import "hex-base/internal/core/adapters/logger"

type Error struct {
	log        logger.ILogger
	err        error
	key        string
	httpStatus int
}

func NewError(log logger.ILogger, err error, key string, httpStatus int) Error {
	return Error{
		log:        log,
		err:        err,
		key:        key,
		httpStatus: httpStatus,
	}
}

func (e Error) Log(msg string) {
	e.log.WithFields(logger.Fields{
		"key":         e.key,
		"error":       e.err.Error(),
		"http_status": e.httpStatus,
	}).Errorf(msg)
}