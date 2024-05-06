package user

import "errors"

var (
	ErrNameIsRequired     = errors.New("name is required")
	ErrEmailIsRequired    = errors.New("email is required")
	ErrEmailIsInvalid     = errors.New("email is invalid")
	ErrPasswordIsRequired = errors.New("password is required")
	ErrIDIsRequired       = errors.New("id is required")
)
