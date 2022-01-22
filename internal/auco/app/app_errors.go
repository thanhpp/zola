package app

import "errors"

var (
	ErrBlocked     = errors.New("blocked")
	ErrInvalidUser = errors.New("invalid user")
)
