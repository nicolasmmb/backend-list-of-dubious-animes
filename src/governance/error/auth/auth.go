package auth

import "errors"

var (
	ErrEmailIsRequired    = errors.New("email is required")
	ErrEmailIsInvalid     = errors.New("email is invalid")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserIsDeleted      = errors.New("user is deactivated")
)
