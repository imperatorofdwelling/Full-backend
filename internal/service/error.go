package service

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrValid             = errors.New("invalid data")

	ErrLocationNotFound = errors.New("location not found")

	ErrStayNotFound = errors.New("stay not found")

	ErrAdvantageNotFound = errors.New("advantage not found")

	ErrNotFoundReservation = errors.New("reservation not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrNotFound           = errors.New("not found")
	ErrUpdateFailed       = errors.New("update failed")
)
