package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/repo"
	_ "github.com/lib/pq"
	"time"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	const op = "repo.user.CheckUserExists"

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

func (r *Repository) GetUserIDByEmail(ctx context.Context, email string) (string, error) {
	const op = "repo.user.GetUserIDByEmail"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id FROM users WHERE email = $1")
	if err != nil {
		return "", fmt.Errorf("%s: Error while prepating query %w", op, err)
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRowContext(ctx, email).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("%s: Error while executing query %w", op, err)
	}

	return id, nil
}

func (r *Repository) GetUserPasswordByEmail(ctx context.Context, email string) (string, error) {
	const op = "repo.user.GetUserPasswordByEmail"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT password FROM users WHERE email = $1")
	if err != nil {
		return "", fmt.Errorf("%s: Error while prepating query %w", op, err)
	}
	defer stmt.Close()

	var password string
	err = stmt.QueryRowContext(ctx, email).Scan(&password)
	if err != nil {
		return "", fmt.Errorf("%s: Error while executing query %w", op, err)
	}

	return password, nil
}

func (r *Repository) FindUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	const op = "repo.user.FindUserByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, name, email, phone, avatar, birth_date, national, gender, country, city, created_at, updated_at FROM users WHERE id = $1")
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, repo.ErrUserNotFound)
	}

	defer stmt.Close()

	var user model.User

	err = stmt.QueryRowContext(ctx, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Avatar,
		&user.BirthDate,
		&user.National,
		&user.Gender,
		&user.Country,
		&user.City,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (r *Repository) UpdateUserByID(ctx context.Context, id uuid.UUID, user model.User) error {
	const op = "repo.user.UpdateUser"

	stmt, err := r.Db.PrepareContext(ctx, `
		UPDATE users 
		SET 
			name = $2, 
			email = $3, 
			phone = $4, 
			avatar = $5, 
			birth_date = $6, 
			national = $7, 
			gender = $8, 
			country = $9, 
			city = $10, 
			updatedAt = $11
		WHERE id = $1
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, repo.ErrUpdateFailed)
	}

	defer stmt.Close()
	currentTime := time.Now()
	rfc1123zTime := currentTime.Format(time.RFC1123Z)

	result, err := stmt.ExecContext(ctx,
		id,
		user.Name,
		user.Email,
		user.Phone,
		user.Avatar,
		user.BirthDate,
		user.National,
		user.Gender,
		user.Country,
		user.City,
		rfc1123zTime,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repository) UpdateUserPasswordByID(ctx context.Context, id uuid.UUID, newPassword string) error {
	const op = "repo.user.UpdateUserPasswordByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users SET password = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newPassword, id.String())
	if err != nil {
		return fmt.Errorf("%s: failed to execute update query: %w", op, err)
	}

	return nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.user.DeleteUserByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM users WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, repo.ErrUserNotFound)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
