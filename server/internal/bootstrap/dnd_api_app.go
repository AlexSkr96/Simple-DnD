package bootstrap

import (
	"github.com/AlexSkr96/Simple-DnD/internal/api"
	"github.com/AlexSkr96/Simple-DnD/internal/configs"
	"github.com/AlexSkr96/Simple-DnD/pkg/common"
	"github.com/AlexSkr96/Simple-DnD/pkg/health"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/AlexSkr96/Simple-DnD/pkg/servers"
)

func NewDnDAPIApp(
	conf *configs.DnDAPIConfig,
	logger logging.Logger,
	server *api.Server,
) *common.App {
	appServers := []servers.Server{
		health.NewServer(conf.HealthBind, logger),
		server,
	}

	return common.NewApp(
		appServers...,
	)
}
