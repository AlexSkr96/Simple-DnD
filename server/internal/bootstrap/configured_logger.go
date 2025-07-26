package bootstrap

import "github.com/AlexSkr96/Simple-DnD/pkg/logging"

func NewConfiguredLogger(conf logging.LoggerConfig) logging.Logger {
	logger := logging.NewZeroLogger(nil)

	return logger.
		SetLevel(conf.GetLogLevel()).
		SetPretty(conf.GetLogPretty()).
		WithField("module", conf.GetModule())
}
