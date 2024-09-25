package staysadvantage

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysadvantage"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.StaysAdvantageService
	Log *slog.Logger
}

func (h *Handler) NewStaysAdvantageHandler(r chi.Router) {
	r.Route("/staysadvantage", func(r chi.Router) {

	})
}

// CreateStaysAdvantage godoc
//
//		@Summary		Create StaysAdvantage
//		@Description	Create staysadvantage
//		@Tags			staysadvantage
//		@Accept			application/json
//	 	@Param			stay_id	body	string			true	"stay id"
//	 	@Param			advantage_id	body	string			true	"advantage id"
//		@Produce		json
//		@Success		201	{object}		string	"created"
//		@Failure		400		{object}	responseApi.ResponseError			"Error"
//		@Failure		default		{object}	responseApi.ResponseError			"Error"
//		@Router			/staysadvantage/create [post]
func (h *Handler) CreateStaysAdvantage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysadvantage.CreateStaysAdvantage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req staysadvantage.StayAdvantageCreateReq

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.Log.Error("failed to parse request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateStaysAdvantage(r.Context(), &req)
	if err != nil {
		h.Log.Error("failed to create stay advantage", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created stay advantage")
}

// DeleteStaysAdvantageByID godoc
//
//	@Summary		Create StaysAdvantage
//	@Description	Create staysadvantage
//	@Tags			staysadvantage
//	@Accept			application/json
//	@Param			id	path		string		true	"stay advantage id"
//	@Produce		json
//	@Success		204	{object}		string	"no content"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/staysadvantage/{id} [delete]
func (h *Handler) DeleteStaysAdvantageByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysadvantage.DeleteStaysAdvantageByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteStaysAdvantageByID(context.Background(), uuID)
	if err != nil {
		h.Log.Error("failed to delete stay advantage", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "successfully deleted stay advantage")
}
