package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"net/http"
)

//go:generate mockery --name AuthRepository
type (
	AuthRepository interface {
		Register(ctx context.Context, user auth.Registration) (uuid.UUID, error)
		Login(ctx context.Context, user auth.Login) (uuid.UUID, error)
		EmailVerification(ctx context.Context, userId string) error
		PasswordVerification(ctx context.Context, email string) error
		CheckIfUserEmailValidated(ctx context.Context, userId string) (bool, error)
		ConfirmEmailChangeOTP(ctx context.Context, userId string) error
	}
)

//go:generate mockery --name AuthService
type (
	AuthService interface {
		Register(ctx context.Context, user auth.Registration) (uuid.UUID, error)
		Login(ctx context.Context, user auth.Login) (uuid.UUID, error)
		CheckEmailOTP(ctx context.Context, userID, otp string) error
		CheckPasswordOTP(ctx context.Context, email, otp string) error
		CheckEmailChangeOTP(ctx context.Context, userID, otp string) error
	}
)

type (
	AuthHandler interface {
		Registration(w http.ResponseWriter, r *http.Request)
		LoginUser(w http.ResponseWriter, r *http.Request)
		ConfirmEmailOTP(w http.ResponseWriter, r *http.Request)
	}
)
