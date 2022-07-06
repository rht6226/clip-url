package errors

import "fmt"

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error.",
	}
}

func NewNotFound(reason string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("NotFound: %v", reason),
	}
}
