package service

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrNotFound           = errors.New("not found")
	ErrUpdateFailed       = errors.New("update failed")
	ErrValid              = errors.New("invalid data")
)
