package models

import (
	"github.com/google/uuid"
)

type Spell struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name                 string    `gorm:"type:varchar(255);not null" json:"name"`
	Description          string    `gorm:"type:text" json:"description"`
	Damage               string    `gorm:"type:varchar(255)" json:"damage"`
	DamageTypeID         uuid.UUID `gorm:"type:uuid" json:"damage_type_id"`
	SavingThrowAbilityID uuid.UUID `gorm:"type:uuid" json:"saving_throw_ability_id"`
	Range                int       `gorm:"type:integer" json:"range"`
	SpellSlotLevel       int       `gorm:"type:integer" json:"spell_slot_level"`
	School               string    `gorm:"type:varchar(255)" json:"school"`

	// Relationships
	DamageType         DamageType `gorm:"foreignKey:DamageTypeID;references:ID" json:"damage_type,omitempty"`
	SavingThrowAbility Ability    `gorm:"foreignKey:SavingThrowAbilityID;references:ID" json:"saving_throw_ability,omitempty"`
}
