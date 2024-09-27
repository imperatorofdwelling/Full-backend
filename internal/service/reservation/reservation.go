package reservation

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
)

type Service struct {
	Repo interfaces.ReservationRepo
}

func (s *Service) CreateReservation(ctx context.Context, reserv *reservation.ReservationEntity) error {
	const op = "service.reservation.CreateReservation"

	err := s.Repo.CreateReservation(ctx, reserv)
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

	isExists, err := s.Repo.CheckReservationIfExists(ctx, id)
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
