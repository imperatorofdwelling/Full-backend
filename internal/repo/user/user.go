package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/repo"
	"time"
)

type Repository struct {
	Db *sql.DB
}

func (r *Repository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	const op = "repo.user.CheckUserIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, repo.ErrNotFound)
	}

	defer stmt.Close()

	var exists bool

	err = stmt.QueryRowContext(ctx, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return exists, nil
}

func (r *Repository) FindUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	const op = "repo.user.FindUserByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, name, email, phone, avatar, birth_date, national, gender, country, city, createdAt, updatedAt FROM users WHERE id = $1")
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, repo.ErrNotFound)
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

func (r *Repository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	const op = "repo.user.DeleteUserByID"

	stmt, err := r.Db.PrepareContext(ctx, "DELETE FROM users WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", op, repo.ErrNotFound)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
