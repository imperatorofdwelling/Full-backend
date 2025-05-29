package notify

import "fmt"

var deviceTokens = make(map[string][]string)

func getUserDeviceTokens(userID string) ([]string, error) {
	// В реальном приложении - запрос к базе данных
	tokens, ok := deviceTokens[userID]
	if !ok {
		return nil, fmt.Errorf("no tokens found for user %s", userID)
	}
	return tokens, nil
}

// Notification представляет структуру уведомления
type Notification struct {
	Tokens   []string `json:"tokens"`
	Platform int      `json:"platform"`
	Message  string   `json:"message"`
	Title    string   `json:"title,omitempty"`
	Topic    string   `json:"topic,omitempty"` // Для iOS (APNs)
}

func sendNotificationToUser(userID, message string) error {
	tokens, err := getUserDeviceTokens(userID)
	if err != nil {
		return err
	}

	notification := Notification{
		Tokens:   tokens,
		Platform: 2, // Определите платформу на основе данных из БД
		Message:  message,
	}

	return SendPushNotification("http://localhost:8088", notification)
}
