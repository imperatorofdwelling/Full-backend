package auth

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"net/mail"
	"strings"
)

type Service struct {
	AuthRepo         interfaces.AuthRepository
	UserRepo         interfaces.UserRepository
	ConfirmEmailRepo interfaces.ConfirmEmailRepository
}

func (s *Service) Register(ctx context.Context, user model.Registration) (uuid.UUID, error) {
	const op = "service.auth.Registration"

	userExists, err := s.UserRepo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if userExists {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrUserAlreadyExists)
	}
	if !s.validate(user) {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrValid)
	}

	id, err := s.AuthRepo.Register(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}

	userFound, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		return uuid.Nil, err
	}

	if userFound.ID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	return userFound.ID, nil
}

func (s *Service) Login(ctx context.Context, user model.Login) (uuid.UUID, error) {
	const op = "service.auth.Login"
	userExists, err := s.UserRepo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if !userExists {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	id, err := s.AuthRepo.Login(ctx, user)
	if err != nil {
		return id, err
	}

	return id, err
}

func (s *Service) CheckEmailOTP(ctx context.Context, userID, otp string) error {
	const op = "service.auth.CheckEmailOTP"

	isVerified, err := s.AuthRepo.CheckIfUserEmailValidated(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if isVerified {
		return fmt.Errorf("user is already verified")
	}

	exist, err := s.ConfirmEmailRepo.CheckEmailOTPExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.ConfirmEmailRepo.CheckEmailOTPNotExpired(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if expired {
			return fmt.Errorf("OTP is expired")
		}
	}

	otpFromDB, err := s.ConfirmEmailRepo.GetEmailOTP(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if otpFromDB != otp {
		return fmt.Errorf("invalid OTP")
	}

	err = s.AuthRepo.EmailVerification(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (s *Service) CheckPasswordOTP(ctx context.Context, email, otp string) error {
	const op = "service.auth.CheckPasswordOTP"

	exist, err := s.ConfirmEmailRepo.CheckPasswordOTPExists(ctx, email)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.ConfirmEmailRepo.CheckPasswordOTPNotExpired(ctx, email)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if expired {
			return fmt.Errorf("OTP is expired")
		}
	}

	otpFromDB, err := s.ConfirmEmailRepo.GetPasswordOTP(ctx, email)
	if err != nil {
		return fmt.Errorf("%s: failed to get password OTP %w", op, err)
	}

	if otpFromDB != otp {
		return fmt.Errorf("invalid OTP")
	}

	err = s.AuthRepo.PasswordVerification(ctx, email)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}

func (s *Service) validate(user model.Registration) bool {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" {
		return false
	}

	_, err := mail.ParseAddress(user.Email)
	return err == nil
}
