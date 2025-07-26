package models

import (
	"github.com/google/uuid"
)

type CharacterPreparedSpell struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SpellID     uuid.UUID `gorm:"type:uuid;not null" json:"spell_id"`
	CharacterID uuid.UUID `gorm:"type:uuid;not null" json:"character_id"`

	// Relationships
	Spell Spell `gorm:"foreignKey:SpellID;references:ID" json:"spell,omitempty"`
}
