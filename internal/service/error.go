package service

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")

	ErrLocationNotFound = errors.New("location not found")

	ErrStayNotFound = errors.New("stay not found")
)
