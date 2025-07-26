package infra

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	errpkg "github.com/AlexSkr96/Simple-DnD/pkg/errors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type GORMRepository struct {
	db *gorm.DB
}

func (r *GORMRepository) FindSomethingByID(ctx context.Context, id uuid.UUID) (*models.Something, error) {
	var something models.Something

	err := r.db.WithContext(ctx).First(&something, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errpkg.ErrNoRows // nolint: nilnil
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &something, nil
}

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	return &GORMRepository{db: db}
}
