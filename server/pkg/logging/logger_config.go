package logging

type LoggerConfig interface {
	GetLogPretty() bool
	GetLogLevel() Level
	GetModule() string
}
