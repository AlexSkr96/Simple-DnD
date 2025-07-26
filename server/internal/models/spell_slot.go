package models

import (
	"github.com/google/uuid"
)

type SpellSlot struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CharacterID uuid.UUID `gorm:"type:uuid;not null" json:"character_id"`
	Level       int       `gorm:"type:integer" json:"level"`
	Max         int       `gorm:"type:integer" json:"max"`
	SlotsLeft   int       `gorm:"type:integer" json:"slots_left"`
}
