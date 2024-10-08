package staysadvantage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysadvantage"
	"time"
)

type Repo struct {
	Db *sql.DB
}

func (r *Repo) CreateStaysAdvantage(ctx context.Context, stayAdv *models.StayAdvantageEntity) error {
	const op = "repo.staysadvantage.CreateStaysAdvantage"

	stmt, err := r.Db.PrepareContext(ctx, "INSERT INTO stays_advantages(stay_id, advantage_id, title, image, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, stayAdv.StayID, stayAdv.AdvantageID, stayAdv.Title, stayAdv.Image, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repo) DeleteStaysAdvantageByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.staysadvantage.DeleteStaysAdvantageByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM stays_advantages WHERE stay_id = $1")
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

func (r *Repo) CheckStaysAdvantageIfExists(ctx context.Context, id uuid.UUID) (bool, error) {
	const op = "repo.staysadvantage.CheckStaysAdvantageIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM stays_advantages WHERE id = $1)")
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
