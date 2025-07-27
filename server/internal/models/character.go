package models

import (
	"github.com/google/uuid"
)

type Character struct {
	ID                uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name              string    `gorm:"size:255;not null" json:"name"`
	Index             string    `gorm:"size:255" json:"index"`
	Class             string    `gorm:"size:255" json:"class"`
	Race              string    `gorm:"size:255" json:"race"`
	Alignment         string    `gorm:"size:255" json:"alignment"`
	UserID            uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User              User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ArmourClass       int       `gorm:"not null" json:"armour_class"`
	CurrentExperience int       `gorm:"default:0" json:"current_experience"`
	DeathSavePos      int       `gorm:"default:0" json:"death_save_pos"`
	DeathSaveNeg      int       `gorm:"default:0" json:"death_save_neg"`
	HitDiceLeft       string    `gorm:"size:255" json:"hit_dice_left"`
	TotalHitDice      string    `gorm:"size:255" json:"total_hit_dice"`
	MaxHP             int       `gorm:"not null" json:"max_hp"`
	CurrentHP         int       `gorm:"not null" json:"current_hp"`
	TempHP            int       `gorm:"not null;default:0" json:"temp_hp"`
	ProficiencyBonus  int       `json:"proficiency_bonus"`
	Inspiration       bool      `gorm:"default:false" json:"inspiration"`
	Speed             int       `gorm:"not null;default:30" json:"speed"`

	// Relationships
	Abilities      []CharacterAbility `gorm:"foreignKey:CharacterID" json:"abilities,omitempty"`
	Skills         []CharacterSkill   `gorm:"foreignKey:CharacterID" json:"skills,omitempty"`
	Features       []CharacterFeature `gorm:"foreignKey:CharacterID" json:"features,omitempty"`
	PreparedSpells []PreparedSpell    `gorm:"foreignKey:CharacterID" json:"prepared_spells,omitempty"`
	SpellSlots     []SpellSlot        `gorm:"foreignKey:CharacterID" json:"spell_slots,omitempty"`
}

func (Character) TableName() string {
	return "characters"
}
