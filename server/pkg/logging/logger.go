package logging

import "context"

type Fields map[string]interface{}

type Level string

const (
	LevelInfo  Level = "info"
	LevelDebug Level = "debug"
)

// nolint: interfacebloat
type Logger interface {
	GetLevel() Level
	SetLevel(level Level) Logger
	SetPretty(val bool) Logger

	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger

	Debug(args ...interface{})
	Debugf(format string, v ...interface{})
	Warning(args ...interface{})
	Warningf(format string, v ...interface{})
	Error(args ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
	Infof(format string, v ...interface{})
}
