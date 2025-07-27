package infra

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	"github.com/google/uuid"
)

// nolint: interfacebloat
type Repository interface {
	FindSomethingByID(ctx context.Context, id uuid.UUID) (*models.Something, error)
	// User methods
	CreateUser(ctx context.Context, user *models.User) error
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	// Auth methods
	CreateSession(ctx context.Context, session *models.UserSession) error
	FindSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
	DeleteSession(ctx context.Context, token string) error
	// GameRoom methods
	FindGameRoomByID(ctx context.Context, id uuid.UUID) (*models.GameRoom, error)
	FindGameRoomOwnerID(ctx context.Context, roomID uuid.UUID) (uuid.UUID, error)
	FindCharacterByIDAndRoomID(ctx context.Context, id uuid.UUID, roomID uuid.UUID) (*models.Character, error)
	// Experience methods
	GrantExperience(ctx context.Context, grant *models.ExperienceGrant) error
}
