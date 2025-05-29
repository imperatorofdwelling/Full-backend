package interfaces

import (
	"context"
	"net/http"
)

//go:generate mockery --name NotificationRepository
type (
	NotificationRepository interface {
		SetDeviceTokenToUser(ctx context.Context, token string, userId string) (string, error)
		GetDeviceTokenFromUser(ctx context.Context, userId int64) (string, error)
	}
)

//go:generate mockery --name NotificationService
type (
	NotificationService interface {
		SetDeviceTokenToUser(ctx context.Context, token string, userId string) (string, error)
		GetDeviceTokenFromUser(ctx context.Context, userId int64) (string, error)
	}
)

type (
	NotificationHandler interface {
		SetDeviceToken(w http.ResponseWriter, r *http.Request)
	}
)
