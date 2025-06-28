package apperror

import (
	"errors"
	"net/http"
)

type HTTPResponseError struct {
	Status      int    `json:"code"`
	UserMessage string `json:"message"`
	Debug       string `json:"debug,omitempty"`
}

func ToHTTPError(err error, debug bool) HTTPResponseError {
	var initErr *InitError
	if errors.As(err, &initErr) {
		for _, mapping := range defaultErrorMap {
			if mapping.Code == initErr.Code {
				resp := HTTPResponseError{
					Status:      mapping.Status,
					UserMessage: mapping.UserMessage,
				}
				if debug {
					resp.Debug = initErr.Message
				}
				return resp
			}
		}

		status := initErr.HTTPStatus
		if status == 0 {
			status = http.StatusInternalServerError
		}

		return HTTPResponseError{
			Status:      status,
			UserMessage: initErr.Message,
			Debug:       ternary(debug, initErr.Message, ""),
		}
	}

	return HTTPResponseError{
		Status:      http.StatusInternalServerError,
		UserMessage: "Terjadi kesalahan server",
		Debug:       ternary(debug, err.Error(), ""),
	}
}

func ternary(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}
