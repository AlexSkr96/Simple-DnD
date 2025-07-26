package configs

import (
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/kelseyhightower/envconfig"
)

type DnDAPIConfig struct {
	*GORMConfig

	Module string `default:"dnd-api" split_words:"true"`

	LogPretty bool          `envconfig:"LOG_PRETTY"`
	LogLevel  logging.Level `envconfig:"LOG_LEVEL"  required:"true"`

	HealthBind string `envconfig:"HEALTH_BIND"  required:"true"`
	ServerBind string `envconfig:"SERVER_BIND"  required:"true"`
}

func (c DnDAPIConfig) GetLogPretty() bool {
	return c.LogPretty
}

func (c DnDAPIConfig) GetLogLevel() logging.Level {
	return c.LogLevel
}

func (c DnDAPIConfig) GetModule() string {
	return c.Module
}

func DnDAPIConfigConfig() (*DnDAPIConfig, error) {
	conf := &DnDAPIConfig{}

	err := envconfig.Process("", conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
