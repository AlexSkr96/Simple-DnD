package main

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/api"
	"github.com/AlexSkr96/Simple-DnD/internal/bootstrap"
	"github.com/AlexSkr96/Simple-DnD/internal/configs"
	"github.com/AlexSkr96/Simple-DnD/internal/infra"
	"github.com/AlexSkr96/Simple-DnD/internal/services/auth"
	"github.com/AlexSkr96/Simple-DnD/pkg/common"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
)

func main() {
	ctx := context.Background()

	app, cleanup, err := initApp()

	defer func() {
		if cleanup != nil {
			cleanup()
		}
	}()

	if err != nil {
		logging.PanicStack(err)
	}

	err = app.Run(ctx)
	if err != nil {
		logging.PanicStack(err)
	}
}

func initApp() (app *common.App, cleanup func(), err error) {
	dndAPIConfig, err := configs.DnDAPIConfigConfig()
	if err != nil {
		return nil, nil, err
	}

	logger := bootstrap.NewConfiguredLogger(dndAPIConfig)
	dndRouter := bootstrap.NewDnDAPIRouter(logger)
	gormConfig := dndAPIConfig.GORMConfig
	driverName := bootstrap.NewPGDBDriverName(gormConfig, logger)
	db, cleanup, err := bootstrap.NewGORMDB(logger, gormConfig, driverName)
	if err != nil {
		cleanup()
		return nil, nil, err
	}

	repository := infra.NewGORMRepository(db)
	authService := auth.NewService(repository)
	server := api.NewServer(logger, dndAPIConfig.ServerBind, dndRouter, repository, authService)
	app = bootstrap.NewDnDAPIApp(dndAPIConfig, logger, server)
	return app, func() {
		cleanup()
	}, nil
}
