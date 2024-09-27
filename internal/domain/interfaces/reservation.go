package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"net/http"
)

//go:generate mockery --name ReservationRepo
type ReservationRepo interface {
	CreateReservation(context.Context, *reservation.ReservationEntity) error
	UpdateReservationByID(context.Context, *reservation.ReservationUpdateEntity) error
	DeleteReservationByID(context.Context, uuid.UUID) error
	CheckReservationIfExists(context.Context, uuid.UUID) (bool, error)
	GetReservationByID(context.Context, uuid.UUID) (*reservation.Reservation, error)
	GetAllReservationsByUserID(context.Context, uuid.UUID) (*[]reservation.Reservation, error)
}

//go:generate mockery --name ReservationService
type ReservationService interface {
	CreateReservation(context.Context, *reservation.ReservationEntity) error
	UpdateReservation(context.Context, *reservation.ReservationUpdateEntity) error
	DeleteReservationByID(context.Context, uuid.UUID) error
	GetReservationByID(context.Context, uuid.UUID) (*reservation.Reservation, error)
	GetAllReservationsByUser(context.Context, uuid.UUID) (*[]reservation.Reservation, error)
}

type ReservationHandler interface {
	CreateReservation(http.ResponseWriter, *http.Request)
	UpdateReservation(http.ResponseWriter, *http.Request)
	DeleteReservationByID(http.ResponseWriter, *http.Request)
	GetReservationByID(http.ResponseWriter, *http.Request)
	GetAllReservationsByUser(http.ResponseWriter, *http.Request)
}
