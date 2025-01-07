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

func (r *Repo) CreateEmailOTP(ctx context.Context, userID string) (string, error) {
	const op = "repo.confirmEmail.CreateOTP"

	userOTP := otp.GenerateOTP()
	expireAt := time.Now().UTC().Add(5 * time.Minute)

	stmt, err := r.DB.PrepareContext(ctx, "INSERT INTO email_verifications(user_id, confirmation_code, expires_at) VALUES($1, $2, $3)")
	if err != nil {
		return "", fmt.Errorf("%s: failed to prepare query for inserting OTP: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, userOTP, expireAt)
	if err != nil {
		return "", fmt.Errorf("%s: failed to execute query for inserting OTP: %w", op, err)
	}

	return userOTP, nil
}

func (r *Repo) CreatePasswordOTP(ctx context.Context, email string) (string, error) {
	const op = "repo.confirmEmail.CreatePasswordOTP"

	userOTP := otp.GenerateOTP()
	expireAt := time.Now().UTC().Add(2 * time.Minute)

	stmt, err := r.DB.PrepareContext(ctx, "INSERT INTO password_verifications(email, confirmation_code, expires_at) VALUES($1, $2, $3)")
	if err != nil {
		return "", fmt.Errorf("%s: failed to prepare query for inserting OTP: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, email, userOTP, expireAt)
	if err != nil {
		return "", fmt.Errorf("%s: failed to execute query for inserting OTP: %w", op, err)
	}

	return userOTP, nil
}

func (r *Repo) GetEmailOTP(ctx context.Context, userID string) (string, error) {
	const op = "repo.confirmEmail.GetEmailOTP"

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

func (r *Repo) GetPasswordOTP(ctx context.Context, email string) (string, error) {
	const op = "repo.confirmEmail.GetPasswordOTP"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT confirmation_code FROM password_verifications WHERE email = $1")
	if err != nil {
		return "", fmt.Errorf("%s: failed to prepare query for getting OTP: %w", op, err)
	}
	defer stmt.Close()

	var code string
	err = stmt.QueryRowContext(ctx, email).Scan(&code)
	if err != nil {
		return "", fmt.Errorf("%s: failed to execute query for getting OTP: %w", op, err)
	}

	return code, nil
}

func (r *Repo) CheckEmailOTPExists(ctx context.Context, userID string) (bool, error) {
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

func (r *Repo) CheckEmailOTPNotExpired(ctx context.Context, userID string) (bool, error) {
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

func (r *Repo) CheckPasswordOTPExists(ctx context.Context, email string) (bool, error) {
	const op = "repo.confirmEmail.CheckPasswordOTPExists"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT id FROM password_verifications WHERE email = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var userId string
	err = stmt.QueryRowContext(ctx, email).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return true, nil
}

func (r *Repo) CheckPasswordOTPNotExpired(ctx context.Context, email string) (bool, error) {
	const op = "repo.confirmEmail.CheckPasswordOTPNotExpired"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT expires_at FROM password_verifications where email = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var expiresAt time.Time
	err = stmt.QueryRowContext(ctx, email).Scan(&expiresAt)
	if err != nil {
		return false, fmt.Errorf("%s: failed to query expires_at: %w", op, err)
	}

	return time.Now().UTC().After(expiresAt.UTC()), nil
}

func (r *Repo) CheckPasswordOTPVerified(ctx context.Context, email string) (bool, error) {
	const op = "repo.confirmEmail.CheckPasswordOTPVerified"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT is_verified FROM password_verifications where email = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var verified bool
	err = stmt.QueryRowContext(ctx, email).Scan(&verified)
	if err != nil {
		return false, fmt.Errorf("%s: failed to query verified: %w", op, err)
	}

	return verified, nil
}

func (r *Repo) CheckPasswordOTPVerifiedForTooLong(ctx context.Context, email string) (bool, error) {
	const op = "repo.confirmEmail.CheckPasswordOTPVerified"

	stmt, err := r.DB.PrepareContext(ctx, "SELECT expires_at FROM password_verifications where email = $1")
	if err != nil {
		return false, fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	var expiresAt time.Time
	err = stmt.QueryRowContext(ctx, email).Scan(&expiresAt)
	if err != nil {
		return false, fmt.Errorf("%s: failed to query expires_at: %w", op, err)
	}

	tenMinutesAgo := time.Now().UTC().Add(-10 * time.Minute)

	return expiresAt.Before(tenMinutesAgo), nil
}

func (r *Repo) UpdateEmailOTP(ctx context.Context, userID string) error {
	const op = "repo.confirmEmail.UpdateEmailOTP"

	newOTP := otp.GenerateOTP()
	newExpiresAt := time.Now().UTC().Add(5 * time.Minute)

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

func (r *Repo) UpdatePasswordOTP(ctx context.Context, email string) error {
	const op = "repo.confirmEmail.UpdatePasswordOTP"

	newOTP := otp.GenerateOTP()
	newExpiresAt := time.Now().UTC().Add(2 * time.Minute)

	stmt, err := r.DB.PrepareContext(ctx, "UPDATE password_verifications SET expires_at = $1, confirmation_code = $2 WHERE email = $3")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newExpiresAt, newOTP, email)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update: %w", op, err)
	}

	return nil
}

func (r *Repo) UpdatePasswordOTPFalse(ctx context.Context, email string) error {
	const op = "repo.confirmEmail.UpdatePasswordOTP"

	stmt, err := r.DB.PrepareContext(ctx, "UPDATE password_verifications SET is_verified = false WHERE email = $1")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, email)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update: %w", op, err)
	}

	return nil
}

func (r *Repo) ResetPasswordOTP(ctx context.Context, email string) error {
	const op = "repo.confirmEmail.UpdatePasswordOTP"

	newOTP := otp.GenerateOTP()
	newExpiresAt := time.Now().UTC().Add(2 * time.Minute)

	stmt, err := r.DB.PrepareContext(ctx, "UPDATE password_verifications SET expires_at = $1, confirmation_code = $2, is_verified = false WHERE email = $3")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare query: %w", op, err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newExpiresAt, newOTP, email)
	if err != nil {
		return fmt.Errorf("%s: failed to execute update: %w", op, err)
	}

	return nil
}
