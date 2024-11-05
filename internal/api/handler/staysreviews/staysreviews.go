package staysreviews

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreviews"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.StaysReviewsService
	Log *slog.Logger
}

func (h *Handler) NewStaysReviewsHandler(r chi.Router) {
	r.Route("/staysreviews", func(r chi.Router) {
		r.Post("/create", h.CreateStaysReview)
		r.Put("/update/{id}", h.UpdateStaysReview)
		r.Delete("/{id}", h.DeleteStaysReview)
		r.Get("/{id}", h.FindOneStaysReview)
		r.Get("/", h.FindAllStaysReviews)
	})
}

// CreateStaysReview godoc
//
//		@Summary		Create Stays_review
//		@Description	Create stays_review
//		@Tags			staysReviews
//		@Accept			application/json
//		@Produce		json
//	 	@Param			request	body	model.StaysReviewEntity			true	"stays review request"
//		@Success		201	{string}		string	"created"
//		@Failure		400		{object}	responseApi.ResponseError			"Error"
//		@Failure		default		{object}	responseApi.ResponseError			"Error"
//		@Router			/staysreviews/create [post]
func (h *Handler) CreateStaysReview(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysreviews.CreateStaysReview"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var newStaysReview model.StaysReviewEntity

	err := render.DecodeJSON(r.Body, &newStaysReview)
	if err != nil {
		h.Log.Error("failed to decode body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateStaysReview(r.Context(), &newStaysReview)
	if err != nil {
		h.Log.Error("failed to create stay review", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "stay review created")
}

// UpdateStaysReview godoc
//
//	@Summary		Update Stays Review
//	@Description	Update a stays review by its ID
//	@Tags			staysReviews
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string		true	"ID of the stays review to update"
//	@Param			request	body	model.StaysReviewEntity	true	"Details to update the stays review"
//	@Success		200	{object}	map[string]interface{}	"Successfully updated stays review"
//	@Failure		400	{object}	responseApi.ResponseError		"Invalid request"
//	@Failure		404	{object}	responseApi.ResponseError		"Stays review not found"
//	@Failure		500	{object}	responseApi.ResponseError		"Internal server error"
//	@Router			/staysreviews/update/{id} [put]
func (h *Handler) UpdateStaysReview(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysreviews.UpdateStaysReview"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var newStaysReview model.StaysReviewEntity

	err = render.DecodeJSON(r.Body, &newStaysReview)
	if err != nil {
		h.Log.Error("failed to decode body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	stayRev, err := h.Svc.UpdateStaysReview(r.Context(), &newStaysReview, uuID)
	if err != nil {
		h.Log.Error("failed to update stay review", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"message": "Updated stay review", "review": stayRev})
}

// DeleteStaysReview godoc
//
//	@Summary		Delete Stays_review
//	@Description	Delete Stays_review by id
//	@Tags			staysReviews
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string		true	"stays review id"
//	@Success		200	{string}		string	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/staysreviews/{id} [delete]
func (h *Handler) DeleteStaysReview(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysreviews.DeleteStaysReview"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteStaysReview(r.Context(), uuID)
	if err != nil {
		h.Log.Error("failed to delete stay review", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "stay review deleted")
}

// FindOneStaysReview godoc
//
//	@Summary		Get Stays review
//	@Description	Get Stays review by id
//	@Tags			staysReviews
//	@Accept			application/json
//	@Produce		json
//	@Param			id	path		string		true	"stays review id"
//	@Success		200	{object}		model.StaysReview	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/staysreviews/{id} [get]
func (h *Handler) FindOneStaysReview(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysreviews.FindOneStaysReview"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "id")
	uuID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	foundStayReview, err := h.Svc.FindOneStaysReview(r.Context(), uuID)
	if err != nil {
		h.Log.Error("failed to find stay review", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, foundStayReview)
}

// FindAllStaysReviews godoc
//
//	@Summary		Get all Stays review
//	@Description	Get all Stays reviews
//	@Tags			staysReviews
//	@Accept			application/json
//	@Produce		json
//	@Success		200	{object}		[]model.StaysReview	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/staysreviews [get]
func (h *Handler) FindAllStaysReviews(w http.ResponseWriter, r *http.Request) {
	const op = "handler.staysreviews.FindAllStaysReviews"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	foundStayReviews, err := h.Svc.FindAllStaysReviews(r.Context())
	if err != nil {
		h.Log.Error("failed to find all stay reviews", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, foundStayReviews)
}
