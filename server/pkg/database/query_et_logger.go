package database

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/pkg/errors"
	"time"
)

type ctxQETLStartTime struct{}

const QETLAnyQuery = "*"

func NewQueryETLogger(dbname string, logger logging.Logger, thresholdsQueryMS map[string]int64) *QueryETLogger {
	return &QueryETLogger{
		logger:            logger,
		dbname:            dbname,
		thresholdsQueryMS: thresholdsQueryMS,
	}
}

type QueryETLogger struct {
	logger            logging.Logger
	dbname            string
	thresholdsQueryMS map[string]int64
}

func (l QueryETLogger) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return context.WithValue(ctx, ctxQETLStartTime{}, time.Now()), nil
}

func (l QueryETLogger) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	start, ok := ctx.Value(ctxQETLStartTime{}).(time.Time)
	if !ok {
		l.logger.
			WithContext(ctx).
			WithField("query", query).
			WithField("db", l.dbname).
			Error(errors.New("broken query logging"))

		return ctx, nil
	}

	qname := GetQueryName(ctx, query)
	tookMS := time.Since(start).Milliseconds()

	logger := l.logger.
		WithField("db", l.dbname).
		WithField("query", query).
		WithField("args", args).
		WithField("took_ms", tookMS)

	if qname != query {
		logger = logger.WithField("qname", qname)
	}

	if l.isWarn(qname, tookMS) {
		logger.Warning("too long query")
	} else {
		logger.Debug("query execution time")
	}

	return ctx, nil
}

func (l QueryETLogger) isWarn(qname string, tookMS int64) bool {
	threshold, ok := l.thresholdsQueryMS[QETLAnyQuery]
	if ok && threshold < tookMS {
		return true
	}

	threshold, ok = l.thresholdsQueryMS[qname]
	if ok && threshold < tookMS {
		return true
	}

	return false
}
