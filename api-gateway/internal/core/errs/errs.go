package errs

import (
	"errors"
	"net/http"
)

var (
	ErrEmptyMessage    = errors.New("failed empty message")
	ErrInvalidUserID   = errors.New("invalid user_id")
	ErrTooManyRequests = errors.New("too many requests")
)

type HTTPError struct {
	Code    int
	Message string
}

func (h *HTTPError) Error() string {
	return h.Message
}

var errToHTTP = map[error]*HTTPError{
	ErrEmptyMessage:    {Code: http.StatusBadRequest, Message: "empty message"},
	ErrInvalidUserID:   {Code: http.StatusBadRequest, Message: "invalid user_id"},
	ErrTooManyRequests: {Code: http.StatusTooManyRequests, Message: "too many requests"},
}

func ToHTTPError(err error) *HTTPError {
	switch {
	case errors.Is(err, ErrEmptyMessage):
		return &HTTPError{Code: http.StatusBadRequest, Message: "empty message"}
	case errors.Is(err, ErrInvalidUserID):
		return &HTTPError{Code: http.StatusBadRequest, Message: "invalid user_id"}
	case errors.Is(err, ErrTooManyRequests):
		return &HTTPError{Code: http.StatusTooManyRequests, Message: "too many requests"}
	default:
		return &HTTPError{Code: http.StatusInternalServerError, Message: "internal server error"}
	}
}
