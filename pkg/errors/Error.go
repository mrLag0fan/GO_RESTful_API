package errors

import "fmt"

type Error struct {
	cause      *Error
	shortTitle string
	message    string
}

func (e *Error) Cause() *Error {
	return e.cause
}

func NewError(shortTitle string, message string, cause *Error) *Error {
	return &Error{
		shortTitle: shortTitle,
		cause:      cause,
		message:    message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.shortTitle, e.message)
}
