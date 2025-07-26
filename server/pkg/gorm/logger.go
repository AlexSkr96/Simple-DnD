package gorm

import (
	"context"
	"fmt"
	"time"

	loggergorm "gorm.io/gorm/logger"

	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
)

func NewLogger(logger logging.Logger) *Logger {
	return &Logger{
		logger: logger.WithField("module", "gorm-db"),
	}
}

type Logger struct {
	logger logging.Logger
}

func (l *Logger) LogMode(loggergorm.LogLevel) loggergorm.Interface {
	panic("use common logger for setting mode")
}

func (l Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.WithContext(ctx).Infof(msg, data...) //nolint
}

func (l Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.WithContext(ctx).Warningf(msg, data...) //nolint
}

func (l Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.WithContext(ctx).Errorf(msg, data...) //nolint
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logger.GetLevel() != logging.LevelDebug {
		return
	}

	elapsed := time.Since(begin)

	sql, rows := fc()

	l.logger.Debug(
		fmt.Sprintf(
			"[%.3fms][rows:%d] %s",
			float64(elapsed.Nanoseconds())/1e6, //nolint
			rows,
			sql,
		),
	)
}
