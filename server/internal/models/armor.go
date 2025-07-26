package models

import (
	"github.com/google/uuid"
)

type Armor struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name                string    `gorm:"type:varchar(255);not null" json:"name"`
	AC                  int       `gorm:"type:integer;not null" json:"ac"`
	StrReq              int       `gorm:"type:integer;not null;default:0" json:"str_req"`
	StealthDisadvantage bool      `gorm:"type:boolean;not null;default:false" json:"stealth_disadvantage"`
	Weight              int       `gorm:"type:integer" json:"weight"`
	Cost                int       `gorm:"type:integer" json:"cost"`
}
