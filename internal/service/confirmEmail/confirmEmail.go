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

	err := s.Repo.CreateOTP(ctx, userID)
	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
