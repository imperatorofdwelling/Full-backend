package payment

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"log/slog"
	"net/http"
)

type Handler struct {
	Kafka interfaces.KafkaProducer
	Log   *slog.Logger
}

func (h *Handler) NewPaymentHandler(r chi.Router) {
	r.Route("/payment", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/", h.MakePayment)
		})
	})
}

func (h *Handler) MakePayment(w http.ResponseWriter, r *http.Request) {
	const op = "handler.payment.MakePayment"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	message := "Hello World!"
	messageJSON, err := json.Marshal(message)

	err = h.Kafka.SendMessage(kafka.PaymentTopic, messageJSON)
	if err != nil {
		h.Log.Error("failed to send message", "error", err.Error())
		responseApi.WriteError(w, r, http.StatusServiceUnavailable, err)
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully sent")
}
