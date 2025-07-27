package auth

import "github.com/pkg/errors"

var (
	ErrUserAlreadyExists      = errors.New("user with this email already exists")
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
)
