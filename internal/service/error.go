package service

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrValid             = errors.New("invalid data")

	ErrLocationNotFound = errors.New("location not found")

	ErrStayNotFound      = errors.New("stay not found")
	ErrStayImageNotFound = errors.New("stay image not found")

	ErrAdvantageNotFound = errors.New("advantage not found")

	ErrNotFoundReservation = errors.New("reservation not found")

	ErrStaysReviewNotFound = errors.New("stays review not found")
	ErrEmailAlreadyExists  = errors.New("email already exists")
	ErrNotFound            = errors.New("not found")
	ErrUpdateFailed        = errors.New("update failed")

	ErrAlreadyReserved      = errors.New("already reserved")
	ErrAlreadyReservedDate  = errors.New("already reserved date")
	ErrInvalidArrivalDate   = errors.New("invalid arrival date")
	ErrInvalidDepartureDate = errors.New("invalid departure date")
	ErrNoReservations       = errors.New("no reservations")
	ErrTimeNotCome          = errors.New("time not come")
	ErrReservationNotFound  = errors.New("reservation not found")
	ErrTimeHasNotCome       = errors.New("time has not come")

	ErrUserNotOwner = errors.New("user not owner")
)
