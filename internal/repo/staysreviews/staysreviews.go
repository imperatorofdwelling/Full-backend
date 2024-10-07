package staysreviews

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreviews"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateStaysReview(ctx context.Context, stayReview *staysreviews.StaysReviewEntity) error {
	const op = "repo.staysreviews.CreateStaysReview"

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO stays_reviews (stay_id, user_id, title, description, rating, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, stayReview.StayID, stayReview.UserID, stayReview.Title, stayReview.Description, stayReview.Rating, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) UpdateStaysReviewByID(ctx context.Context, stayReview *staysreviews.StaysReviewEntity, id uuid.UUID) error {
	const op = "repo.staysreviews.UpdateStaysReviewByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE stays_reviews SET title = $1, description = $2, rating = $3, updated_at = $4 WHERE id = $5")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, stayReview.Title, stayReview.Description, stayReview.Rating, time.Now(), id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteStaysReviewByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.staysreviews.DeleteStaysReviewByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM stays_reviews WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) FindOneStaysReviewByID(ctx context.Context, id uuid.UUID) (*staysreviews.StaysReview, error) {
	const op = "repo.staysreviews.FindOneStaysReviewByID"

	var stayReview staysreviews.StaysReview

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays_reviews WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, id).Scan(&stayReview.ID, &stayReview.StayID, &stayReview.UserID, &stayReview.Title, &stayReview.Description, &stayReview.Rating, &stayReview.CreatedAt, &stayReview.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &stayReview, nil
}

func (r *Repo) FindAllStaysReviews(ctx context.Context) ([]staysreviews.StaysReview, error) {
	const op = "repo.staysreviews.FindAllStaysReviews"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM stays_reviews")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var stayReviews []staysreviews.StaysReview

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var stayReview staysreviews.StaysReview

		err := rows.Scan(&stayReview.ID, &stayReview.StayID, &stayReview.UserID, &stayReview.Title, &stayReview.Description, &stayReview.Rating, &stayReview.CreatedAt, &stayReview.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		stayReviews = append(stayReviews, stayReview)
	}

	return stayReviews, nil
}

func (r *Repo) CheckIfExists(ctx context.Context, id uuid.UUID) (bool, error) {
	const op = "repo.staysreviews.CheckIfExists"

	stmt, err := r.Db.PrepareContext(ctx, `SELECT EXISTS (SELECT 1 FROM stays_reviews WHERE id = $1)`)
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
