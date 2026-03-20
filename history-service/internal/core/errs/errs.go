package errs

import "errors"

var (
	ErrNotFound = errors.New("notification not found")
	ErrInvalidID = errors.New("invalid id")
)
