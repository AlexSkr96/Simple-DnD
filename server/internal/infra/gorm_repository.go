package infra

import (
	"context"
	"github.com/AlexSkr96/Simple-DnD/internal/models"
	errpkg "github.com/AlexSkr96/Simple-DnD/pkg/errors"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

func NewGORMRepository(db *gorm.DB) *GORMRepository {
	return &GORMRepository{db: db}
}

type GORMRepository struct {
	db *gorm.DB
}

func (r *GORMRepository) FindSomethingByID(ctx context.Context, id uuid.UUID) (*models.Something, error) {
	//var something models.Something

	//err := r.db.WithContext(ctx).First(&something, id).Error
	//if errors.Is(err, gorm.ErrRecordNotFound) {
	//	return nil, errpkg.ErrNoRows // nolint: nilnil
	//} else if err != nil {
	//	return nil, errors.WithStack(err)
	//}

	return &models.Something{
		ID:          uuid.Nil,
		Description: "some description",
	}, nil
}

func (r *GORMRepository) CreateUser(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *GORMRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errpkg.ErrNoRows
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *GORMRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errpkg.ErrNoRows
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return &user, nil
}

func (r *GORMRepository) CreateSession(ctx context.Context, session *models.UserSession) error {
	err := r.db.WithContext(ctx).Create(session).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *GORMRepository) FindSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var session models.UserSession
	err := r.db.WithContext(ctx).
		Where("token = ? AND expires_at > ?", token, time.Now()).
		Preload("User").
		First(&session).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errpkg.ErrNoRows
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return &session, nil
}

func (r *GORMRepository) DeleteSession(ctx context.Context, token string) error {
	err := r.db.WithContext(ctx).Where("token = ?", token).Delete(&models.UserSession{}).Error
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
