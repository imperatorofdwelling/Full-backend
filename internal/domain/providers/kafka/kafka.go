package providers

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
)

var KafkaProviderSet wire.ProviderSet = wire.NewSet(
	kafka.NewKafkaProducer,
	kafka.NewKafkaConsumer,
	kafka.NewClient,
)
