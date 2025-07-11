package apperror

import (
	"errors"
	"fmt"
)

type InitError struct {
	Code           string
	Message        string
	Err            error
	HTTPStatus     int    // Optional: custom HTTP status code
	ResponseStatus string // Optional: custom JSON "status" (e.g. "fail", "unauthorized")
}

func (e *InitError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s | %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func New(code, message string, err error, httpStatus ...int) *InitError {
	status := 0
	if len(httpStatus) > 0 {
		status = httpStatus[0]
	}
	return &InitError{
		Code:       code,
		Message:    message,
		Err:        err,
		HTTPStatus: status,
	}
}

func Is(err error, code string) bool {
	var e *InitError
	return errors.As(err, &e) && e.Code == code
}

func (e *InitError) WithResponseStatus(status string) *InitError {
	e.ResponseStatus = status
	return e
}

func NewWithStatus(code, message string, err error, httpStatus int, responseStatus string) *InitError {
	return &InitError{
		Code:           code,
		Message:        message,
		Err:            err,
		HTTPStatus:     httpStatus,
		ResponseStatus: responseStatus,
	}
}
