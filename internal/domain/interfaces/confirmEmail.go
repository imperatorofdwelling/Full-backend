package interfaces

import (
	"context"
	"net/http"
)

//go:generate mockery --name ConfirmEmailRepository
type (
	ConfirmEmailRepository interface {
		CreateOTP(ctx context.Context, userId string) error
		GetOTP(ctx context.Context, userId string) (string, error)
		CheckOTPExists(ctx context.Context, userID string) (bool, error)
		CheckOTPNotExpired(ctx context.Context, userID string) (bool, error)
		UpdateOTP(ctx context.Context, userID string) error
	}
)

//go:generate mockery --name ConfirmEmailService
type (
	ConfirmEmailService interface {
		CreateOTP(ctx context.Context, userID string) error
	}
)

type (
	ConfirmEmailHandler interface {
		CreateOTP(w http.ResponseWriter, r *http.Request)
	}
)
