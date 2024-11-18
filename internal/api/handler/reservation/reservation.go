package reservation

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
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
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Post("/create", h.CreateReservation)
			r.Put("/update/{reservationID}", h.UpdateReservation)
			r.Delete("/{reservationID}", h.DeleteReservationByID)
			r.Get("/{reservationID}", h.GetReservationByID)
			r.Get("/user/userID", h.GetAllReservationsByUser)
		})
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
//	@Description	Update an existing reservation by its ID
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			reservationId	path		string		true	"ID of the reservation to update"
//	@Param			request		body	reservation.ReservationUpdateEntity	true	"Details to update the reservation"
//	@Success		200	{object}	map[string]interface{}	"Successfully updated reservation"
//	@Failure		400	{object}	responseApi.ResponseError		"Invalid request"
//	@Failure		404	{object}	responseApi.ResponseError		"Reservation not found"
//	@Failure		500	{object}	responseApi.ResponseError		"Internal server error"
//	@Router			/reservation/update/{reservationId} [put]
func (h *Handler) UpdateReservation(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.UpdateReservation"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "reservationID")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse UUID", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var newReserv reservation.ReservationUpdateEntity

	err = render.DecodeJSON(r.Body, &newReserv)
	if err != nil {
		h.Log.Error("failed to decode JSON", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	newReserv.ID = uuID

	err = h.Svc.UpdateReservation(context.Background(), &newReserv)
	if err != nil {
		h.Log.Error("failed to update reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	reserv, err := h.Svc.GetReservationByID(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to find reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"Updated reservation": reserv})
}

// DeleteReservationByID godoc
//
//	@Summary		Delete Reservation
//	@Description	Delete reservation by id
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			reservationID	path		string		true	"reservation id"
//	@Success		200	{string}		string	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/reservation/{reservationID} [delete]
func (h *Handler) DeleteReservationByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.DeleteReservationByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "reservationID")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse UUID", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteReservationByID(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to delete reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully deleted reservation")
}

// GetReservationByID godoc
//
//	@Summary		Get Reservation
//	@Description	Get reservation by id
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			reservationID	path		string		true	"reservation id"
//	@Success		200	{object}		reservation.Reservation	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/reservation/{reservationID} [get]
func (h *Handler) GetReservationByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.GetReservationByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "reservationID")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse UUID", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	reserv, err := h.Svc.GetReservationByID(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to fetch reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, reserv)
}

// GetAllReservationsByUser godoc
//
//	@Summary		Get all Reservations
//	@Description	Get reservation by user id
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			userID	path		string		true	"user id"
//	@Success		200	{object}		[]reservation.Reservation	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/reservation/user/{userID} [get]
func (h *Handler) GetAllReservationsByUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.GetAllReservationsByUser"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "userID")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse UUID", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	reservs, err := h.Svc.GetAllReservationsByUser(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to fetch reservations", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, reservs)
}
