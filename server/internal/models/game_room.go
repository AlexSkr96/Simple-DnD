package models

import (
	"github.com/google/uuid"
)

type GameRoom struct {
	ID      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name    string    `gorm:"size:255;not null" json:"name"`
	OwnerID uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`

	// Relationships
	GameMaster User `gorm:"foreignKey:OwnerID" json:"game_master,omitempty"`
}

func (GameRoom) TableName() string {
	return "game_rooms"
}

type GameRoomParticipants struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	GameRoomID  uuid.UUID `gorm:"type:uuid;not null" json:"party_id"`
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	CharacterID uuid.UUID `gorm:"type:uuid;not null" json:"character_id"`

	// Relationships
	GameRoom  GameRoom  `gorm:"foreignKey:GameRoomID" json:"game_room,omitempty"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Character Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
}

func (GameRoomParticipants) TableName() string {
	return "game_room_participants"
}
