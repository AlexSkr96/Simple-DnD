package infra

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	"github.com/google/uuid"
)

// nolint: interfacebloat
type Repository interface {
	FindSomethingByID(ctx context.Context, id uuid.UUID) (*models.Something, error)
}
