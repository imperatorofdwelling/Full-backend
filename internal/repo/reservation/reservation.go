package reservation

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"github.com/pkg/errors"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) ConfirmCheckOutReservation(ctx context.Context, userID string, stayID string) error {
	const op = "repo.reservation.ConfirmCheckOutReservation"

	query := `
		UPDATE reservations 
		SET check_out = TRUE 
		WHERE stay_id = $1 AND user_id = $2;
	`

	_, err := r.Db.ExecContext(ctx, query, stayID, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) CheckTimeForCheckOutReservation(ctx context.Context, userID string, stayID string) (bool, error) {
	const op = "repo.reservation.CheckTimeForCheckOutReservation"

	var departure time.Time

	query := `
       SELECT arrived 
       FROM reservations 
       WHERE user_id = $1 AND stay_id = $2 
       LIMIT 1;
    `

	err := r.Db.QueryRowContext(ctx, query, userID, stayID).Scan(&departure)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: reservation not found: %w", op, err)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	now := time.Now().Truncate(24 * time.Hour)
	departureDate := departure.Truncate(24 * time.Hour)

	result := departureDate.After(now)

	if result {
		return true, fmt.Errorf("%s: checkout time not yet reached. Current date: %v, Departure date: %v",
			op, now, departureDate)
	}

	return false, nil
}

func (r *Repo) CheckIfUserIsOwner(ctx context.Context, userID uuid.UUID, stayID uuid.UUID) (bool, error) {
	const op = "repo.reservation.CheckIfUserIsOwner"

	query := `SELECT COUNT(*) FROM stays WHERE id = $1 AND user_id = $2`

	var count int
	err := r.Db.QueryRowContext(ctx, query, stayID, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	if count == 0 {
		return false, fmt.Errorf("%s: %w", op, service.ErrUserNotOwner)
	}

	return true, nil
}

func (r *Repo) CheckReservationIsFree(ctx context.Context, reservationObject *reservation.ReservationEntity) error {
	const op = "repo.reservation.CheckReservationIsFree"

	query := `
        SELECT COUNT(*) 
        FROM reservations 
        WHERE stay_id = $1 
          AND (arrived, departure) OVERLAPS ($2, $3)
    `

	var count int
	err := r.Db.QueryRowContext(ctx, query, reservationObject.StayID, reservationObject.Arrived, reservationObject.Departure).Scan(&count)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if count > 0 {
		return fmt.Errorf("%s: %w", op, service.ErrAlreadyReservedDate)
	}

	return nil
}

func (r *Repo) CreateReservation(ctx context.Context, reservation *reservation.ReservationEntity, userID string) error {
	const op = "repo.reservation.CreateReservation"

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO reservations (stay_id, user_id, arrived, departure, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, reservation.StayID, userID, reservation.Arrived, reservation.Departure, time.Now(), time.Now())
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

func (r *Repo) CheckReservationIfExistsByUserId(ctx context.Context, id uuid.UUID) (bool, error) {
	const op = "repo.reservation.CheckReservationIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM reservations WHERE user_id = $1 and check_out = false)")
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

func (r *Repo) CheckIfReservationExists(ctx context.Context, id uuid.UUID) (bool, error) {
	const op = "repo.reservation.CheckIfReservationExists"

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

func (r *Repo) CheckIfArrivalIsCorrect(ctx context.Context, userID uuid.UUID, stayID uuid.UUID, reserv reservation.ReservationCheckInEntity) (bool, error) {
	const op = "repo.reservation.CheckIfArrivalIsCorrect"

	var arrived time.Time

	query := `
       SELECT arrived 
       FROM reservations 
       WHERE user_id = $1 AND stay_id = $2 
       LIMIT 1;
    `

	err := r.Db.QueryRowContext(ctx, query, reserv.UserID, stayID).Scan(&arrived)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	now := time.Now().Truncate(24 * time.Hour)
	arrivedDate := arrived.Truncate(24 * time.Hour)

	result := now.After(arrivedDate) || now.Equal(arrivedDate)

	if !result {
		return false, fmt.Errorf("%s: date check failed. Now: %v, Arrived: %v, Now(truncated): %v, Arrived(truncated): %v",
			op, time.Now(), arrived, now, arrivedDate)
	}

	return true, nil
}

func (r *Repo) CheckInApproval(ctx context.Context, stayID uuid.UUID, reservationUser reservation.ReservationCheckInEntity) error {
	const op = "repo.reservation.CheckInApproval"

	query := `
		UPDATE reservations 
		SET check_in = TRUE 
		WHERE stay_id = $1 AND user_id = $2;
	`

	result, err := r.Db.ExecContext(ctx, query, stayID, reservationUser.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, service.ErrReservationNotFound)
	}

	return nil
}

func (r *Repo) CheckIfReservationExistsByStayID(ctx context.Context, stayID uuid.UUID) (bool, error) {
	const op = "repo.reservation.CheckIfReservationExistsByStayID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM reservations WHERE stay_id = $1)")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, stayID).Scan(&exists)
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

func (r *Repo) GetFreeReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.Stay, error) {
	const op = "repo.reservation.GetFreeReservationsByUserID"

	query := `
		SELECT 
			id, user_id, location_id, name, type, number_of_bedrooms, 
			number_of_beds, number_of_bathrooms, guests, rating, 
			amenities, is_smoking_prohibited, square, street, house, 
			entrance, floor, room, price, created_at, updated_at
		FROM stays 
		WHERE user_id = $1 
		AND id NOT IN (SELECT stay_id FROM reservations WHERE check_out = false)`

	rows, err := r.Db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var reservations []stays.Stay

	for rows.Next() {
		var reserv stays.Stay
		var amenitiesJSON []byte

		err = rows.Scan(
			&reserv.ID, &reserv.UserID, &reserv.LocationID, &reserv.Name, &reserv.Type,
			&reserv.NumberOfBedrooms, &reserv.NumberOfBeds, &reserv.NumberOfBathrooms,
			&reserv.Guests, &reserv.Rating, &amenitiesJSON, &reserv.IsSmokingProhibited,
			&reserv.Square, &reserv.Street, &reserv.House, &reserv.Entrance,
			&reserv.Floor, &reserv.Room, &reserv.Price, &reserv.CreatedAt, &reserv.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		err = json.Unmarshal(amenitiesJSON, &reserv.Amenities)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to parse amenities JSON: %w", op, err)
		}

		reservations = append(reservations, reserv)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &reservations, nil
}

func (r *Repo) GetOccupiedReservationsByUserID(ctx context.Context, id uuid.UUID) (*[]stays.StayOccupied, error) {
	const op = "repo.reservation.GetOccupiedReservationsByUserID"

	query := `
		SELECT 
			s.id, s.user_id, s.location_id, s.name, s.type, s.number_of_bedrooms, 
			s.number_of_beds, s.number_of_bathrooms, s.guests, s.rating, 
			s.amenities, s.is_smoking_prohibited, s.square, s.street, s.house, 
			s.entrance, s.floor, s.room, s.price, s.created_at, s.updated_at,
			r.arrived, r.departure
		FROM stays s
		JOIN reservations r ON s.id = r.stay_id
		WHERE s.user_id = $1 
		AND r.check_out = false
	`

	rows, err := r.Db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var reservations []stays.StayOccupied

	for rows.Next() {
		var reserv stays.StayOccupied
		var amenitiesJSON []byte

		err = rows.Scan(
			&reserv.ID, &reserv.UserID, &reserv.LocationID, &reserv.Name, &reserv.Type,
			&reserv.NumberOfBedrooms, &reserv.NumberOfBeds, &reserv.NumberOfBathrooms,
			&reserv.Guests, &reserv.Rating, &amenitiesJSON, &reserv.IsSmokingProhibited,
			&reserv.Square, &reserv.Street, &reserv.House, &reserv.Entrance,
			&reserv.Floor, &reserv.Room, &reserv.Price, &reserv.CreatedAt, &reserv.UpdatedAt,
			&reserv.ArrivedAt, &reserv.DepartureAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		err = json.Unmarshal(amenitiesJSON, &reserv.Amenities)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to parse amenities JSON: %w", op, err)
		}

		reservations = append(reservations, reserv)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &reservations, nil
}
