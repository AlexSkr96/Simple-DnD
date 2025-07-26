package health

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/pkg/logging"
	"github.com/pkg/errors"
	"net"
	"net/http"
)

func NewServer(
	bind string,
	logger logging.Logger,
) *Server {
	return &Server{
		bind:   bind,
		logger: logger,
	}
}

type Server struct {
	bind   string
	logger logging.Logger
}

func (r Server) Serve(_ context.Context) error {
	handleFunc := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(`{"message": "OK"}`))
		if err != nil {
			r.logger.Error(err)
		}
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/readiness", handleFunc)
	serveMux.HandleFunc("/liveness", handleFunc)

	listener, err := net.Listen("tcp", r.bind)
	if err != nil {
		return errors.WithStack(err)
	}

	r.logger.Info("starting health server on ", r.bind)

	err = http.Serve(listener, serveMux) //nolint: gosec
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
