package apperror

import (
	"errors"
	"fmt"
)

type InitError struct {
	Code    string
	Message string
	Err     error
}

func (e *InitError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s | %v", e.Code, e.Message, e.Err)
	}

	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func New(code, message string, err error) *InitError {
	return &InitError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func Is(err error, code string) bool {
	var e *InitError
	if errors.As(err, &e) {
		return e.Code == code
	}

	return false
}
