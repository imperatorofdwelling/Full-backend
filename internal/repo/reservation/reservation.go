package reservation

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateReservation(ctx context.Context, reservation *reservation.ReservationEntity) error {
	const op = "repo.reservation.CreateReservation"

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO reservations (stay_id, user_id, arrived, departure, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, reservation.StayID, reservation.UserID, reservation.Arrived, reservation.Departure, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) UpdateReservationByID(ctx context.Context, reservation *reservation.ReservationUpdateEntity) error {
	const op = "repo.reservation.UpdateReservationByID"

	exists, err := checkers.CheckReservationExists(ctx, r.Db, reservation.ID.String())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: reservation not found", op)
	}

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE reservations SET arrived = $1, departure = $2, updated_at = $3 WHERE id = $4")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, reservation.Arrived, reservation.Departure, time.Now(), reservation.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteReservationByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.reservation.deleteReservationByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM reservations WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) CheckReservationIfExists(ctx context.Context, id uuid.UUID) (bool, error) {
	const op = "repo.reservation.CheckReservationIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM reservations WHERE id = $1)")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return exists, nil
}
func (r *Repo) GetReservationByID(ctx context.Context, id uuid.UUID) (*reservation.Reservation, error) {
	const op = "repo.reservation.GetReservationByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM reservations WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var reserv reservation.Reservation

	err = stmt.QueryRowContext(ctx, id).Scan(&reserv.ID, &reserv.UserID, &reserv.StayID, &reserv.Arrived, &reserv.Departure, &reserv.UpdatedAt, &reserv.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no reservation found with id %v", op, id)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &reserv, nil
}

func (r *Repo) GetAllReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]reservation.Reservation, error) {
	const op = "repo.reservation.GetAllReservationsByUserID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM reservations WHERE user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var reservations []reservation.Reservation

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var reserv reservation.Reservation

		err = rows.Scan(&reserv.ID, &reserv.UserID, &reserv.StayID, &reserv.Arrived, &reserv.Departure, &reserv.UpdatedAt, &reserv.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		reservations = append(reservations, reserv)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &reservations, nil
}
