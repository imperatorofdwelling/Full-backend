package favourite

import (
	"context"
	"database/sql"
	"fmt"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) AddFavourite(ctx context.Context, userId, stayID string) error {
	const op = "repo.Favourite.AddFavourite"

	exists, err := r.CheckStayExists(ctx, stayID)
	if err != nil {
		return fmt.Errorf("%s: checking stay existence: %w", op, err)
	}
	if !exists {
		return fmt.Errorf("%s: stay does not exist: %s", op, stayID)
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

func (r *Repo) CheckStayExists(ctx context.Context, stayID string) (bool, error) {
	var exists bool
	err := r.Db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM stays WHERE id = $1)", stayID).Scan(&exists)
	return exists, err
}
