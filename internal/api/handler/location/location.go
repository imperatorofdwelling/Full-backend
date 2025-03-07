package location

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.LocationService
	Log *slog.Logger
}

func (h *Handler) NewLocationHandler(r chi.Router) {
	r.Route("/locations", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Get("/{id}", h.GetOneByID)
			r.Delete("/{id}", h.DeleteByID)
			r.Put("/{id}", h.UpdateByID)
		})

		r.Group(func(r chi.Router) {
			r.Get("/{locationName}", h.FindByNameMatch)
			r.Get("/", h.GetAll)
		})
	})
}

// FindByNameMatch godoc
//
//	@Summary		Find city by name
//	@Description	Find city by matching name
//	@Tags			locations
//	@Accept			json
//	@Produce		json
//	@Param			locationName	path		string		true	"location name match"
//	@Success		200	{object}		[]string	"ok"
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
//	@Router			/locations/{locationName} [get]
func (h *Handler) FindByNameMatch(w http.ResponseWriter, r *http.Request) {
	const op = "handler.location.FindByNameMatch"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	locationName := chi.URLParam(r, "locationName")

	locations, err := h.Svc.FindByNameMatch(context.Background(), locationName)
	if err != nil {
		h.Log.Error("failed to find location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, locations)
}

// GetAll godoc
//
//	@Summary		Get all locations
//	@Description	Get all locations
//	@Tags			locations
//	@Accept			application/json
//	@Produce		json
//	@Success		200	{object}		[]location.Location	"ok"
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
//	@Router			/locations [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	const op = "handler.location.GetAll"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	locations, err := h.Svc.GetAll(r.Context())
	if err != nil {
		h.Log.Error("failed to find locations", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, locations)
}

// GetOneByID godoc
//
//	@Summary		Find location by id
//	@Description	Find location by id
//	@Tags			locations
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string		true	"location id"
//	@Success		200	{object}		location.Location	"ok"
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
//	@Router			/locations/{id} [get]
func (h *Handler) GetOneByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.location.GetOneByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	idUUID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to find location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := h.Svc.GetByID(r.Context(), idUUID)
	if err != nil {
		h.Log.Error("failed to find location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, result)
}

// DeleteByID godoc
//
//	@Summary		Delete location by id
//	@Description	Delete location by id
//	@Tags			locations
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string		true	"location id"
//	@Success		200	{object}		string	"ok"
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
//	@Router			/locations/{id} [delete]
func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.location.DeleteByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	idUUID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to find location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteByID(r.Context(), idUUID)
	if err != nil {
		h.Log.Error("failed to delete location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "successfully deleted")
}

// UpdateByID godoc
//
//	@Summary		Update location by id
//	@Description	Update location by id
//	@Tags			locations
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string		true	"location id"
//	@Param			request	body		location.LocationEntity		true	"location request"
//	@Success		200	{object}		string	"ok"
//	@Failure		400		{object}	response.ResponseError			"Error"
//	@Failure		default		{object}	response.ResponseError			"Error"
//	@Router			/locations/{id} [put]
func (h *Handler) UpdateByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.location.UpdateByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	idUUID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to find location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var loc location.LocationEntity

	err = render.DecodeJSON(r.Body, &loc)
	if err != nil {
		h.Log.Error("failed to decode body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.UpdateByID(r.Context(), idUUID, loc)
	if err != nil {
		h.Log.Error("failed to update location", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully updated")
}
