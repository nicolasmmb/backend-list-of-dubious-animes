package auth

import "errors"

var (
	ErrEmailIsRequired    = errors.New("email is required")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserIsDeleted      = errors.New("user is deactivated")
)
