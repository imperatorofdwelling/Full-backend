package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/repo"
	"time"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) Register(ctx context.Context, user model.Registration) (uuid.UUID, error) {
	// Create a new UUID for the user
	id, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}

	query := `INSERT INTO users (id, name, email, password, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6)`

	currentTime := time.Now()
	rfc1123zTime := currentTime.Format(time.RFC1123Z)

	_, err = r.Db.ExecContext(ctx, query, id, user.Name, user.Email, user.Password, rfc1123zTime, rfc1123zTime)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *Repository) Login(ctx context.Context, user model.Login) (uuid.UUID, error) {
	const op = "repo.user.Login"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id FROM users WHERE email = $1 AND password = $2")
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, repo.ErrNotFound)
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
