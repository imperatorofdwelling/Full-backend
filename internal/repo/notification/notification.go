package notification

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/pkg/encryption"
)

// Запрос к бд за токеном

type Repo struct {
	Db *sql.DB
}

func (r *Repo) SetDeviceTokenToUser(ctx context.Context, token string, userId string) (string, error) {
	const op = "repo.notification.SetDeviceTokenToUser"

	// Шифруем токен перед сохранением
	encryptedToken, err := encryption.EncryptToken(token)
	if err != nil {
		return "", fmt.Errorf("%s: encryption failed: %w", op, err)
	}

	stmt, err := r.Db.PrepareContext(ctx, `
        UPDATE users 
        SET device_token = $1 
        WHERE user_id = $2
        RETURNING device_token
    `)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var updatedToken string
	err = stmt.QueryRowContext(ctx, encryptedToken, userId).Scan(&updatedToken)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return updatedToken, nil
}

func (r *Repo) GetDeviceTokenFromUser(ctx context.Context, userId int64) (string, error) {
	const op = "repo.notification.GetDeviceTokenFromUser"

	stmt, err := r.Db.PrepareContext(ctx, `
        SELECT device_token FROM users WHERE user_id = $1
    `)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var encryptedToken string
	err = stmt.QueryRowContext(ctx, userId).Scan(&encryptedToken)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	// Дешифруем токен перед возвратом
	decryptedToken, err := encryption.DecryptToken(encryptedToken)
	if err != nil {
		return "", fmt.Errorf("%s: decryption failed: %w", op, err)
	}

	return decryptedToken, nil
}
