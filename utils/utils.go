package utils

import "github.com/google/uuid"

func Ptr[T any](v T) *T {
	return &v
}

func Uuid() string {
	// generate uuid
	return uuid.New().String()
}
