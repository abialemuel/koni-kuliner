package entity

import (
	"net/http"
)

type CustomError struct {
	Message    string `json:"message"`
	Code       int    `json:"code"`
	Field      string `json:"field,omitempty"`
	HttpStatus int    `json:"http_status"`
}

func (e CustomError) Error() string {
	return e.Message
}

var (
	InvalidTokenError = CustomError{
		Message:    "Invalid token",
		Code:       67101,
		Field:      "",
		HttpStatus: http.StatusUnauthorized,
	}
	UserUnauthorizedError = CustomError{
		Message:    "User unauthorized",
		Code:       67102,
		Field:      "",
		HttpStatus: http.StatusForbidden,
	}
	FailedDecodeJSONError = CustomError{
		Message:    "Invalid parameters",
		Code:       67103,
		Field:      "",
		HttpStatus: http.StatusBadRequest,
	}
	AuthorizationFailedError = CustomError{
		Message:    "Authorization failed",
		Code:       67104,
		HttpStatus: http.StatusUnauthorized,
	}
	UnprocessableEntityError = CustomError{
		Message:    "Unprocessable Entity",
		Code:       67105,
		HttpStatus: http.StatusUnprocessableEntity,
	}
)
