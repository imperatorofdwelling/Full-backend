package providers

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"sync"
)

var (
	producer     *kafka.Producer
	producerOnce sync.Once
)

var KafkaProviderSet wire.ProviderSet = wire.NewSet(
	ProvideKafkaProducer,

	wire.Bind(new(interfaces.KafkaProducer), new(*kafka.Producer)),
)

func ProvideKafkaProducer() *kafka.Producer {
	producerOnce.Do(func() {
		producer = &kafka.Producer{}
	})

	return producer
}
