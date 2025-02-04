package providers

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"sync"
)

var (
	client     *kafka.Client
	clientOnce sync.Once
)

var KafkaProviderSet wire.ProviderSet = wire.NewSet(
	kafka.NewKafkaProducer,
	kafka.NewKafkaConsumer,
	kafka.NewClient,

	//ProvideKafkaClient,
)

//func ProvideKafkaClient(producer *kafka.Producer, consumer *kafka.Consumer, log *slog.Logger) *kafka.Client {
//	clientOnce.Do(func() {
//		client = &kafka.Client{
//			Producer: producer,
//			Consumer: consumer,
//			Log:      log,
//		}
//	})
//
//	return client
//}
