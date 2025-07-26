package models

import (
	"github.com/google/uuid"
)

type DamageType struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name string    `gorm:"size:255;not null" json:"name"`
}
