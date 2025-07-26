package common

import (
	"context"
	"sync"
	"time"

	"go.uber.org/multierr"

	"github.com/AlexSkr96/Simple-DnD/pkg/servers"
)

func NewApp(servers ...servers.Server) *App {
	return &App{
		servers: servers,
	}
}

type App struct {
	servers []servers.Server
}

func (a App) Run(ctx context.Context) error {
	errCh := make(chan error)
	wg := sync.WaitGroup{}

	for _, server := range a.servers {
		wg.Add(1)

		go func() {
			wg.Done()

			err := server.Serve(ctx)
			if err != nil {
				errCh <- err
			}
		}()

		wg.Wait()
		time.Sleep(1 * time.Second)
	}

	var merr error

	for err := range errCh {
		merr = multierr.Append(merr, err)
	}

	close(errCh)

	return merr
}
