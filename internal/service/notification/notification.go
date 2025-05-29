package notification

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
)

type Service struct {
	Repo interfaces.NotificationRepository
}

func (s *Service) SetDeviceTokenToUser(ctx context.Context, token string, userId string) (string, error) {
	const op = "service.notification.SetDeviceTokenToUser"

	updatedToken, err := s.Repo.SetDeviceTokenToUser(ctx, token, userId)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return updatedToken, nil
}
