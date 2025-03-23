package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"net/http"
)

//go:generate mockery --name ReservationRepo
type ReservationRepo interface {
	CheckReservationIsFree(context.Context, *reservation.ReservationEntity) error
	CreateReservation(context.Context, *reservation.ReservationEntity, string) error
	UpdateReservationByID(context.Context, *reservation.ReservationUpdateEntity) error
	DeleteReservationByID(context.Context, uuid.UUID) error
	CheckReservationIfExistsByUserId(context.Context, uuid.UUID) (bool, error)
	CheckIfArrivalIsCorrect(context.Context, uuid.UUID, uuid.UUID, reservation.ReservationCheckInEntity) (bool, error)
	CheckInApproval(context.Context, uuid.UUID, reservation.ReservationCheckInEntity) error
	CheckIfReservationExists(context.Context, uuid.UUID) (bool, error)
	CheckIfReservationExistsByStayID(context.Context, uuid.UUID) (bool, error)
	GetReservationByID(context.Context, uuid.UUID) (*reservation.Reservation, error)
	GetAllReservationsByUserID(context.Context, uuid.UUID) (*[]reservation.Reservation, error)
	CheckIfUserIsOwner(context.Context, uuid.UUID, uuid.UUID) (bool, error)
	ConfirmCheckOutReservation(context.Context, string, string) error
	CheckTimeForCheckOutReservation(context.Context, string, string) (bool, error)
	GetFreeReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.Stay, error)
	GetOccupiedReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.StayOccupied, error)
}

//go:generate mockery --name ReservationService
type ReservationService interface {
	CheckReservation(context.Context, *reservation.ReservationEntity, string) error
	CreateReservation(context.Context, *reservation.ReservationEntity, string) error
	UpdateReservation(context.Context, *reservation.ReservationUpdateEntity) error
	DeleteReservationByID(context.Context, uuid.UUID) error
	GetReservationByID(context.Context, uuid.UUID) (*reservation.Reservation, error)
	GetAllReservationsByUser(context.Context, uuid.UUID) (*[]reservation.Reservation, error)
	ConfirmCheckInReservation(context.Context, string, string, reservation.ReservationCheckInEntity) error
	ConfirmCheckOutReservation(context.Context, string, string) error
	GetFreeReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.Stay, error)
	GetOccupiedReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.StayOccupied, error)
}

type ReservationHandler interface {
	CreateReservation(http.ResponseWriter, *http.Request)
	UpdateReservation(http.ResponseWriter, *http.Request)
	DeleteReservationByID(http.ResponseWriter, *http.Request)
	GetReservationByID(http.ResponseWriter, *http.Request)
	GetAllReservationsByUser(http.ResponseWriter, *http.Request)
	ConfirmCheckInReservation(http.ResponseWriter, *http.Request)
	ConfirmCheckOutReservation(http.ResponseWriter, *http.Request)
	GetFreeReservationsByUserID(http.ResponseWriter, *http.Request)
	GetOccupiedReservationsByUserID(http.ResponseWriter, *http.Request)
}
