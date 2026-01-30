package services

import (
	"errors"
)

var (
	ErrValidation = errors.New("validation error")
	ErrUserExists = errors.New("user with email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrInternal = errors.New("internal server error")
	ErrPasswordTooLong = errors.New("password is too long")
)

