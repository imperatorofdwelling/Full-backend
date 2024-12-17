package confirmEmail

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type Repo struct {
	DB *sql.DB
}

func (r *Repo) CreateOTP(ctx context.Context, userID string) error {
	const op = "repo.confirmEmail.CreateOTP"

	exist, err := r.CheckOTPExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		notExpired, err := r.CheckOTPNotExpired(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if notExpired {
			return fmt.Errorf("%s : OTP already exists and is not expired", op)
		}

		err = r.UpdateOTP(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to update expired OTP: %w", op, err)
		}
	}

	return nil
}

func (r *Repo) CheckOTPExists(ctx context.Context, userID string) (bool, error) {
	const op = "repo.confirmEmail.CheckOTPExists"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT id FROM email_verifications WHERE user_id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var userId string
	err = stmt.QueryRowContext(ctx, userID).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return true, nil
}

func (r *Repo) CheckOTPNotExpired(ctx context.Context, userID string) (bool, error) {
	const op = "repo.confirmEmail.CheckOTPNotExpired"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT expires_at FROM email_verifications where id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var expiresAt time.Time
	err = stmt.QueryRowContext(ctx, userID).Scan(&expiresAt)
	if err != nil {
		return false, fmt.Errorf("%s: failed to query expires_at: %w", op, err)
	}

	if time.Now().After(expiresAt) {
		return false, nil
	}

	return true, nil
}

func (r *Repo) UpdateOTP(ctx context.Context, userID string) error {
	const op = "repo.confirmEmail.UpdateOTP"

	newExpiresAt := time.Now().Add(5 * time.Minute)

	stmt, err := r.DB.PrepareContext(ctx, "UPDATE email_verifications SET expires_at = $1 WHERE user_id = $2")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newExpiresAt, userID)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update: %w", op, err)
	}

	return nil
}
