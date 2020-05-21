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
	ProductNotFoundError = CustomError{
		Message:    "Product not found",
		Code:       67111,
		HttpStatus: http.StatusNotFound,
	}
	ProductNotValidError = CustomError{
		Message:    "Product not valid",
		Code:       67112,
		HttpStatus: http.StatusUnprocessableEntity,
	}
	OutletNotFoundError = CustomError{
		Message:    "Outlet not found",
		Code:       67111,
		HttpStatus: http.StatusNotFound,
	}
	OutletNotValidError = CustomError{
		Message:    "Outlet not valid",
		Code:       67112,
		HttpStatus: http.StatusUnprocessableEntity,
	}
	BrandNotFoundError = CustomError{
		Message:    "Brand not found",
		Code:       67111,
		HttpStatus: http.StatusNotFound,
	}
	BrandNotValidError = CustomError{
		Message:    "Brand not valid",
		Code:       67112,
		HttpStatus: http.StatusUnprocessableEntity,
	}
	CustomerNotFoundError = CustomError{
		Message:    "Customer not found",
		Code:       67111,
		HttpStatus: http.StatusNotFound,
	}
	OutletProductNotFoundError = CustomError{
		Message:    "Outlet Product not found",
		Code:       67111,
		HttpStatus: http.StatusNotFound,
	}
	TransactionNotFoundError = CustomError{
		Message:    "Transaction not found",
		Code:       67111,
		HttpStatus: http.StatusNotFound,
	}
	TransactionNotValidError = CustomError{
		Message:    "Transaction not valid",
		Code:       67112,
		HttpStatus: http.StatusUnprocessableEntity,
	}
)
