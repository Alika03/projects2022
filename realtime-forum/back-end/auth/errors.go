package auth

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrUnauthorized    = errors.New("unauthorized")
)
