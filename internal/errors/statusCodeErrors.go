package error

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound             = errors.New("record not found")
	ErrConflict             = errors.New("an conflict occurred")
	ErrInternal             = errors.New("an error occurred processing the request")
	ErrBadRequest           = errors.New("bad request")
	ErrUnprocessableContent = errors.New("unprocessable content")
	ErrUnauthorized         = errors.New("access unauthorized")
)

func statusCodes() map[int][]error {
	return map[int][]error{
		http.StatusNotFound: {
			ErrNotFound,
		},
		http.StatusBadRequest: {
			ErrBadRequest,
		},
		http.StatusConflict: {
			ErrConflict,
		},
		http.StatusInternalServerError: {
			ErrInternal,
		},
		http.StatusUnprocessableEntity: {
			ErrUnprocessableContent,
		},
	}
}
