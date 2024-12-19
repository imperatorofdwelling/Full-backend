package confirmEmail

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/pkg/otp"
	"github.com/pkg/errors"
	"time"
)

type Repo struct {
	DB *sql.DB
}

func (r *Repo) CreateOTP(ctx context.Context, userID string) error {
	const op = "repo.confirmEmail.CreateOTP"

	userOTP := otp.GenerateOTP()
	expireAt := time.Now().UTC().Add(5 * time.Minute)

	stmt, err := r.DB.PrepareContext(ctx, "INSERT INTO email_verifications(user_id, confirmation_code, expires_at) VALUES($1, $2, $3)")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query for inserting OTP: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, userOTP, expireAt)
	if err != nil {
		return fmt.Errorf("%s: failed to execute query for inserting OTP: %w", op, err)
	}

	return nil
}

func (r *Repo) GetOTP(ctx context.Context, userID string) (string, error) {
	const op = "repo.confirmEmail.GetOTP"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT confirmation_code FROM email_verifications WHERE user_id = $1")
	if err != nil {
		return "", fmt.Errorf("%s: failed to prepare query for getting OTP: %w", op, err)
	}
	defer stmt.Close()

	var code string
	err = stmt.QueryRowContext(ctx, userID).Scan(&code)
	if err != nil {
		return "", fmt.Errorf("%s: failed to execute query for getting OTP: %w", op, err)
	}

	return code, nil
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

	stmt, err := r.DB.PrepareContext(ctx, "SELECT expires_at FROM email_verifications where user_id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var expiresAt time.Time
	err = stmt.QueryRowContext(ctx, userID).Scan(&expiresAt)
	if err != nil {
		return false, fmt.Errorf("%s: failed to query expires_at: %w", op, err)
	}

	// For checking purposes
	// fmt.Println(time.Now().UTC())
	// fmt.Println(expiresAt.UTC())
	// fmt.Println(time.Now().UTC().After(expiresAt.UTC()))

	return time.Now().UTC().After(expiresAt.UTC()), nil
}

func (r *Repo) UpdateOTP(ctx context.Context, userID string) error {
	const op = "repo.confirmEmail.UpdateOTP"

	newOTP := otp.GenerateOTP()
	newExpiresAt := time.Now().UTC().Add(5 * time.Minute)

	fmt.Println("Я тут")

	stmt, err := r.DB.PrepareContext(ctx, "UPDATE email_verifications SET expires_at = $1, confirmation_code = $2 WHERE user_id = $3")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newExpiresAt, newOTP, userID)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update: %w", op, err)
	}

	return nil
}
