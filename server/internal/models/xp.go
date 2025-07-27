package models

import (
	"github.com/google/uuid"
	"time"
)

type ExperienceGrant struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CharacterID uuid.UUID `gorm:"type:uuid;not null" json:"character_id"`
	Amount      int       `gorm:"not null" json:"amount"`
	Reason      string    `gorm:"size:255" json:"reason"`
	GrantedBy   uuid.UUID `gorm:"type:uuid;not null" json:"granted_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	Character   Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
	GameMaster  User      `gorm:"foreignKey:GrantedBy" json:"game_master,omitempty"`
}

type GrantExperienceParams struct {
	XRequestID  uuid.UUID `header:"X-Request-Id"`
	GameRoomID  uuid.UUID `path:"gameRoomId" required:"true"`
	CharacterID uuid.UUID `path:"characterId" required:"true"`
	Body        GrantExperienceBody
}

type GrantExperienceBody struct {
	Amount int    `json:"amount" required:"true" minimum:"1"`
	Reason string `json:"reason" required:"true"`
}
