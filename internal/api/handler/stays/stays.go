package stays

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

const (
	MaxImageMemorySize = 5 * (1024 * 1024)
)

type Handler struct {
	Svc interfaces.StaysService
	Log *slog.Logger
}

func (h *Handler) NewStaysHandler(r chi.Router) {
	r.Route("/stays", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Post("/create", h.CreateStay)
			r.Get("/{stayId}", h.GetStayByID)
			r.Delete("/{stayId}", h.DeleteStayByID)
			r.Put("/{stayId}", h.UpdateStayByID)
			r.Post("/images", h.CreateImages)
			r.Post("/images/main", h.CreateMainImage)
			r.Delete("/images/delete/{imageId}", h.DeleteStayImage)
		})

		r.Group(func(r chi.Router) {
			r.Get("/", h.GetStays)
			r.Get("/user/{userId}", h.GetStaysByUserID)
			r.Get("/images/{stayId}", h.GetStayImagesByStayID)
			r.Get("/images/main/{stayId}", h.GetMainImageByStayID)
		})
	})
}

// CreateStay godoc
//
//	@Summary		Create Stay
//	@Description	Create stay
//	@Tags			stays
//	@Accept			application/json
//	@Produce		json
//	@Param	request body model.StayEntity	true	"request stay data"
//	@Success		201	{string}		string		"created"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/create [post]
func (h *Handler) CreateStay(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.CreateStay"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var newStay model.StayEntity

	err := render.DecodeJSON(r.Body, &newStay)
	if err != nil {
		h.Log.Error("failed to decode form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateStay(r.Context(), &newStay)
	if err != nil {
		h.Log.Error("failed to create stay: ", slogError.Err(err))
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
//	@Success		200	{object}		model.Stay		"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/{stayId} [get]
func (h *Handler) GetStayByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStayByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "stayId")
	idUuid, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("%s: %v", op, slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	stay, err := h.Svc.GetStayByID(r.Context(), idUuid)
	if err != nil {
		h.Log.Error("failed to fetch stay by id %s: %v", slogError.Err(err))
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
//	@Success		200	{object}		[]model.Stay	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays [get]
func (h *Handler) GetStays(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStays"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stays, err := h.Svc.GetStays(r.Context())
	if err != nil {
		h.Log.Error("failed to fetch stays: ", slogError.Err(err))
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
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteStayByID(r.Context(), idUuid)
	if err != nil {
		h.Log.Error("failed to delete stay by id %s: %v", slogError.Err(err))
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
//	@Accept			application/json
//	@Produce		json
//	@Param	request body model.StayEntity	true	"request stay data"
//	@Success		200	{object}		model.Stay	"ok"
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
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var newStay model.StayEntity

	err = render.DecodeJSON(r.Body, &newStay)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	updatedStay, err := h.Svc.UpdateStayByID(r.Context(), &newStay, idUuid)
	if err != nil {
		h.Log.Error("failed to update stay by id %s: %v", slogError.Err(err))
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
//	@Success		200	{object}		[]model.Stay	"ok"
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
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	stays, err := h.Svc.GetStaysByUserID(r.Context(), idUuid)
	if err != nil {
		h.Log.Error("failed to fetch stays: %v", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stays)
}

// GetStayImagesByStayID godoc
//
//	@Summary		Get all stays images by stay id
//	@Description	Get stays images by stay id
//	@Tags			stays
//	@Accept			application/json
//	@Param			stayId	path		string		true	"stay id"
//	@Produce		json
//	@Success		200	{object}		[]model.StayImage	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/images/{stayId} [get]
func (h *Handler) GetStayImagesByStayID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetStayImagesByStayID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stayId := chi.URLParam(r, "stayId")
	idUuid, err := uuid.FromString(stayId)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	stayImages, err := h.Svc.GetImagesByStayID(r.Context(), idUuid)
	if err != nil {
		h.Log.Error("failed to fetch stay images by id %s: %v", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stayImages)
}

// GetMainImageByStayID godoc
//
//	@Summary		Get main stays image by stay id
//	@Description	Get main stays image by stay id
//	@Tags			stays
//	@Accept			application/json
//	@Param			stayId	path		string		true	"stay id"
//	@Produce		json
//	@Success		200	{object}		model.StayImage	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/images/main/{stayId} [get]
func (h *Handler) GetMainImageByStayID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.GetMainImageByStayID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	stayId := chi.URLParam(r, "stayId")
	idUuid, err := uuid.FromString(stayId)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	stayImage, err := h.Svc.GetMainImageByStayID(r.Context(), idUuid)
	if err != nil {
		h.Log.Error("failed to fetch stay image by id %s: %v", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, stayImage)
}

// CreateImages godoc
//
//	@Summary		Create images
//	@Description	Create images
//	@Tags			stays
//	@Accept			multipart/form-data
//	@Param			images	formData		[]file		true	"images"
//	@Param			stay_id	formData		string		true	"stay id"
//	@Produce		json
//	@Success		200	{object}		string	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/images [post]
func (h *Handler) CreateImages(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.CreateImages"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	err := r.ParseMultipartForm(MaxImageMemorySize)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	formValues := r.MultipartForm.Value

	formFiles := r.MultipartForm.File

	images := formFiles["images"]
	stayID := formValues["stay_id"][0]

	stayIDUuid, err := uuid.FromString(stayID)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateImages(r.Context(), images, stayIDUuid)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created")
}

// CreateMainImage godoc
//
//	@Summary		Create main image
//	@Description	Create main image
//	@Tags			stays
//	@Accept			multipart/form-data
//	@Param			images	formData		file		true	"images"
//	@Param			stay_id	formData		string		true	"stay id"
//	@Produce		json
//	@Success		200	{object}		string	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/images/main [post]
func (h *Handler) CreateMainImage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.CreateMainImage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	err := r.ParseMultipartForm(MaxImageMemorySize)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	formValues := r.MultipartForm.Value
	formFiles := r.MultipartForm.File

	mainImage := formFiles["images"][0]
	stayID := formValues["stay_id"][0]

	stayIDUuid, err := uuid.FromString(stayID)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.CreateMainImage(r.Context(), mainImage, stayIDUuid)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created")
}

// DeleteStayImage godoc
//
//	@Summary		Delete image by id
//	@Description	Delete image by id
//	@Tags			stays
//	@Accept			application/json
//	@Param			imageId	path		string		true	"stay image id"
//	@Produce		json
//	@Success		200	{object}		string	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/stays/images/delete/{imageId} [delete]
func (h *Handler) DeleteStayImage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.stays.DeleteStayImage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	imageID := chi.URLParam(r, "imageId")
	imageIDUuid, err := uuid.FromString(imageID)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.DeleteStayImage(r.Context(), imageIDUuid)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "successfully deleted")

}
