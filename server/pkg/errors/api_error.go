package errors

import (
	"github.com/danielgtaylor/huma/v2"
)

type APIError struct {
	Status    int                 `json:"-"`
	Title     string              `json:"title"`
	Detail    string              `json:"detail"`
	RequestID string              `json:"requestId"`
	Origin    string              `json:"origin"`
	Errors    []*huma.ErrorDetail `doc:"Optional list of individual error details" json:"errors,omitempty"`
}

func (e *APIError) Error() string {
	return e.Detail
}

func (e *APIError) GetStatus() int {
	return e.Status
}

func (e *APIError) ContentType(string) string {
	return "application/problem+json"
}
