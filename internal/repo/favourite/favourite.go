package favourite

import (
	"context"
	"database/sql"
	"fmt"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/favourite"
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

func (r *Repo) GetAllFavourites(ctx context.Context, userID string) ([]model.Favourite, error) {
	const op = "repo.Favourite.GetAllFavourites"

	stmt, err := r.Db.PrepareContext(ctx, `
		SELECT f.user_id, 
		       f.stay_id, 
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
		return nil, fmt.Errorf("%s: querying favourites: %w", op, err)
	}
	defer rows.Close()

	var favourites []model.Favourite

	for rows.Next() {
		var fav model.Favourite
		// Now scanning only the necessary fields
		if err := rows.Scan(&fav.UserID, &fav.StayID, &fav.City); err != nil {
			return nil, fmt.Errorf("%s: scanning row: %w", op, err)
		}
		favourites = append(favourites, fav)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows error: %w", op, err)
	}

	return favourites, nil
}
