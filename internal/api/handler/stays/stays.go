package stays

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
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

	var newStay stays.StayEntity

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
// @Success		200	{object}		models.Stay		"ok"
// @Failure		400		{object}	responseApi.ResponseError			"Error"
// @Failure		default		{object}	responseApi.ResponseError			"Error"
// @Router			/stays/{stayId} [get]
func (h *Handler) GetStayByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStayByID"

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
//	@Success		200	{object}		[]models.Stay	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays [get]
func (h *Handler) GetStays(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStays"

	stays, err := h.Svc.GetStays(context.Background())
	if err != nil {
		h.Log.Error("failed to fetch stays: ", err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stays)
}
