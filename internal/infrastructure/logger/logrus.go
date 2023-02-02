package logger

import (
	"github.com/sirupsen/logrus"
	logger2 "hex-base/internal/core/adapters/logger"
)

type logrusLogger struct {
	logger *logrus.Logger
}

func (l *logrusLogger) WithError(err error) logger2.ILogger {
	return &logrusLogEntry{
		entry: l.logger.WithError(err),
	}
}

func (l *logrusLogger) WithFields(fields logger2.Fields) logger2.ILogger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func NewLogrusLogger() logger2.ILogger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &logrusLogger{logger: log}
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}


type logrusLogEntry struct {
	entry *logrus.Entry
}

func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogEntry) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

func (l *logrusLogEntry) WithFields(fields logger2.Fields) logger2.ILogger {
	return &logrusLogEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogEntry) WithError(err error) logger2.ILogger {
	return &logrusLogEntry{
		entry: l.entry.WithError(err),
	}
}

func convertToLogrusFields(fields logger2.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, field := range fields {
		logrusFields[index] = field
	}

	return logrusFields
}
