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

func (s *Service) CreateOTPEmail(ctx context.Context, userID string) error {
	const op = "service.confirmEmail.CreateOTP"

	exist, err := s.ConfirmEmailRepo.CheckEmailOTPExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.ConfirmEmailRepo.CheckEmailOTPNotExpired(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if !expired {
			return fmt.Errorf("OTP already exists and is not expired")
		}

		err = s.ConfirmEmailRepo.UpdateEmailOTP(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to update expired OTP: %w", op, err)
		}

		userOTP, err := s.ConfirmEmailRepo.GetEmailOTP(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to get OTP: %w", op, err)
		}

		err = s.SendOTPEmail(ctx, userID, userOTP, "Registration")
		if err != nil {
			return fmt.Errorf("%s : failed to send OTP: %w", op, err)
		}

		return nil
	}

	userOTP, err := s.ConfirmEmailRepo.CreateEmailOTP(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	err = s.SendOTPEmail(ctx, userID, userOTP, "Registration")
	if err != nil {
		return fmt.Errorf("%s : failed to send OTP: %w", op, err)
	}

	return nil
}

func (s *Service) CreateOTPPassword(ctx context.Context, email string) error {
	const op = "service.confirmEmail.CreateOTPPassword"

	userExists, err := s.UserRepo.CheckUserExists(ctx, email)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if !userExists {
		return fmt.Errorf("%s : user does not exist", op)
	}

	userID, err := s.UserRepo.GetUserIDByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("%s : failed to get user ID: %w", op, err)
	}

	exist, err := s.ConfirmEmailRepo.CheckPasswordOTPExists(ctx, email)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.ConfirmEmailRepo.CheckPasswordOTPNotExpired(ctx, email)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if !expired {
			return fmt.Errorf("OTP already exists and is not expired")
		}

		err = s.ConfirmEmailRepo.UpdatePasswordOTP(ctx, email)
		if err != nil {
			return fmt.Errorf("%s : failed to update expired OTP: %w", op, err)
		}

		userOTP, err := s.ConfirmEmailRepo.GetPasswordOTP(ctx, email)
		if err != nil {
			return fmt.Errorf("%s : failed to get OTP: %w", op, err)
		}

		err = s.SendOTPEmail(ctx, userID, userOTP, "Password Reset")
		if err != nil {
			return fmt.Errorf("%s : failed to send OTP: %w", op, err)
		}

		return nil
	}

	userOTP, err := s.ConfirmEmailRepo.CreatePasswordOTP(ctx, email)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	err = s.SendOTPEmail(ctx, userID, userOTP, "Password Reset")
	if err != nil {
		return fmt.Errorf("%s : failed to send OTP: %w", op, err)
	}

	return nil
}

func (s *Service) SendOTPEmail(ctx context.Context, userID, userOTP, title string) error {
	const op = "service.confirmEmail.sendOTPEmail"

	userMail, err := s.UserRepo.FindUserByID(ctx, uuid.FromStringOrNil(userID))
	if err != nil {
		return fmt.Errorf("%s : failed to get user email: %w", op, err)
	}

	err = sendMail.SimpleEmailSend(userMail.Email, userOTP, title)
	if err != nil {
		return fmt.Errorf("%s : failed to send email to user: %w", op, err)
	}

	return nil
}
