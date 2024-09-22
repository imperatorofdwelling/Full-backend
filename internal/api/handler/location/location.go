package location

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	responseApi "github.com/imperatorofdwelling/Website-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Website-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.LocationService
	Log *slog.Logger
}

func (h *Handler) NewLocationHandler(r chi.Router) {
	r.Route("/locations", func(r chi.Router) {
		r.Get("/{locationName}", h.FindByNameMatch)
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
//	@Success		200	{object}		[]models.Location	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
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
		responseApi.WriteError(w, r, http.StatusInternalServerError, fmt.Errorf("%s: %s", op, err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, locations)
}
