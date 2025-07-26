package logging

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func NewStacktraceLogger(logger Logger) *StacktraceLogger {
	return &StacktraceLogger{
		logger: logger,
	}
}

type StacktraceLogger struct {
	logger Logger
}

func (r StacktraceLogger) GetLevel() Level {
	return r.logger.GetLevel()
}

func (r StacktraceLogger) SetLevel(level Level) Logger {
	return NewStacktraceLogger(r.logger.SetLevel(level))
}

func (r StacktraceLogger) SetPretty(val bool) Logger {
	return NewStacktraceLogger(r.logger.SetPretty(val))
}

func (r StacktraceLogger) WithField(key string, value interface{}) Logger {
	return NewStacktraceLogger(r.logger.WithField(key, value))
}

func (r StacktraceLogger) WithError(err error) Logger {
	err = r.wrapError(err)
	return NewStacktraceLogger(r.logger.WithError(err))
}

func (r StacktraceLogger) WithContext(ctx context.Context) Logger {
	return NewStacktraceLogger(r.logger.WithContext(ctx))
}

func (r StacktraceLogger) Debugf(format string, args ...interface{}) {
	r.logger.Debugf(format, r.wrapErrors(args)...)
}

func (r StacktraceLogger) Debug(args ...interface{}) {
	r.logger.Debug(r.wrapErrors(args)...)
}

func (r StacktraceLogger) Warningf(format string, v ...interface{}) {
	r.logger.Warningf(format, v...)
}

func (r StacktraceLogger) Warning(args ...interface{}) {
	r.logger.Warning(r.wrapErrors(args)...)
}

func (r StacktraceLogger) Errorf(format string, args ...interface{}) {
	r.logger.Errorf(format, r.wrapErrors(args)...)
}

func (r StacktraceLogger) Error(args ...interface{}) {
	r.logger.Error(r.wrapErrors(args)...)
}

func (r StacktraceLogger) Fatal(args ...interface{}) {
	r.logger.Fatal(r.wrapErrors(args)...)
}

func (r StacktraceLogger) Info(args ...interface{}) {
	r.logger.Info(r.wrapErrors(args)...)
}

func (r StacktraceLogger) Infof(format string, v ...interface{}) {
	r.logger.Infof(format, v...)
}

func (r StacktraceLogger) wrapErrors(args []interface{}) []interface{} {
	result := make([]interface{}, len(args))

	for i := range args {
		err, ok := args[i].(error)
		if !ok {
			result[i] = args[i]

			continue
		}

		result[i] = r.wrapError(err)
	}

	return result
}

func (r StacktraceLogger) wrapError(err error) error {
	if err == nil {
		return nil
	}

	trErr := UnwrapStacktrace(err)
	if trErr == nil {
		return err
	}

	st := trErr.StackTrace()
	if len(st) == 0 {
		return err
	}

	message := fmt.Sprintf("%+v", st)

	return errors.Wrap(err, message)
}
