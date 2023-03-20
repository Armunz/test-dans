package handler

import "errors"

var (
	// ErrCtxTimeout - error returned when context is timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrRepoNil - error returned when repository is nil
	ErrRepoNil = errors.New("repository is nil")

	// ErrEmptyID - error returned when ID is empty
	ErrEmptyID = errors.New("id is empty")

	// ErrPageInvalid - error returned when page is invalid
	ErrPageInvalid = errors.New("page is invalid")
)
