package payment

import (
	"errors"
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka"
	"github.com/imperatorofdwelling/Full-backend/internal/api/kafka/consumer"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"log/slog"
	"net/http"
	"time"
)

type Handler struct {
	Kafka           *kafka.Client
	Log             *slog.Logger
	WaitForResponse map[string]chan consumer.PaymentResponse
}

func (h *Handler) NewPaymentHandler(r chi.Router) {
	r.Route("/payment", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/", h.MakePayment)
		})
	})
}

// MakePayment godoc
//
//	@Summary		Create payment
//	@Description	Create payment (with yookassa model)
//	@Tags			payment
//	@Accept			application/json
//	@Produce		json
//	@Param	Idempotence-Key header string	true	"Idempotence-Key"
//	@Param	_ body yoomodel.Payment	true	"request yookassa payment"
//	@Success		200	{object}		yoomodel.PaymentRes	"ok"
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
//	@Router			/payment [post]
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

	responseChan := make(chan consumer.PaymentResponse)
	h.WaitForResponse[idempotenceKey] = responseChan

	err = h.Kafka.Producer.SendMessage(kafka.PaymentReqTopic, idempotenceKey, payment)
	if err != nil {
		h.Log.Error("failed to send message", "error", err.Error())
		delete(h.WaitForResponse, idempotenceKey)
		responseApi.WriteError(w, r, http.StatusServiceUnavailable, err)
		return
	}

	select {
	case response := <-responseChan:
		responseApi.WriteJson(w, r, http.StatusOK, response.Result)
	case <-time.After(10 * time.Second):
		h.Log.Error("timed out waiting for response")
		responseApi.WriteError(w, r, http.StatusGatewayTimeout, errors.New("timed out waiting for response"))
	}

	delete(h.WaitForResponse, idempotenceKey)
}
