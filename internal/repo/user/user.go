package user

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/repo"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
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

	stmt, err := r.Db.PrepareContext(ctx, "SELECT id, name, email, phone, avatar, birth_date, national, gender, country, city, role_id, created_at, updated_at FROM users WHERE id = $1")
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
		&user.RoleID,
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
			name = COALESCE($2, name), 
			avatar = COALESCE($3, avatar), 
			birth_date = COALESCE($4, birth_date), 
			national = COALESCE($5, national), 
			gender = COALESCE($6, gender), 
			country = COALESCE($7, country), 
			city = COALESCE($8, city), 
			updated_at = $9
		WHERE id = $1
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, repo.ErrUpdateFailed)
	}
	defer stmt.Close()

	currentTime := time.Now()

	result, err := stmt.ExecContext(ctx,
		id,
		user.Name,
		string(user.Avatar),
		user.BirthDate,
		user.National,
		user.Gender,
		user.Country,
		user.City,
		currentTime.Format(time.RFC1123Z),
	)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, "Error: no rows affected")
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

func (r *Repository) UpdateUserEmailByID(ctx context.Context, id uuid.UUID, newEmail string) error {
	const op = "repo.user.UpdateUserEmailByID"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users SET email = $1, is_email_verified = false WHERE id = $2")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newEmail, id)
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

func (r *Repository) CreateUserPfp(ctx context.Context, userId, imagePath string) error {
	const op = "repo.user.CreateUserPfp"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users SET avatar = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, imagePath, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Repository) GetUserPfp(ctx context.Context, userId string) (string, error) {
	const op = "repo.user.GetUserPfp"

	var avatarPath sql.NullString
	query := "SELECT avatar FROM users WHERE id = $1"

	err := r.Db.QueryRowContext(ctx, query, userId).Scan(&avatarPath)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("%s: failed to query avatar path: %w", op, err)
	}

	if !avatarPath.Valid {
		return "", nil
	}

	return avatarPath.String, nil
}

func (r *Repository) DeleteUserPfp(ctx context.Context, userId uuid.UUID) error {
	const op = "repo.user.DeleteUserPfp"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users SET avatar = NULL WHERE id = $1")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userId)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update query: %w", op, err)
	}

	return nil
}

func (r *Repository) UpdateUserPfp(ctx context.Context, userId uuid.UUID, imagePath string) error {
	const op = "repo.user.UpdateUserPfp"

	stmt, err := r.Db.PrepareContext(ctx, "UPDATE users SET avatar = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, imagePath, userId)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update query: %w", op, err)
	}
	return nil
}
