package api

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	errpkg "github.com/AlexSkr96/Simple-DnD/pkg/errors"
	"github.com/danielgtaylor/huma/v2"
	"github.com/pkg/errors"
	"net/http"
)

func (s *Server) GetSomethingByID(ctx context.Context, p *models.GetSomethingByIDParams) (*models.GetSomethingByIDResponse, error) {
	resp, err := s.repository.FindSomethingByID(ctx, p.SomethingID)
	if errors.Is(err, errpkg.ErrNoRows) {
		return nil, huma.NewError(http.StatusNotFound, "Something not found")
	}

	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	return &models.GetSomethingByIDResponse{ContentType: JSONContentType, Body: *resp}, nil
}
