package reservation

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/reservation"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
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
			r.Post("/", h.CreateReservation)

			r.Post("/checkin/{stayId}", h.ConfirmCheckInReservation)
			r.Post("/checkout/{stayId}", h.ConfirmCheckOutReservation)

			r.Put("/{reservationID}", h.UpdateReservation)
			r.Delete("/{reservationID}", h.DeleteReservationByID)

			r.Get("/{reservationID}", h.GetReservationByID)

			r.Get("/user/userID", h.GetAllReservationsByUser)

			r.Get("/free/{userID}", h.GetFreeReservationsByUserID)
			r.Get("/occupied/{userID}", h.GetOccupiedReservationsByUserID)
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
//		@Failure		400		{object}	response.ResponseError			"Error"
//		@Failure		default		{object}	response.ResponseError			"Error"
//		@Router			/reservation [post]
func (h *Handler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.CreateReservation"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logged in")))
		return
	}

	h.Log.Info(userID)

	var reserv reservation.ReservationEntity

	err := render.DecodeJSON(r.Body, &reserv)
	if err != nil {
		h.Log.Error("failed to decode JSON", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CheckReservation(context.Background(), &reserv, userID)
	if err != nil {
		h.Log.Error("failed to check reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	err = h.Svc.CreateReservation(context.Background(), &reserv, userID)
	if err != nil {
		h.Log.Error("failed to create reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created reservation")
}

// ConfirmCheckInReservation godoc
//
//	@Summary		Confirm Check-In Reservation
//	@Description	Confirm check-in for a reservation by stay ID
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			stayId	path		string		true	"ID of the stay to confirm check-in"
//	@Param			request	body		reservation.ReservationCheckInEntity	true	"Check-in details"
//	@Success		200		{string}	string	"successfully confirmed reservation"
//	@Failure		400		{object}	response.ResponseError	"Error"
//	@Failure		401		{object}	response.ResponseError	"Unauthorized"
//	@Failure		500		{object}	response.ResponseError	"Error"
//	@Router			/reservation/checkin/{stayId} [post]
func (h *Handler) ConfirmCheckInReservation(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.ConfirmCheckInReservation"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stayId := chi.URLParam(r, "stayId")

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logged in")))
		return
	}

	var reserv reservation.ReservationCheckInEntity

	err := render.DecodeJSON(r.Body, &reserv)
	if err != nil {
		h.Log.Error("failed to decode JSON", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.ConfirmCheckInReservation(context.Background(), userID, stayId, reserv)
	if err != nil {
		h.Log.Error("failed to confirm reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully confirmed reservation")
}

// ConfirmCheckOutReservation godoc
//
//	@Summary		Confirm Check-Out Reservation
//	@Description	Confirm check-out for a reservation by stay ID
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			stayId	path		string		true	"ID of the stay to confirm check-out"
//	@Success		200		{string}	string	"successfully confirmed checkout reservation"
//	@Failure		400		{object}	response.ResponseError	"Error"
//	@Failure		401		{object}	response.ResponseError	"Unauthorized"
//	@Failure		500		{object}	response.ResponseError	"Error"
//	@Router			/reservation/checkout/{stayId} [post]
func (h *Handler) ConfirmCheckOutReservation(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.ConfirmCheckOutReservation"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stayId := chi.URLParam(r, "stayId")

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logged in")))
		return
	}

	err := h.Svc.ConfirmCheckOutReservation(context.Background(), userID, stayId)
	if err != nil {
		h.Log.Error("failed to confirm reservation", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully confirmed checkout reservation")
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
//	@Failure		400	{object}	response.ResponseError		"Invalid request"
//	@Failure		404	{object}	response.ResponseError		"Reservation not found"
//	@Failure		500	{object}	response.ResponseError		"Internal server error"
//	@Router			/reservation/{reservationId} [put]
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
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
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
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
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
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
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

// GetFreeReservationsByUserID godoc
//
//	@Summary		Get free reservations
//	@Description	Get free reservations by user ID
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			userID	path		string		true	"User ID"
//	@Success		200	{object}		[]stays.Stay	"List of free reservations"
//	@Failure		400	{object}		response.ResponseError	"Invalid user ID"
//	@Failure		500	{object}		response.ResponseError	"Failed to fetch reservations"
//	@Router			/reservation/free/{userID} [get]
func (h *Handler) GetFreeReservationsByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.GetFreeReservationsByUserID"

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

	freeReservations, err := h.Svc.GetFreeReservationsByUserID(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to get free reservations", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, freeReservations)
}

// GetOccupiedReservationsByUserID godoc
//
//	@Summary		Get occupied reservations
//	@Description	Get reservations by user ID where check-out is false
//	@Tags			reservations
//	@Accept			application/json
//	@Produce		json
//	@Param			userID	path		string		true	"user id"
//	@Success		200	{object}		[]stays.Stay	"ok"
//	@Failure		400	{object}		response.ResponseError	"Error"
//	@Failure		500	{object}		response.ResponseError	"Error"
//	@Router			/reservation/occupied/{userID} [get]
func (h *Handler) GetOccupiedReservationsByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.reservation.GetOccupiedReservationsByUserID"

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

	occupiedReservations, err := h.Svc.GetOccupiedReservationsByUserID(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to get occupied reservations", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, occupiedReservations)
}
