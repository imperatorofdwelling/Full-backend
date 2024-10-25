package searchhistory

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	history "github.com/imperatorofdwelling/Full-backend/internal/domain/models/searchhistory"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) AddHistory(ctx context.Context, userId, name string) error {
	const op = "repo.SearchHistory.AddHistory"

	id, _ := uuid.NewV4()
	createdAt := time.Now()

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO search_history (id, user_id, name, created_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, userId, name, createdAt)
	if err != nil {
		return fmt.Errorf("%s: executing statement: %w", op, err)
	}

	return nil
}

func (r *Repo) GetAllHistoryByUserId(ctx context.Context, userId string) ([]history.SearchHistory, error) {
	const op = "repo.SearchHistory.GetAllHistoryByUserId"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT s.user_id, s.name FROM search_history s WHERE s.user_id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: preparing statement: %w", op, err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: executing query: %w", op, err)
	}
	defer rows.Close()

	var searchHistory []history.SearchHistory

	for rows.Next() {
		var hist history.SearchHistory
		if err := rows.Scan(&hist.ID, &hist.Name); err != nil {
			return nil, fmt.Errorf("%s: scanning row: %w", op, err)
		}
		searchHistory = append(searchHistory, hist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows error: %w", op, err)
	}

	return searchHistory, nil
}
