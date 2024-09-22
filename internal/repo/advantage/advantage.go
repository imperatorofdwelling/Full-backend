package advantage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateAdvantage(ctx context.Context, advTitle string, imgPath string) error {
	const op = "repo.advantage.CreateAdvantage"

	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO advantages (id, title, image) VALUES ($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id, advTitle, imgPath)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) RemoveAdvantage(ctx context.Context, id uuid.UUID) error {
	const op = "repo.advantage.RemoveAdvantage"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM advantages WHERE id = $1")
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

func (r *Repo) CheckAdvantageIfExists(ctx context.Context, advName string) (bool, error) {
	const op = "repo.advantage.CheckAdvantageIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 from advantages WHERE title=$1)")
	if err != nil {
		return false, fmt.Errorf("%s: %v", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRow(advName).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		} else {
			return false, fmt.Errorf("%s: %v", op, err)
		}
	}

	return exists, nil
}

func (r *Repo) FindAdvantageByID(ctx context.Context, id uuid.UUID) (*models.Advantage, error) {
	const op = "repo.advantage.FindAdvantageByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM advantages WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var advantage models.Advantage

	err = row.Scan(&advantage.ID, &advantage.Title, &advantage.Image, &advantage.CreatedAt, &advantage.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &advantage, nil
}

func (r *Repo) GetAllAdvantages(ctx context.Context) ([]models.Advantage, error) {
	const op = "repo.advantage.GetAllAdvantages"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM advantages")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var advantages []models.Advantage

	for rows.Next() {
		var advantage models.Advantage

		if err = rows.Scan(&advantage.ID, &advantage.Title, &advantage.Image, &advantage.CreatedAt, &advantage.UpdatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		advantages = append(advantages, advantage)
	}

	return advantages, nil
}

func (r *Repo) UpdateAdvantageByID(ctx context.Context, id uuid.UUID, adv *models.Advantage) error {
	const op = "repo.advantage.UpdateAdvantageByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE advantages SET title = $1, image = $2, updated_at = $3 WHERE id = $4")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, adv.Title, adv.Image, time.Now(), id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
