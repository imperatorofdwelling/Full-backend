package confirmEmail

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/pkg/sendMail"
)

type Service struct {
	ConfirmEmailRepo interfaces.ConfirmEmailRepository
	UserRepo         interfaces.UserRepository
}

func (s *Service) CreateOTP(ctx context.Context, userID string) error {
	const op = "service.confirmEmail.CreateOTP"

	exist, err := s.ConfirmEmailRepo.CheckOTPExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.ConfirmEmailRepo.CheckOTPNotExpired(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if !expired {
			return fmt.Errorf("OTP already exists and is not expired")
		}

		err = s.ConfirmEmailRepo.UpdateOTP(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to update expired OTP: %w", op, err)
		}

		userOTP, err := s.ConfirmEmailRepo.GetOTP(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to get OTP: %w", op, err)
		}

		err = s.sendOTPEmail(ctx, userID, userOTP)
		if err != nil {
			return fmt.Errorf("%s : failed to send OTP: %w", op, err)
		}

		return nil
	}

	userOTP, err := s.ConfirmEmailRepo.CreateOTP(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	err = s.sendOTPEmail(ctx, userID, userOTP)
	if err != nil {
		return fmt.Errorf("%s : failed to send OTP: %w", op, err)
	}

	return nil
}

func (s *Service) sendOTPEmail(ctx context.Context, userID, userOTP string) error {
	const op = "service.confirmEmail.sendOTPEmail"

	userMail, err := s.UserRepo.FindUserByID(ctx, uuid.FromStringOrNil(userID))
	if err != nil {
		return fmt.Errorf("%s : failed to get user email: %w", op, err)
	}

	err = sendMail.SimpleEmailSend(userMail.Email, userOTP)
	if err != nil {
		return fmt.Errorf("%s : failed to send email to user: %w", op, err)
	}

	return nil
}
