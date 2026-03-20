package errs

import "errors"

var (
	ErrEmptyMessage = errors.New("failed empty message")
	ErrNotFound     = errors.New("notification not found")
)
