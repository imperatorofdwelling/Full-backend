package kafka

import (
	"log/slog"
)

type Client struct {
	Producer *Producer
	Log      *slog.Logger
}

func NewClient(producer *Producer, log *slog.Logger) *Client {
	return &Client{
		Producer: producer,
		Log:      log,
	}
}
