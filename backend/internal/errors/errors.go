package errors

import "errors"

var (
	ErrForbidden        = errors.New("forbidden")
	ErrUnsupportedMedia = errors.New("unsupported media")
	ErrNotFound         = errors.New("not found")
	ErrFileTooLarge     = errors.New("file too large")
)
