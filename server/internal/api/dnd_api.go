package api

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/infra"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

const prefix = "/api/v1"

func NewServer(
	logger logging.Logger,
	bind string,
	rapi huma.API,
	repository infra.Repository,
) *Server {
	return &Server{
		logger:     logger,
		bind:       bind,
		repository: repository,
		api:        rapi,
	}
}

type Server struct {
	logger     logging.Logger
	bind       string
	api        huma.API
	repository infra.Repository
}

// nolint: funlen
func (s *Server) Serve(ctx context.Context) error {
	httpServer := &http.Server{ //nolint
		Handler: s.api.Adapter(),
		Addr:    s.bind,
	}

	huma.Register(s.api, huma.Operation{
		Method:        http.MethodGet,
		Path:          prefix + "/something/{id}",
		DefaultStatus: http.StatusOK,
		Errors:        []int{http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError},
		Tags:          []string{"something"},
		Summary:       "Get something",
		Description:   "Get something",
		OperationID:   "get-something",
	}, s.GetSomethingByID)

	go func() {
		s.logger.Info("starting dnd-api server on ", s.bind)

		err := httpServer.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	return nil
}
