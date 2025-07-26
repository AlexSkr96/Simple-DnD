package models

import (
	"github.com/google/uuid"
)

type Weapon struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name         string     `gorm:"type:varchar(255);not null" json:"name"`
	IsRanged     bool       `gorm:"type:boolean" json:"is_ranged"`
	OptimalRange int        `gorm:"type:integer" json:"optimal_range"`
	MaxRange     int        `gorm:"type:integer" json:"max_range"`
	Cost         int        `gorm:"type:integer" json:"cost"`
	Damage       string     `gorm:"type:varchar(255)" json:"damage"`
	DamageTypeID *uuid.UUID `gorm:"type:uuid" json:"damage_type_id"`

	// Relationships
	DamageType     DamageType     `gorm:"foreignKey:DamageTypeID;references:ID" json:"damage_type,omitempty"`
	WeaponProperty WeaponProperty `gorm:"foreignKey:WeaponPropertyID;references:ID" json:"weapon_property,omitempty"`
}
