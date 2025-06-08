package notification

import (
	"context"
	"database/sql"
	"fmt"
	domainNoticitaion "github.com/imperatorofdwelling/Full-backend/internal/domain/models/notification"
	"github.com/imperatorofdwelling/Full-backend/pkg/encryption"
)

// Запрос к бд за токеном

type Repo struct {
	Db *sql.DB
}

func (r *Repo) SetDeviceTokenToUser(ctx context.Context, deviceInfo domainNoticitaion.DeviceInfo, userId string) (domainNoticitaion.DeviceInfo, error) {
	const op = "repo.notification.SetDeviceTokenToUser"

	// Шифруем токен перед сохранением
	encryptedToken, err := encryption.EncryptToken(deviceInfo.Token)
	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: encryption failed: %w", op, err)
	}

	stmt, err := r.Db.PrepareContext(ctx, `
        UPDATE users 
        SET device_token = $1 , device_platform = $2
        WHERE user_id = $3
        RETURNING device_token, device_platform
    `)
	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var updatedDeviceToken domainNoticitaion.DeviceInfo
	err = stmt.QueryRowContext(ctx, encryptedToken, deviceInfo.Platform, userId).Scan(&updatedDeviceToken)
	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	return updatedDeviceToken, nil
}

func (r *Repo) GetDeviceTokenFromUser(ctx context.Context, userId string) (domainNoticitaion.DeviceInfo, error) {
	const op = "repo.notification.GetDeviceTokenFromUser"

	stmt, err := r.Db.PrepareContext(ctx, `
        SELECT device_token, device_platform FROM users WHERE user_id = $1
    `)

	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var deviceInfo domainNoticitaion.DeviceInfo
	err = stmt.QueryRowContext(ctx, userId).Scan(&deviceInfo.Token, &deviceInfo.Platform)
	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: %w", op, err)
	}

	// Дешифруем токен перед возвратом
	deviceInfo.Token, err = encryption.DecryptToken(deviceInfo.Token)
	if err != nil {
		return domainNoticitaion.DeviceInfo{}, fmt.Errorf("%s: decryption failed: %w", op, err)
	}

	return deviceInfo, nil
}
