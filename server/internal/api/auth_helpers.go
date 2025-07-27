package api

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	"github.com/AlexSkr96/Simple-DnD/pkg/middleware"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

func (s *Server) getCurrentUser(ctx context.Context) (*models.User, error) {
	user, ok := middleware.GetUserFromContext(ctx)
	if !ok {
		return nil, huma.NewError(http.StatusUnauthorized, "Authentication required")
	}
	return user, nil
}
