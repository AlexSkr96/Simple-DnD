package api

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	"github.com/AlexSkr96/Simple-DnD/internal/services/auth"
	errpkg "github.com/AlexSkr96/Simple-DnD/pkg/errors"
	"github.com/danielgtaylor/huma/v2"
	"github.com/pkg/errors"
	"net/http"
	"strings"
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

func (s *Server) Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error) {
	resp, err := s.authService.Register(ctx, req)
	if errors.Is(err, auth.ErrUserAlreadyExists) {
		return nil, huma.NewError(http.StatusConflict, err.Error())
	}

	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	return &models.AuthResponse{
		ContentType: JSONContentType,
		Body:        *resp,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error) {
	resp, err := s.authService.Login(ctx, req)
	if errors.Is(err, auth.ErrInvalidEmailOrPassword) {
		return nil, huma.NewError(http.StatusUnauthorized, err.Error())
	}

	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	return &models.AuthResponse{
		ContentType: JSONContentType,
		Body:        *resp,
	}, nil
}

func (s *Server) Logout(ctx context.Context, req *models.LogoutRequest) (*struct{}, error) {
	token := strings.TrimPrefix(req.Authorization, "Bearer ")

	err := s.authService.Logout(ctx, token)
	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	return &struct{}{}, nil
}

func (s *Server) GrantExperience(ctx context.Context, p *models.GrantExperienceParams) (*struct{}, error) {
	user, err := s.getCurrentUser(ctx)
	if err != nil {
		return nil, err
	}

	gmID, err := s.repository.FindGameRoomOwnerID(ctx, p.GameRoomID)
	if errors.Is(err, errpkg.ErrNoRows) {
		return nil, huma.NewError(http.StatusNotFound, "Room not found")
	}

	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	if gmID != user.ID {
		return nil, huma.NewError(http.StatusForbidden, "You are not allowed to grant experience to this room")
	}

	_, err = s.repository.FindCharacterByIDAndRoomID(ctx, p.CharacterID, p.GameRoomID)
	if errors.Is(err, errpkg.ErrNoRows) {
		return nil, huma.NewError(http.StatusNotFound, "Character not found")
	}

	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	grant := &models.ExperienceGrant{
		CharacterID: p.CharacterID,
		Amount:      p.Body.Amount,
		Reason:      p.Body.Reason,
		GrantedBy:   user.ID,
	}

	err = s.repository.GrantExperience(ctx, grant)
	if err != nil {
		s.logger.Error(err)
		return nil, huma.NewError(http.StatusInternalServerError, "Internal server error")
	}

	return &struct{}{}, nil
}
