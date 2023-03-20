package mysql

import "errors"

var (
	// ErrCtxTimeout - error returned when context is timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrDBConnNil - error returned when mysql connection is nil
	ErrDBConnNil = errors.New("database connection is nil")

	// ErrUsernameEmpty - error returned when username is empty
	ErrUsernameEmpty = errors.New("username is empty")

	// ErrUsernameEmpty - error returned when table name is empty
	ErrTableNameEmpty = errors.New("table name is empty")

	// ErrEmptyPassword - error returned when password is empty
	ErrEmptyPassword = errors.New("password is empty")
)
