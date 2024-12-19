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
	const op = "service.user.Registration"

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
	const op = "service.user.Login"
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

func (s *Service) CheckOTP(ctx context.Context, userID, otp string) error {
	const op = "service.confirmEmail.CheckOTP"

	exist, err := s.ConfirmEmailRepo.CheckOTPExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.ConfirmEmailRepo.CheckOTPNotExpired(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if expired {
			return fmt.Errorf("OTP is expired")
		}
	}

	otpFromDB, err := s.ConfirmEmailRepo.GetOTP(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if otpFromDB != otp {
		return fmt.Errorf("invalid OTP")
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
