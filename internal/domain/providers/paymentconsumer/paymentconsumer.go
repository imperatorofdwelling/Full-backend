package paymentconsumer

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka/consumer"
	"log/slog"
	"sync"
)

var (
	con     *consumer.PaymentConsumer
	conOnce sync.Once
)

var PaymentConsumerProviderSet wire.ProviderSet = wire.NewSet(
	ProvidePaymentConsumer,
)

func ProvidePaymentConsumer(log *slog.Logger) *consumer.PaymentConsumer {
	conOnce.Do(func() {
		con = &consumer.PaymentConsumer{
			Log: log,
		}
	})

	return con
}
