package paymentconsumer

import (
	"github.com/google/wire"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka/consumer"
	"log/slog"
	"sync"
)

var (
	con     *consumer.PaymentConsumerHdl
	conOnce sync.Once
)

var PaymentConsumerProviderSet wire.ProviderSet = wire.NewSet(
	ProvideWaitPaymentForResponseChan,
	ProvidePaymentConsumer,
)

func ProvidePaymentConsumer(log *slog.Logger, waitForResponse map[string]chan consumer.PaymentResponse) *consumer.PaymentConsumerHdl {
	conOnce.Do(func() {
		con = &consumer.PaymentConsumerHdl{
			Log:             log,
			WaitForResponse: waitForResponse,
		}
	})

	return con
}

func ProvideWaitPaymentForResponseChan() map[string]chan consumer.PaymentResponse {
	return make(map[string]chan consumer.PaymentResponse)
}
