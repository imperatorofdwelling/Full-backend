package notification

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	domainNoticitaion "github.com/imperatorofdwelling/Full-backend/internal/domain/models/notification"
)

type Service struct {
	Repo       interfaces.NotificationRepository
	PushClient interfaces.PushNotificationSender
	//cfg        *config.Server
}

func (s *Service) SetDeviceTokenToUser(ctx context.Context, request domainNoticitaion.DeviceInfo, userId string) (domainNoticitaion.DeviceInfo, error) {
	const op = "service.notification.SetDeviceTokenToUser"

	updatedToken, err := s.Repo.SetDeviceTokenToUser(ctx, request, userId)
	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	return updatedToken, nil
}

func (s *Service) SendNotificationToUser(ctx context.Context, userId string, message string) error {
	const op = "service.notification.SendNotificationToUser"
	deviceInfo, err := s.Repo.GetDeviceTokenFromUser(ctx, userId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	notification := domainNoticitaion.Notification{
		Tokens:   []string{deviceInfo.Token},
		Platform: deviceInfo.Platform,
		Message:  message,
	}

	return s.PushClient.SendPushNotification(ctx, notification)
}

//
//func (s *Service) sendPushNotification(notification domainNoticitaion.Notification) (string, error) {
//	// Преобразуем уведомление в JSON
//	jsonData, err := json.Marshal(notification)
//	if err != nil {
//		return "", err
//	}
//
//	// Отправляем POST-запрос на Gorush
//	resp, err := http.Post(fmt.Sprintf("http://%s:%s", s.cfg.Addr, s.cfg.Port)+"/push", "application/json", bytes.NewBuffer(jsonData))
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("failed to send notification, status code: %d", resp.StatusCode)
//	}
//	return "Notification sent successfully", nil
//}
