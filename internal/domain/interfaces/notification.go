package interfaces

import (
	"context"
	domainNoticitaion "github.com/imperatorofdwelling/Full-backend/internal/domain/models/notification"
	"net/http"
)

//go:generate mockery --name NotificationRepository
type (
	NotificationRepository interface {
		SetDeviceTokenToUser(ctx context.Context, deviceInfo domainNoticitaion.DeviceInfo, userId string) (domainNoticitaion.DeviceInfo, error)
		GetDeviceTokenFromUser(ctx context.Context, userId string) (domainNoticitaion.DeviceInfo, error)
	}
)

//go:generate mockery --name NotificationService
type (
	NotificationService interface {
		SetDeviceTokenToUser(ctx context.Context, request domainNoticitaion.DeviceInfo, userId string) (domainNoticitaion.DeviceInfo, error)
		SendNotificationToUser(ctx context.Context, userId string, message string) error
	}
)

type (
	NotificationHandler interface {
		SetDeviceToken(w http.ResponseWriter, r *http.Request)
	}
	PushNotificationSender interface {
		SendPushNotification(ctx context.Context, notify domainNoticitaion.Notification) error
	}
)
