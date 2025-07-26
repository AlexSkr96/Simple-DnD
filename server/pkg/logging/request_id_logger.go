package logging

import (
	"context"
)

const loggerKeyRequestID = "requestID"

type ctxKeyRequestID struct{}

func GetRequestIDFromCtx(ctx context.Context) string {
	val := ctx.Value(ctxKeyRequestID{})
	if val == nil {
		return ""
	}

	return ctx.Value(ctxKeyRequestID{}).(string) //nolint: forcetypeassert
}

func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ctxKeyRequestID{}, requestID)
}

func NewRequestIDLogger(logger Logger) *RequestIDLogger {
	return &RequestIDLogger{
		logger: logger,
	}
}

type RequestIDLogger struct {
	logger Logger
}

func (l *RequestIDLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l RequestIDLogger) WithField(key string, value interface{}) Logger {
	return NewRequestIDLogger(l.logger.WithField(key, value))
}

func (l RequestIDLogger) WithError(err error) Logger {
	return NewRequestIDLogger(l.logger.WithError(err))
}

func (l RequestIDLogger) WithContext(ctx context.Context) Logger {
	requestID := GetRequestIDFromCtx(ctx)
	if len(requestID) == 0 {
		return NewRequestIDLogger(l.logger.WithContext(ctx))
	}

	logger := l.logger.WithField(loggerKeyRequestID, requestID)

	return NewRequestIDLogger(logger.WithContext(ctx))
}

func (l RequestIDLogger) GetLevel() Level {
	return l.logger.GetLevel()
}

func (l RequestIDLogger) SetLevel(level Level) Logger {
	return NewRequestIDLogger(l.logger.SetLevel(level))
}

func (l RequestIDLogger) SetPretty(val bool) Logger {
	return NewRequestIDLogger(l.logger.SetPretty(val))
}

func (l RequestIDLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l RequestIDLogger) Warningf(format string, v ...interface{}) {
	l.logger.Warningf(format, v...)
}

func (l RequestIDLogger) Warning(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l RequestIDLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l RequestIDLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l RequestIDLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l RequestIDLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l RequestIDLogger) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}
