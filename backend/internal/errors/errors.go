package errors

import "net/http"

type AppError struct {
	Code int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrNotFound         = &AppError{Code: http.StatusNotFound, Message: "not found"}
	ErrForbidden        = &AppError{Code: http.StatusForbidden, Message: "forbidden"}
	ErrUnauthorized     = &AppError{Code: http.StatusUnauthorized, Message: "unauthorized"}
	ErrConflict         = &AppError{Code: http.StatusConflict, Message: "conflict"}
	ErrInvalidInput     = &AppError{Code: http.StatusBadRequest, Message: "invalid input"}
	ErrUnsupportedMedia = &AppError{
		Code:    http.StatusUnsupportedMediaType,
		Message: "unsupported media type",
	}
)
