

package logger
import (
"go.uber.org/zap"
	"hex-base/internal/core/adapters/logger"
)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (logger.ILogger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sugar := log.Sugar()
	defer log.Sync()

	return &zapLogger{logger: sugar}, nil
}

func (l *zapLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *zapLogger) Fatalln(args ...interface{}) {
	l.logger.Fatal(args)
}

func (l *zapLogger) WithFields(fields logger.Fields) logger.ILogger {
	var f = make([]interface{}, 0)
	for index, field := range fields {
		f = append(f, index)
		f = append(f, field)
	}

	log := l.logger.With(f...)
	return &zapLogger{logger: log}
}

func (l *zapLogger) WithError(err error) logger.ILogger {
	var log = l.logger.With(err.Error())
	return &zapLogger{logger: log}
}