package providers

import (
	"github.com/google/wire"
	paymentHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/payment"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka/consumer"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"log/slog"
	"sync"
)

var (
	hdl     *paymentHdl.Handler
	hdlOnce sync.Once
)

var PaymentProviderSet wire.ProviderSet = wire.NewSet(
	ProvidePaymentHandler,

	wire.Bind(new(interfaces.PaymentHandler), new(*paymentHdl.Handler)),
)

func ProvidePaymentHandler(kafka *kafka.Client, log *slog.Logger, waitForResponse map[string]chan consumer.PaymentResponse) *paymentHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &paymentHdl.Handler{
			Kafka:           kafka,
			Log:             log,
			WaitForResponse: waitForResponse,
		}
	})

	return hdl
}
