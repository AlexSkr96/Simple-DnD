package models

import (
	"github.com/google/uuid"
)

type Ability struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name string    `gorm:"size:255;not null" json:"name"`

	// Relationships
	CharacterAbilities []CharacterAbility `gorm:"foreignKey:AbilityID" json:"-"`
	Skills             []Skill            `gorm:"foreignKey:BaseAbilityID" json:"-"`
}

type CharacterAbility struct {
	ID                      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	AbilityID               uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_character_ability" json:"ability_id"`
	CharacterID             uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_character_ability" json:"character_id"`
	Value                   int       `json:"value"`
	IsProficientSavingThrow bool      `json:"is_proficient_saving_throw"`

	// Relationships
	Ability   Ability   `gorm:"foreignKey:AbilityID" json:"ability,omitempty"`
	Character Character `gorm:"foreignKey:CharacterID" json:"-"`
}
