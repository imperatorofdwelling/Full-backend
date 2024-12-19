package confirmEmail

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
)

type Service struct {
	Repo interfaces.ConfirmEmailRepository
}

func (s *Service) CreateOTP(ctx context.Context, userID string) error {
	const op = "service.confirmEmail.CreateOTP"

	exist, err := s.Repo.CheckOTPExists(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}
	if exist {
		expired, err := s.Repo.CheckOTPNotExpired(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to check if OTP is expired: %w", op, err)
		}
		if !expired {
			return fmt.Errorf("OTP already exists and is not expired")
		}

		err = s.Repo.UpdateOTP(ctx, userID)
		if err != nil {
			return fmt.Errorf("%s : failed to update expired OTP: %w", op, err)
		}
		return nil
	}

	err = s.Repo.CreateOTP(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
