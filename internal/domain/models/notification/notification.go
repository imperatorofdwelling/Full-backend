package notification

type (
	DeviceInfo struct {
		Token    string `json:"device_token"`
		Platform int    `json:"device_platform"`
	}

	Notification struct {
		Tokens   []string `json:"tokens"`
		Platform int      `json:"platform"`
		Message  string   `json:"message"`
		Title    string   `json:"title,omitempty"`
		Topic    string   `json:"topic,omitempty"` // Для iOS (APNs)
	}
)
