package providers

import (
	"github.com/IBM/sarama"
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

func ProvideKafkaProducer(sarama sarama.SyncProducer) *kafka.Producer {
	producerOnce.Do(func() {
		producer = &kafka.Producer{
			Sarama: sarama,
		}
	})

	return producer
}
