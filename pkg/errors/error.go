package errors

import "fmt"

type Error struct {
	cause      *error
	shortTitle string
	message    string
}

func (e *Error) Cause() *error {
	return e.cause
}

func NewError(shortTitle string, message string, cause *error) *Error {
	return &Error{
		shortTitle: shortTitle,
		cause:      cause,
		message:    message,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.shortTitle, e.message)
}
