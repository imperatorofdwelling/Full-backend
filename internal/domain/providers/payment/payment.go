package providers

import (
	"github.com/google/wire"
	pHdl "github.com/imperatorofdwelling/Full-backend/internal/api/handler/payment"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"log/slog"
	"sync"
)

var (
	hdl     *pHdl.Handler
	hdlOnce sync.Once
)

var PaymentProviderSet wire.ProviderSet = wire.NewSet(
	ProvidePaymentHandler,

	wire.Bind(new(interfaces.PaymentHandler), new(*pHdl.Handler)),
)

func ProvidePaymentHandler(kafka *kafka.Producer, log *slog.Logger) *pHdl.Handler {
	hdlOnce.Do(func() {
		hdl = &pHdl.Handler{
			Kafka: kafka,
			Log:   log,
		}
	})

	return hdl
}
