package auth

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"log"
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
		log.Println("Error checking user existence: " + err.Error())
		return uuid.Nil, err
	}

	if userExists {
		log.Println("User already exists")
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrUserAlreadyExists)
	}
	if !s.validate(user) {
		log.Println("Invalid user")
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrValid)
	}

	id, err := s.AuthRepo.Register(ctx, user)
	if err != nil {
		log.Println("Error registering user: " + err.Error())
		return uuid.Nil, err
	}

	userFound, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		log.Println("Error finding user: " + err.Error())
		return uuid.Nil, err
	}

	if userFound.ID == uuid.Nil {
		log.Println("User not found")
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	return userFound.ID, nil
}

func (s *Service) Login(ctx context.Context, user model.Login) (uuid.UUID, int, error) {
	const op = "service.auth.Login"
	userExists, err := s.UserRepo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return uuid.Nil, -1, err
	}

	if !userExists {
		return uuid.Nil, -1, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	id, roleID, err := s.AuthRepo.Login(ctx, user)
	if err != nil {
		return id, -1, err
	}

	return id, roleID, err
}

func (s *Service) CheckOTP(ctx context.Context, userID, otp string) error {
	const op = "service.auth.CheckOTP"

	isVerified, err := s.AuthRepo.CheckIfUserValidated(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if isVerified {
		return fmt.Errorf("user is already verified")
	}

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

	err = s.AuthRepo.EmailVerification(ctx, userID)
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
