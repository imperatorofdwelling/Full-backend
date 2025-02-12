package payment

//type Handler struct {
//	Kafka           *kafka.Client
//	Log             *slog.Logger
//	WaitForResponse map[string]chan consumer.PaymentResponse
//}
//
//func (h *Handler) NewPaymentHandler(r chi.Router) {
//	r.Route("/payment", func(r chi.Router) {
//		r.Group(func(r chi.Router) {
//			r.Post("/", h.MakePayment)
//		})
//	})
//}
//
//func (h *Handler) MakePayment(w http.ResponseWriter, r *http.Request) {
//	const op = "handler.payment.MakePayment"
//
//	h.Log = h.Log.With(
//		slog.String("op", op),
//		slog.String("request_id", middleware.GetReqID(r.Context())),
//	)
//
//	idempotenceKey := r.Header.Get("Idempotence-Key")
//	if idempotenceKey == "" {
//		h.Log.Error(handler.ErrGettingIdempotenceKey.Error())
//		responseApi.WriteError(w, r, http.StatusBadRequest, handler.ErrGettingIdempotenceKey.Error())
//		return
//	}
//
//	var payment yoomodel.Payment
//
//	err := render.DecodeJSON(r.Body, &payment)
//	if err != nil {
//		h.Log.Error("error decoding payment json", err.Error())
//		responseApi.WriteError(w, r, http.StatusBadRequest, err)
//		return
//	}
//
//	responseChan := make(chan consumer.PaymentResponse)
//	h.WaitForResponse[idempotenceKey] = responseChan
//
//	err = h.Kafka.Producer.SendMessage(kafka.PaymentReqTopic, idempotenceKey, payment)
//	if err != nil {
//		h.Log.Error("failed to send message", "error", err.Error())
//		delete(h.WaitForResponse, idempotenceKey)
//		responseApi.WriteError(w, r, http.StatusServiceUnavailable, err)
//		return
//	}
//
//	select {
//	case response := <-responseChan:
//		responseApi.WriteJson(w, r, http.StatusOK, response)
//	case <-time.After(10 * time.Second):
//		h.Log.Error("timed out waiting for response")
//		responseApi.WriteError(w, r, http.StatusGatewayTimeout, errors.New("timed out waiting for response"))
//	}
//
//	delete(h.WaitForResponse, idempotenceKey)
//}
