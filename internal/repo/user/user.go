package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	const op = "repo.user.CheckUserIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return exists, nil
}

func (r *Repository) FindUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	const op = "repo.user.FindUserByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, name, email, phone, birth_date, national, gender, country, city, createdAt, updatedAt FROM users WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	var user user.User

	err = stmt.QueryRowContext(ctx, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.BirthDate,
		&user.National,
		&user.Gender,
		&user.Country,
		&user.City,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
