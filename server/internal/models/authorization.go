package models

import (
	"github.com/google/uuid"
	"time"
)

type RegisterRequest struct {
	XRequestID uuid.UUID `header:"X-Request-Id"`
	Email      string    `json:"email" required:"true"`
	Username   string    `json:"username" required:"true" minLength:"3" maxLength:"50"`
	Password   string    `json:"password" required:"true" minLength:"6"`
}

type LoginRequest struct {
	XRequestID uuid.UUID `header:"X-Request-Id"`
	Email      string    `json:"email" required:"true" format:"email"`
	Password   string    `json:"password" required:"true"`
}

type AuthResponse struct {
	ContentType string `header:"Content-Type"`
	Body        AuthResponseBody
}

type AuthResponseBody struct {
	UserID uuid.UUID `json:"user_id"`
	Token  string    `json:"token"`
}

type UserSession struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Token     string    `gorm:"size:255;not null;uniqueIndex" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      User      `gorm:"foreignKey:user_id" json:"user,omitempty"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}

type LogoutRequest struct {
	XRequestID    uuid.UUID `header:"X-Request-Id"`
	Authorization string    `header:"Authorization" required:"true"`
}
