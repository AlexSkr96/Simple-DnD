package models

import (
	"github.com/google/uuid"
)

type Feature struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Uses        int       `gorm:"type:integer" json:"uses"`
	Reset       string    `gorm:"type:varchar(255)" json:"reset"` // 'short rest', 'long rest', 'daily'
	Target      string    `gorm:"type:varchar(255)" json:"target"`
	Modificator int       `gorm:"type:integer" json:"modificator"`
}

type CharacterFeature struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CharacterID uuid.UUID `gorm:"type:uuid;not null" json:"character_id"`
	FeatureID   uuid.UUID `gorm:"type:uuid;not null" json:"feature_id"`

	// Relationships
	Character Character `gorm:"foreignKey:CharacterID;references:ID" json:"character,omitempty"`
	Feature   Feature   `gorm:"foreignKey:FeatureID;references:ID" json:"feature,omitempty"`
}
