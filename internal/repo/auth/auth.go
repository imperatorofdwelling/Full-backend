package auth

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
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

	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	currentTime := time.Now()
	rfc1123zTime := currentTime.Format(time.RFC1123Z)

	_, err = r.Db.ExecContext(ctx, query, id, user.Name, user.Email, user.Password, rfc1123zTime, rfc1123zTime)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *Repository) Login(ctx context.Context, user model.Login) (uuid.UUID, error) {
	const op = "repo.auth.Login"

	var storedPassword string
	var userID uuid.UUID

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, password FROM users WHERE email = $1")
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, user.Email).Scan(&userID, &storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, fmt.Errorf("%s: %w", op, repo.ErrUserNotFound)
		}
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	if user.IsHashed {
		if storedPassword == user.Password {
			return userID, nil
		}
		return uuid.Nil, fmt.Errorf("%s: %w", op, repo.ErrUserNotFound)
	}

	hashedPassword := sha256.Sum256([]byte(user.Password))
	hashedPasswordHex := hex.EncodeToString(hashedPassword[:])

	if hashedPasswordHex != storedPassword {
		return uuid.Nil, fmt.Errorf("%s: %w", op, repo.ErrUserNotFound)
	}

	return userID, nil
}

func (r *Repository) EmailVerification(ctx context.Context, userId string) error {
	const op = "repo.auth.EmailVerification"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users SET is_email_verified = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statement: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId)
	if err != nil {
		return fmt.Errorf("%s: failed to execute statement: %w", op, err)
	}

	return nil
}

func (r *Repository) CheckIfUserValidated(ctx context.Context, userId string) (bool, error) {
	const op = "repo.auth.CheckIfUserValidated"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT is_email_verified FROM users WHERE id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query %w", op, err)
	}
	defer stmt.Close()

	var isVerified bool
	err = stmt.QueryRowContext(ctx, userId).Scan(&isVerified)
	if err != nil {
		return false, fmt.Errorf("%s: failed to execute query %w", op, err)
	}

	return isVerified, nil
}
