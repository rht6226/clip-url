package errors

import (
	"errors"
	"net/http"
)

// Type holds a type string and integer code for the error
type Type string

// "Set" of valid errorTypes
const (
	BadRequest Type = "BADREQUEST" // Validation errors / BadInput
	Internal   Type = "INTERNAL"   // Server (500) and fallback errors
	NotFound   Type = "NOTFOUND"   // For not finding resource
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case BadRequest:
		return http.StatusBadRequest
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}
