package providers

//var (
//	hdl     *pHdl.Handler
//	hdlOnce sync.Once
//)
//
//var PaymentProviderSet wire.ProviderSet = wire.NewSet(
//	ProvidePaymentHandler,
//
//	wire.Bind(new(interfaces.PaymentHandler), new(*pHdl.Handler)),
//)
//
//func ProvidePaymentHandler(kafka *kafka.Client, log *slog.Logger, waitForResponse map[string]chan consumer.PaymentResponse) *pHdl.Handler {
//	hdlOnce.Do(func() {
//		hdl = &pHdl.Handler{
//			Kafka:           kafka,
//			Log:             log,
//			WaitForResponse: waitForResponse,
//		}
//	})
//
//	return hdl
//}
