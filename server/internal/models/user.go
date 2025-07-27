package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string      `gorm:"size:255" json:"email"`
	Username     string      `gorm:"size:255" json:"username"`
	PasswordHash string      `gorm:"size:255" json:"-"`
	Characters   []Character `gorm:"foreignKey:UserID" json:"characters,omitempty"`
}

func (User) TableName() string {
	return "users"
}
