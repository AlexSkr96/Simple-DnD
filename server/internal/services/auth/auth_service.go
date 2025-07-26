package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/AlexSkr96/Simple-DnD/internal/infra"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	errpkg "github.com/AlexSkr96/Simple-DnD/pkg/errors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	repository infra.Repository
}

func NewService(repository infra.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponseBody, error) {
	_, err := s.repository.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	if !errors.Is(err, errpkg.ErrNoRows) {
		return nil, errors.WithStack(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}

	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	token, err := s.generateToken()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	now := time.Now()

	session := &models.UserSession{
		ID:        uuid.New(),
		Token:     token,
		ExpiresAt: now.Add(24 * time.Hour),
		CreatedAt: now,
		User:      *user,
	}

	err = s.repository.CreateSession(ctx, session)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.AuthResponseBody{
		UserID: user.ID,
		Token:  session.Token,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponseBody, error) {
	user, err := s.repository.FindUserByEmail(ctx, req.Email)
	if errors.Is(err, errpkg.ErrNoRows) {
		return nil, ErrInvalidEmailOrPassword
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidEmailOrPassword
	}

	token, err := s.generateToken()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	now := time.Now()

	session := &models.UserSession{
		ID:        uuid.New(),
		Token:     token,
		ExpiresAt: now.Add(24 * time.Hour),
		CreatedAt: now,
		User:      *user,
	}

	err = s.repository.CreateSession(ctx, session)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &models.AuthResponseBody{
		UserID: user.ID,
		Token:  session.Token,
	}, nil
}

func (s *Service) ValidateToken(ctx context.Context, token string) (*models.User, error) {
	session, err := s.repository.FindSessionByToken(ctx, token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &session.User, nil
}

func (s *Service) Logout(ctx context.Context, token string) error {
	return s.repository.DeleteSession(ctx, token)
}

func (s *Service) generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
