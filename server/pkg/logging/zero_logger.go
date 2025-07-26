package logging

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewZeroLogger(zlogger *zerolog.Logger) *ZeroLogger {
	if zlogger == nil {
		t := log.With().Logger()
		zlogger = &t
	}

	zerologger := &ZeroLogger{
		logger: zlogger,
	}

	return zerologger
}

type ZeroLogger struct {
	level  Level
	logger *zerolog.Logger
}

func (r *ZeroLogger) GetLevel() Level {
	return r.level
}

func (r *ZeroLogger) SetLevel(level Level) Logger {
	zlevel, err := zerolog.ParseLevel(string(level))
	if err != nil {
		r.Error(err)

		return r
	}

	zlogger := r.logger.Level(zlevel)
	l := r.new(&zlogger)
	l.level = level

	return l
}

func (r *ZeroLogger) SetPretty(val bool) Logger {
	var writer io.Writer
	if val {
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	} else {
		writer = os.Stderr
	}

	zlogger := r.logger.Output(writer)

	return r.new(&zlogger)
}

func (r *ZeroLogger) WithContext(ctx context.Context) Logger {
	return r
}

func (r *ZeroLogger) WithField(key string, value interface{}) Logger {
	zlogger := r.logger.With().Fields(map[string]interface{}{key: value}).Logger()

	return r.new(&zlogger)
}

func (r *ZeroLogger) WithError(err error) Logger {
	zlogger := r.logger.With().Err(err).Logger()

	return r.new(&zlogger)
}

func (r *ZeroLogger) Debug(args ...interface{}) {
	event := r.logger.Debug()
	r.sendEvent(event, args...)
}

func (r *ZeroLogger) Debugf(format string, v ...interface{}) {
	r.logger.Debug().Msgf(format, v...)
}

func (r *ZeroLogger) Warning(args ...interface{}) {
	event := r.logger.Warn()
	r.sendEvent(event, args...)
}

func (r *ZeroLogger) Warningf(format string, v ...interface{}) {
	r.logger.Warn().Msgf(format, v...)
}

func (r *ZeroLogger) Error(args ...interface{}) {
	event := r.logger.Error()
	r.sendEvent(event, args...)
}

func (r *ZeroLogger) Errorf(format string, v ...interface{}) {
	r.logger.Error().Msgf(format, v...)
}

func (r *ZeroLogger) Fatal(args ...interface{}) {
	event := r.logger.Fatal()
	r.sendEvent(event, args...)
}

func (r *ZeroLogger) Info(args ...interface{}) {
	event := r.logger.Info()
	r.sendEvent(event, args...)
}

func (r *ZeroLogger) Infof(format string, v ...interface{}) {
	r.logger.Info().Msgf(format, v...)
}

func (r *ZeroLogger) sendEvent(event *zerolog.Event, args ...interface{}) {
	event.Msg(fmt.Sprint(args...))
}

func (r *ZeroLogger) new(zlogger *zerolog.Logger) *ZeroLogger {
	logger := NewZeroLogger(zlogger)
	logger.level = r.level

	return logger
}
