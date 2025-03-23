package reservation

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"time"
)

type Service struct {
	Repo interfaces.ReservationRepo
}

func (s *Service) ConfirmCheckOutReservation(ctx context.Context, userID string, stayID string) error {
	const op = "service.reservation.ConfirmCheckOutReservation"

	timeExist, err := s.Repo.CheckTimeForCheckOutReservation(ctx, userID, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if timeExist {
		return fmt.Errorf("%s: %w", op, service.ErrTimeHasNotCome)
	}

	err = s.Repo.ConfirmCheckOutReservation(ctx, userID, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) ConfirmCheckInReservation(ctx context.Context, userID string, stayID string, reserv reservation.ReservationCheckInEntity) error {
	const op = "service.reservation.ConfirmCheckInReservation"

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	stayUUID, err := uuid.FromString(stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	owner, err := s.Repo.CheckIfUserIsOwner(ctx, userUUID, stayUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if !owner {
		return fmt.Errorf("%s: %w", op, service.ErrUserNotOwner)
	}

	exists, err := s.Repo.CheckIfReservationExistsByStayID(ctx, stayUUID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if !exists {
		return fmt.Errorf("%s: %w", op, service.ErrNoReservations)
	}

	dateErr, err := s.Repo.CheckIfArrivalIsCorrect(ctx, userUUID, stayUUID, reserv)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if !dateErr {
		return fmt.Errorf("%s: %w", op, service.ErrTimeNotCome)
	}

	err = s.Repo.CheckInApproval(ctx, stayUUID, reserv)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) CheckReservation(ctx context.Context, reservationObj *reservation.ReservationEntity, userID string) error {
	const op = "service.reservation.CheckReservation"

	now := time.Now().Truncate(24 * time.Hour)

	if reservationObj.Arrived.Before(now) {
		return fmt.Errorf("%s: %w", op, service.ErrInvalidArrivalDate)
	}

	if !reservationObj.Departure.After(reservationObj.Arrived.Add(24*time.Hour - time.Nanosecond)) {
		return fmt.Errorf("%s: %w", op, service.ErrInvalidDepartureDate)
	}

	uuidUserId := uuid.FromStringOrNil(userID)

	exists, err := s.Repo.CheckReservationIfExistsByUserId(ctx, uuidUserId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if exists {
		return fmt.Errorf("%s: %w", op, service.ErrAlreadyReserved)
	}

	err = s.Repo.CheckReservationIsFree(ctx, reservationObj)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) CreateReservation(ctx context.Context, reserv *reservation.ReservationEntity, userID string) error {
	const op = "service.reservation.CreateReservation"

	err := s.Repo.CreateReservation(ctx, reserv, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateReservation(ctx context.Context, reserv *reservation.ReservationUpdateEntity) error {
	const op = "service.reservation.UpdateReservation"

	foundReserv, err := s.Repo.GetReservationByID(ctx, reserv.ID)
	if err != nil {
		return err
	}

	if foundReserv.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrNotFoundReservation)
	}

	err = s.Repo.UpdateReservationByID(ctx, reserv)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteReservationByID(ctx context.Context, id uuid.UUID) error {
	const op = "service.reservation.DeleteReservationByID"

	isExists, err := s.Repo.CheckIfReservationExists(ctx, id)
	if err != nil {
		return err
	}

	if !isExists {
		return fmt.Errorf("%s: %w", op, service.ErrNotFoundReservation)
	}

	err = s.Repo.DeleteReservationByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetReservationByID(ctx context.Context, id uuid.UUID) (*reservation.Reservation, error) {
	const op = "service.reservation.GetReservationByID"

	foundReserv, err := s.Repo.GetReservationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if foundReserv.ID == uuid.Nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrNotFoundReservation)
	}

	return foundReserv, nil
}

func (s *Service) GetAllReservationsByUser(ctx context.Context, id uuid.UUID) (*[]reservation.Reservation, error) {
	const op = "service.reservation.GetAllReservationsByUser"

	// TODO Check user if exists

	reserv, err := s.Repo.GetAllReservationsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return reserv, nil
}

func (s *Service) GetFreeReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.Stay, error) {
	const op = "service.reservation.GetFreeReservationsByUserID"

	reserv, err := s.Repo.GetFreeReservationsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return reserv, nil
}

func (s *Service) GetOccupiedReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.Stay, error) {
	const op = "service.reservation.GetOccupiedReservationsByUserID"

	reserv, err := s.Repo.GetOccupiedReservationsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	return reserv, nil
}
