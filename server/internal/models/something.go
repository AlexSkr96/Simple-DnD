package models

import "github.com/google/uuid"

// nolint: tagalign
type GetSomethingByIDParams struct {
	XRequestID  uuid.UUID `header:"X-Request-Id"`
	SomethingID uuid.UUID `path:"id"              required:"true"`
}

type GetSomethingByIDResponse struct {
	ContentType string `header:"Content-Type"`
	Body        Something
}

type Something struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Description string    `json:"description"`
}

func (s Something) TableName() string {
	return "something"
}
