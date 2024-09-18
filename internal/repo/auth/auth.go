package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) Registration(ctx context.Context, user *auth.Registration) (uuid.UUID, error) {
	const op = "repo.user.Registration"

	stmt, err := r.Db.PrepareContext(ctx,
		"INSERT INTO users (id, name, email, password, phone, birth_date, national, gender, country, city) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	userID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(
		ctx,
		userID,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}

func (r *Repository) Login(ctx context.Context, user *auth.Login) (uuid.UUID, error) {
	const op = "repo.user.Login"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id FROM users WHERE email = ? AND password = ?")
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	defer stmt.Close()

	userID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, user.Email, user.Password).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, fmt.Errorf("%s: %w", op, err)
		}
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userID, nil
}
