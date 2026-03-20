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

	httpErr, ok := errToHTTP[err]
	if ok {
		return httpErr
	}
	return &HTTPError{Code: http.StatusInternalServerError, Message: "internal server error"}
}
