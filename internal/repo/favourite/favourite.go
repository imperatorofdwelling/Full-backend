package favourite

import (
	"context"
	"database/sql"
	"fmt"
	stays2 "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) AddFavourite(ctx context.Context, userId, stayID string) error {
	const op = "repo.Favourite.AddFavourite"

	// Check if the stay exists
	exists, err := checkers.CheckStayExists(ctx, r.Db, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: %w", op, err)
	}

	// Check if the favourite already exists
	favExists, err := checkers.CheckFavouriteExists(ctx, r.Db, userId, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if favExists {
		return fmt.Errorf("%s: favourite already exists for user: %s and stay: %s", op, userId, stayID)
	}

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO favourite (user_id, stay_id, created_at) VALUES ($1, $2, CURRENT_TIMESTAMP) ON CONFLICT (user_id, stay_id) DO NOTHING")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) RemoveFavourite(ctx context.Context, userId, stayID string) error {
	const op = "repo.Favourite.RemoveFavourite"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM favourite WHERE user_id = $1 AND stay_id = $2")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) GetAllFavourites(ctx context.Context, userID string) ([]stays2.StayEntityFav, error) {
	const op = "repo.Favourite.GetAllFavouriteStays"

	stmt, err := r.Db.PrepareContext(ctx, `
		SELECT 
		    s.id,
		    f.user_id,
		    s.location_id,
		    s.name,
		    s.type,
		    s.number_of_bedrooms,
		    s.number_of_beds,
		    s.number_of_bathrooms,
		    s.guests,
		    s.is_smoking_prohibited,
		    s.square,
		    s.street,
		    s.house,
		    s.entrance,
		    s.floor,
		    s.room,
		    s.price,
		    l.city
		FROM favourite f
		JOIN stays s ON f.stay_id = s.id
		JOIN locations l ON s.location_id = l.id
		WHERE f.user_id = $1
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: querying favourite stays: %w", op, err)
	}
	defer rows.Close()

	var stays []stays2.StayEntityFav

	for rows.Next() {
		var stay stays2.StayEntityFav
		if err := rows.Scan(
			&stay.ID,
			&stay.UserID,
			&stay.LocationID,
			&stay.Name,
			&stay.Type,
			&stay.NumberOfBedrooms,
			&stay.NumberOfBeds,
			&stay.NumberOfBathrooms,
			&stay.Guests,
			&stay.IsSmokingProhibited,
			&stay.Square,
			&stay.Street,
			&stay.House,
			&stay.Entrance,
			&stay.Floor,
			&stay.Room,
			&stay.Price,
			&stay.City,
		); err != nil {
			return nil, fmt.Errorf("%s: scanning row: %w", op, err)
		}
		stays = append(stays, stay)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows error: %w", op, err)
	}

	return stays, nil
}
