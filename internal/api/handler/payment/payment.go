package payment

import (
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"log/slog"
	"net/http"
)

type Handler struct {
	Kafka *kafka.Producer
	Log   *slog.Logger
}

func (h *Handler) NewPaymentHandler(r chi.Router) {
	r.Route("/payment", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", h.MakePayment)
		})
	})
}

func (h *Handler) MakePayment(w http.ResponseWriter, r *http.Request) {
	const op = "handler.payment.MakePayment"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	idempotenceKey := r.Header.Get("Idempotence-Key")
	if idempotenceKey == "" {
		h.Log.Error(handler.ErrGettingIdempotenceKey.Error())
		responseApi.WriteError(w, r, http.StatusBadRequest, handler.ErrGettingIdempotenceKey.Error())
		return
	}

	var payment yoomodel.Payment

	err := render.DecodeJSON(r.Body, &payment)
	if err != nil {
		h.Log.Error("error decoding payment json", err.Error())
		responseApi.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	err = h.Kafka.SendMessage(kafka.PaymentTopic, idempotenceKey, payment)
	if err != nil {
		h.Log.Error("failed to send message", "error", err.Error())
		responseApi.WriteError(w, r, http.StatusServiceUnavailable, err)
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully sent")
}
