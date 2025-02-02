package providers

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"log/slog"
	"sync"
)

var (
	client     *kafka.Client
	clientOnce sync.Once
)

var KafkaProviderSet wire.ProviderSet = wire.NewSet(
	kafka.NewKafkaProducer,
	kafka.NewClient,
)

func ProvideKafkaClient(producer *kafka.Producer, log *slog.Logger) *kafka.Client {
	clientOnce.Do(func() {
		client = &kafka.Client{
			Producer: producer,
			Log:      log,
		}
	})

	return client
}
