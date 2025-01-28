package interfaces

import (
	"context"
	"net/http"
)

//go:generate mockery --name ConfirmEmailRepository
type (
	ConfirmEmailRepository interface {
		CreateEmailOTP(ctx context.Context, userId string) (string, error)
		CreatePasswordOTP(ctx context.Context, email string) (string, error)
		CreateEmailChangeOTP(ctx context.Context, userID string) (string, error)
		GetEmailOTP(ctx context.Context, userId string) (string, error)
		GetPasswordOTP(ctx context.Context, email string) (string, error)
		GetEmailChangeOTP(ctx context.Context, email string) (string, error)
		CheckEmailOTPExists(ctx context.Context, userID string) (bool, error)
		CheckEmailOTPNotExpired(ctx context.Context, userID string) (bool, error)
		CheckPasswordOTPExists(ctx context.Context, email string) (bool, error)
		CheckPasswordOTPNotExpired(ctx context.Context, email string) (bool, error)
		CheckPasswordOTPVerified(ctx context.Context, email string) (bool, error)
		CheckPasswordOTPVerifiedForTooLong(ctx context.Context, email string) (bool, error)
		CheckEmailChangeOTPExists(ctx context.Context, userID string) (bool, error)
		CheckEmailChangeOTPNotExpired(ctx context.Context, email string) (bool, error)
		UpdateEmailOTP(ctx context.Context, userID string) error
		UpdatePasswordOTP(ctx context.Context, email string) error
		UpdatePasswordOTPFalse(ctx context.Context, email string) error
		UpdateEmailChangeOTP(ctx context.Context, userID string) error
		ResetPasswordOTP(ctx context.Context, email string) error
	}
)

//go:generate mockery --name ConfirmEmailService
type (
	ConfirmEmailService interface {
		CreateOTPEmail(ctx context.Context, userID string) error
		CreateOTPPassword(ctx context.Context, email string) error
		SendOtpForEmailChange(ctx context.Context, userID string) error
		SendOTPEmail(ctx context.Context, userID, userOTP, title string) error
	}
)

type (
	ConfirmEmailHandler interface {
		CreateOTPEmail(w http.ResponseWriter, r *http.Request)
		CreateOTPPassword(w http.ResponseWriter, r *http.Request)
		SendOtpForEmailChange(w http.ResponseWriter, r *http.Request)
	}
)
