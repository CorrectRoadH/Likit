package domain

import "errors"

var (
	ErrBusinessNotExist = errors.New("business is not exist")

	ErrDatabaseTypeNotSupport = errors.New("database type not support")

	ErrUserHasVoted = errors.New("user has voted")

	ErrUserNotVoted = errors.New("user has not voted")
)
