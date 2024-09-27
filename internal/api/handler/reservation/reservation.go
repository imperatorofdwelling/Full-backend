package reservation

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.ReservationService
	Log *slog.Logger
}

func (h *Handler) NewReservationHandler(r chi.Router) {
	r.Route("/reservation", func(r chi.Router) {

	})
}

// CreateReservation godoc
//
//		@Summary		Create Reservation
//		@Description	Create reservation (arrived and departure should be TIMESTAMP type)
//		@Tags			reservations
//		@Accept			application/json
//		@Produce		json
//	 	@Param			request 	body	reservation.ReservationEntity	true	"Create reservation request"
//		@Success		201	{object}		string	"created"
//		@Failure		400		{object}	responseApi.ResponseError			"Error"
//		@Failure		default		{object}	responseApi.ResponseError			"Error"
//		@Router			/reservation/create [post]
func (h *Handler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.CreateReservation"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var reserv reservation.ReservationEntity

	err := render.DecodeJSON(r.Body, &reserv)
	if err != nil {
		h.Log.Error("failed to decode JSON", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateReservation(context.Background(), &reserv)
	if err != nil {
		h.Log.Error("failed to create reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created reservation")
}

// UpdateReservation godoc
//
//	@Summary		Update Reservation
//	@Description	Update reservation by id
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			reservationId	path		string		true	"reservation id"
//	@Param			request	body	reservation.ReservationUpdateEntity			true	"update reservation request"
//	@Success		200	{string}		string	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/reservation/{reservationId} [put]
func (h *Handler) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.UpdateReservation"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var newReserv reservation.ReservationUpdateEntity

	err := render.DecodeJSON(r.Body, &newReserv)
	if err != nil {
		h.Log.Error("failed to decode JSON", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.UpdateReservation(context.Background(), &newReserv)
	if err != nil {
		h.Log.Error("failed to update reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully updated reservation")
}
