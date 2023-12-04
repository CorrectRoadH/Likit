package domain

import "errors"

var (
	ErrBusinessNotExist = errors.New("business is not exist")

	ErrDatabaseTypeNotSupport = errors.New("database type not support")
)
