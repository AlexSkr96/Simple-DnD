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

func (r *GORMRepository) GrantExperience(ctx context.Context, grant *models.ExperienceGrant) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(grant).Error; err != nil {
			return errors.WithStack(err)
		}

		// Update character's current experience
		if err := tx.Model(&models.Character{}).
			Where("id = ?", grant.CharacterID).
			Update("current_experience", gorm.Expr("current_experience + ?", grant.Amount)).Error; err != nil {
			return errors.WithStack(err)
		}

		return nil
	})
}

func (r *GORMRepository) FindCharacterByIDAndRoomID(ctx context.Context, id uuid.UUID, roomID uuid.UUID) (*models.Character, error) {
	var character models.Character

	err := r.db.WithContext(ctx).
		Preload("Abilities.Ability").
		Preload("Skills.Skill.BaseAbility").
		Preload("Features.Feature").
		Preload("PreparedSpells.Spell").
		Preload("SpellSlots").
		Joins("JOIN game_room_participants ON game_room_participants.character_id = characters.id").
		Where("characters.id = ? AND game_room_participants.game_room_id = ?", id, roomID).
		First(&character).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errpkg.ErrNoRows
	} else if err != nil {
		return nil, errors.WithStack(err)
	}

	return &character, nil
}

func (r *GORMRepository) FindGameRoomByID(ctx context.Context, id uuid.UUID) (*models.GameRoom, error) {
	var gameRoom models.GameRoom
	err := r.db.WithContext(ctx).First(&gameRoom, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errpkg.ErrNoRows
	} else if err != nil {
		return nil, errors.WithStack(err)
	}
	return &gameRoom, nil
}

func (r *GORMRepository) FindGameRoomOwnerID(ctx context.Context, roomID uuid.UUID) (uuid.UUID, error) {
	var ownerID uuid.UUID
	err := r.db.WithContext(ctx).Model(&models.GameRoom{}).
		Where("id = ?", roomID).
		Select("owner_id").
		Take(&ownerID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return uuid.Nil, errpkg.ErrNoRows
	} else if err != nil {
		return uuid.Nil, errors.WithStack(err)
	}

	return ownerID, nil
}
