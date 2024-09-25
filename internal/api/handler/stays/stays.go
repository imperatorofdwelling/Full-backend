package stays

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.StaysService
	Log *slog.Logger
}

func (h *Handler) NewStaysHandler(r chi.Router) {
	r.Route("/stays", func(r chi.Router) {
		r.Post("/create", h.CreateStay)
		r.Get("/{stayId}", h.GetStayByID)
		r.Get("/", h.GetStays)
		r.Delete("/{stayId}", h.DeleteStayByID)
		r.Put("/{stayId}", h.UpdateStayByID)
		r.Get("/user/{userId}", h.GetStaysByUserID)
	})
}

// CreateStay godoc
//
//		@Summary		Create Stay
//		@Description	Create stay
//		@Tags			stays
//		@Accept			application/json
//		@Produce		json
//	 	@Param			user_id	formData		string			true	"id of user"
//	 	@Param			location_id	formData	string			true	"id of location"
//	 	@Param			name	formData		string			false	"name of stay"
//	 	@Param			image_main	formData	file			false	"main image"
//	 	@Param			images	formData	file			false	"images"
//	 	@Param			type	formData		string			false	"type of stay"
//	 	@Param			number_of_bedrooms	formData		int			false	"number of bedrooms"
//	 	@Param			number_of_beds	formData		int			false	"number of bathrooms"
//	 	@Param			number_of_bathrooms	formData		int			false	"number of beds"
//	 	@Param			guests	formData		int			false	"number of guests"
//	 	@Param			rating	formData		float32			false	"rating"
//	 	@Param			is_smoking_prohibited	formData		boolean			false	"smoking"
//	 	@Param			square	formData		float32			false	"square of home"
//	 	@Param			street	formData		string			true	"street"
//	 	@Param			house	formData		string			true	"house"
//	 	@Param			entrance	formData		string			false	"entrance if exists"
//	 	@Param			floor	formData		string			false	"floor if exists"
//	 	@Param			room	formData		string			false	"room if exists"
//	 	@Param			price	formData		float32			false	"price of stay"
//		@Success		201	{string}		string		"created"
//		@Failure		400		{object}	responseApi.ResponseError			"Error"
//		@Failure		default		{object}	responseApi.ResponseError			"Error"
//		@Router			/stays/create [post]
func (h *Handler) CreateStay(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.CreateStay"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var newStay models.StayEntity

	err := render.DecodeJSON(r.Body, &newStay)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateStay(context.Background(), &newStay)
	if err != nil {
		h.Log.Error("failed to create stay: ", err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created")
}

// GetStayByID godoc
//
//	@Summary		Get Stay by id
//	@Description	get stay
//	@Tags			stays
//	@Accept			application/json
//	@Produce		json
//	@Param			stayId	path		string		true	"stay id"
//
// @Success		200	{object}		stays.Stay		"ok"
// @Failure		400		{object}	responseApi.ResponseError			"Error"
// @Failure		default		{object}	responseApi.ResponseError			"Error"
// @Router			/stays/{stayId} [get]
func (h *Handler) GetStayByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStayByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "stayId")
	idUuid, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	stay, err := h.Svc.GetStayByID(context.Background(), idUuid)
	if err != nil {
		h.Log.Error("failed to fetch stay by id %s: %v", idUuid, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stay)
}

// GetStays godoc
//
//	@Summary		Get all stays
//	@Description	Get all stays
//	@Tags			stays
//	@Accept			application/json
//	@Produce		json
//	@Success		200	{object}		[]stays.Stay	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays [get]
func (h *Handler) GetStays(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStays"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stays, err := h.Svc.GetStays(context.Background())
	if err != nil {
		h.Log.Error("failed to fetch stays: ", err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stays)
}

// DeleteStayByID godoc
//
//	@Summary		Delete Stay
//	@Description	Delete stay by id
//	@Tags			stays
//	@Accept			json
//	@Produce		json
//	@Param			stayId	path		string		true	"stay id"
//	@Success		204	{string}		string	"no content"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/{stayId} [delete]
func (h *Handler) DeleteStayByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.DeleteStay"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stayId := chi.URLParam(r, "stayId")
	idUuid, err := uuid.FromString(stayId)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteStayByID(context.Background(), idUuid)
	if err != nil {
		h.Log.Error("failed to delete stay by id %s: %v", idUuid, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "successfully deleted")
}

// UpdateStayByID godoc
//
//	@Summary		Update Stay
//	@Description	Update stay by id
//	@Tags			stays
//	@Accept			json
//	@Produce		json
//	@Param			stayId	path		string		true	"advantage id"
//	@Param			location_id	formData	string			true	"id of location"
//	@Param			name	formData		string			false	"name of stay"
//	@Param			image_main	formData	file			false	"main image"
//	@Param			images	formData	file			false	"images"
//	@Param			type	formData		string			false	"type of stay"
//	@Param			number_of_bedrooms	formData		int			false	"number of bedrooms"
//	@Param			number_of_beds	formData		int			false	"number of bathrooms"
//	@Param			number_of_bathrooms	formData		int			false	"number of beds"
//	@Param			guests	formData		int			false	"number of guests"
//	@Param			rating	formData		float32			false	"rating"
//	@Param			is_smoking_prohibited	formData		boolean			false	"smoking"
//	@Param			square	formData		float32			false	"square of home"
//	@Param			street	formData		string			true	"street"
//	@Param			house	formData		string			true	"house"
//	@Param			entrance	formData		string			false	"entrance if exists"
//	@Param			floor	formData		string			false	"floor if exists"
//	@Param			room	formData		string			false	"room if exists"
//	@Param			price	formData		float32			false	"price of stay"
//	@Success		200	{object}		stays.Stay	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/{stayId} [put]
func (h *Handler) UpdateStayByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.UpdateStayByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stayId := chi.URLParam(r, "stayId")
	idUuid, err := uuid.FromString(stayId)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	var newStay models.StayEntity

	err = render.DecodeJSON(r.Body, &newStay)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	updatedStay, err := h.Svc.UpdateStayByID(context.Background(), &newStay, idUuid)
	if err != nil {
		h.Log.Error("failed to update stay by id %s: %v", idUuid, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, updatedStay)
}

// GetStaysByUserID godoc
//
//	@Summary		Get all stays by user id
//	@Description	Get stays by user id
//	@Tags			stays
//	@Accept			application/json
//	@Param			userId	path		string		true	"user id"
//	@Produce		json
//	@Success		200	{object}		[]stays.Stay	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/user/{userId} [get]
func (h *Handler) GetStaysByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStaysByUserID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId := chi.URLParam(r, "userId")
	idUuid, err := uuid.FromString(userId)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	stays, err := h.Svc.GetStaysByUserID(context.Background(), idUuid)
	if err != nil {
		h.Log.Error("failed to fetch stays: %v", err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stays)
}
