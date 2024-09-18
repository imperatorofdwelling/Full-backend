package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models/user"
)

type UserRepository struct {
	Db *sql.DB
}

func (r *UserRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	const op = "repo.user.CheckUserIfExists"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)")
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

func (r *UserRepository) CreateUser(ctx context.Context, user *user.Registration) (uuid.UUID, error) {
	const op = "repo.user.CreateUser"

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

func (r *UserRepository) Login(ctx context.Context, user *user.Login) (int64, error) {
	const op = "repo.user.Login"

}

func (r *UserRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	const op = "repo.user.FindUserByID"

	stmt, err := r.Db.PrepareContext(ctx, "SELECT * FROM users WHERE id = ?")
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
