package logger

type ILogger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalln(args ...interface{})
	WithError(err error) ILogger
	WithFields(fields Fields) ILogger
}

type Fields map[string]interface{}

