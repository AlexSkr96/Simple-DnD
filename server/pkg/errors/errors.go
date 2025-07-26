package errors

import (
	"errors"
)

var (
	ErrNoRows          = errors.New("no rows in result set")
	ErrNilIsNotAllowed = errors.New("nil parameter is not allowed")
)
