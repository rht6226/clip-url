package errors

import "fmt"

func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

func NewInternal(reason string) *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error: %v", reason),
	}
}

func NewNotFound(reason string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("NotFound: %v", reason),
	}
}
