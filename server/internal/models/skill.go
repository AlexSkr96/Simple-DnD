package models

import "github.com/google/uuid"

type Skill struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name          string    `gorm:"size:255;not null" json:"name"`
	BaseAbilityID uuid.UUID `gorm:"type:uuid" json:"base_ability_id"`

	// Relationships
	BaseAbility     Ability          `gorm:"foreignKey:BaseAbilityID" json:"base_ability,omitempty"`
	CharacterSkills []CharacterSkill `gorm:"foreignKey:SkillID" json:"-"`
}

type CharacterSkill struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	SkillID      uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_character_skill" json:"skill_id"`
	CharacterID  uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_character_skill" json:"character_id"`
	Value        int       `json:"value"`
	IsProficient bool      `json:"is_proficient"`

	// Relationships
	Skill     Skill     `gorm:"foreignKey:SkillID" json:"skill,omitempty"`
	Character Character `gorm:"foreignKey:CharacterID" json:"-"`
}
