package pushNotification

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	domainNoticitaion "github.com/imperatorofdwelling/Full-backend/internal/domain/models/notification"
	"net/http"
)

type GorushClient struct {
	BaseURL string
}

func (c *GorushClient) SendPushNotification(ctx context.Context, notification domainNoticitaion.Notification) error {
	jsonData, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.BaseURL+"/api/push", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("create request error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("gorush returned status: %d", resp.StatusCode)
	}

	return nil
}
