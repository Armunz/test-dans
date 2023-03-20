package handler

import "errors"

var (
	// ErrCtxTimeout - error returned when context is timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrRepoNil - error returned when repository is nil
	ErrRepoNil = errors.New("repository is nil")

	// ErrEmptyUsername - error returned when username is empty
	ErrEmptyUsername = errors.New("username is empty")

	// ErrEmptyPassword - error returned when password is empty
	ErrEmptyPassword = errors.New("password is empty")
)
