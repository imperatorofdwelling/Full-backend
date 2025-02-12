package kafka

import "log/slog"

type Client struct {
	Producer *Producer
	Consumer *Consumer
	Log      *slog.Logger
}

func NewClient(producer *Producer, consumer *Consumer, log *slog.Logger) *Client {
	return &Client{
		Producer: producer,
		Consumer: consumer,
		Log:      log,
	}
}
