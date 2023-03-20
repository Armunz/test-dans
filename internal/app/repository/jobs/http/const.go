package http

import "errors"

const PAGINATION = 10

var (
	// ErrCtxTimeout - error returned when context is timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrEmptyID - error returned when id is empty
	ErrEmptyID = errors.New("id is empty")
)
