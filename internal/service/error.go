package service

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrValid             = errors.New("invalid data")

	ErrLocationNotFound = errors.New("location not found")

	ErrStayNotFound = errors.New("stay not found")

	ErrAdvantageNotFound = errors.New("advantage not found")
)
