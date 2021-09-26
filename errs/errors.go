package errs

import "net/http"

// Errs ...
type AppErrs struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

// NewNotFoundError ...
func NewNotFoundError(msg string) *AppErrs {
	return &AppErrs{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

// NewServerError ...
func NewServerError(msg string) *AppErrs {
	return &AppErrs{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}
