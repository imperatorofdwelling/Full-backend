package advantage

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/advantage"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
	"strings"
)

const (
	MaxAdvantageImgSize    = 1 * (1024 * 1024)
	MaxAdvantageMemorySize = 2 * (1024 * 1024)
)

type Handler struct {
	Svc interfaces.AdvantageService
	Log *slog.Logger
}

func (h *Handler) NewAdvantageHandler(r chi.Router) {
	r.Route("/advantages", func(r chi.Router) {
		r.Post("/create", h.CreateAdvantage)
		r.Delete("/{advantageId}", h.RemoveAdvantage)
		r.Get("/all", h.GetAllAdvantages)
		r.Patch("/{advantageId}", h.UpdateAdvantage)
	})
}

// CreateAdvantage godoc
//
//		@Summary		Create Advantage
//		@Description	Create advantage
//		@Tags			advantages
//		@Accept			multipart/form-data
//	 	@Param			image	formData	file			true	"image file"
//	 	@Param			title	formData	string			true	"title of advantage"
//		@Produce		json
//		@Success		201	{string}		string	"created"
//		@Failure		400		{object}	responseApi.ResponseError			"Error"
//		@Failure		default		{object}	responseApi.ResponseError			"Error"
//		@Router			/advantages/create [post]
func (h *Handler) CreateAdvantage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.advantage.CreateAdvantage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	r.Body = http.MaxBytesReader(w, r.Body, MaxAdvantageMemorySize)

	err := r.ParseMultipartForm(MaxAdvantageMemorySize)
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	image, hdl, err := r.FormFile("image")
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}
	defer image.Close()

	if hdl.Size > MaxAdvantageImgSize || hdl.Size < 1 {
		h.Log.Error(handler.ErrInvalidImageSize.Error(), slogError.Err(handler.ErrInvalidImageSize))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrInvalidImageSize))
		return
	}

	contentType := hdl.Header.Get("Content-Type")

	if !strings.Contains(contentType, "image/svg+xml") {
		h.Log.Error("content type is not image/svg+xml", slogError.Err(handler.ErrImageTypeNotSvg))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrImageTypeNotSvg))
		return
	}

	title := r.FormValue("title")

	buf := make([]byte, hdl.Size)

	n, err := image.Read(buf)
	if err != nil {
		h.Log.Error("failed to read image", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	err = h.Svc.CreateAdvantage(context.Background(), &advantage.AdvantageEntity{Title: title, Image: buf[:n]})
	if err != nil {
		h.Log.Error("failed to create advantage", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, "successfully created advantage")
}

// RemoveAdvantage godoc
//
//	@Summary		Delete Advantage
//	@Description	Delete advantage by id
//	@Tags			advantages
//	@Accept			json
//	@Produce		json
//	@Param			advantageId	path		string		true	"advantage id"
//	@Success		204	{string}		string	"no content"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/advantages/{advantageId} [delete]
func (h *Handler) RemoveAdvantage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.advantage.RemoveAdvantage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	advId := chi.URLParam(r, "advantageId")

	uuidID, err := uuid.FromString(advId)
	if err != nil {
		h.Log.Error("failed to parse uuid", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.RemoveAdvantage(context.Background(), uuidID)
	if err != nil {
		h.Log.Error("failed to remove advantage", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "successfully removed advantage")
}

// GetAllAdvantages godoc
//
//	@Summary		Get advantages
//	@Description	Get all advantages
//	@Tags			advantages
//	@Accept			application/json
//	@Produce		json
//	@Success		200	{object}		[]advantage.Advantage	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/advantages/all [get]
func (h *Handler) GetAllAdvantages(w http.ResponseWriter, r *http.Request) {
	const op = "handler.advantage.GetAllAdvantage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	adv, err := h.Svc.GetAllAdvantages(context.Background())
	if err != nil {
		h.Log.Error("failed to get all advantages", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, adv)
}

// UpdateAdvantage godoc
//
//	@Summary		Update Advantage
//	@Description	Update advantage by id
//	@Tags			advantages
//	@Accept			application/json
//	@Produce		json
//	@Param			advantageId	path		string		true	"advantage id"
//	@Param			image	formData	file			false	"image file"
//	@Param			title	formData	string			false	"title of advantage"
//	@Success		200	{object}		advantage.Advantage	"ok"
//	@Failure		400		{object}	responseApi.ResponseError			"Error"
//	@Failure		default		{object}	responseApi.ResponseError			"Error"
//	@Router			/advantages/{advantageId} [patch]
func (h *Handler) UpdateAdvantage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.advantage.UpdateAdvantage"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	r.Body = http.MaxBytesReader(w, r.Body, MaxAdvantageMemorySize)

	err := r.ParseMultipartForm(MaxAdvantageMemorySize)
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	advId := chi.URLParam(r, "advantageId")

	uuidID, err := uuid.FromString(advId)
	if err != nil {
		h.Log.Error("failed to parse uuid", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var buf []byte
	var numBytes int

	image, hdl, err := r.FormFile("image")
	if err != nil {
		if !errors.Is(err, http.ErrMissingFile) {
			h.Log.Error("failed to parse form", slogError.Err(err))
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
			return
		}
	} else {
		defer image.Close()

		if hdl.Size > MaxAdvantageImgSize || hdl.Size < 1 {
			h.Log.Error(handler.ErrInvalidImageSize.Error(), slogError.Err(handler.ErrInvalidImageSize))
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrInvalidImageSize))
			return
		}

		imgContentType := hdl.Header.Get("Content-Type")

		if !strings.Contains(imgContentType, "image/svg+xml") {
			h.Log.Error("content type is not image/svg+xml", slogError.Err(handler.ErrImageTypeNotSvg))
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrImageTypeNotSvg))
			return
		}

		buf = make([]byte, hdl.Size)

		n, err := image.Read(buf)
		if err != nil {
			h.Log.Error("failed to read image", slogError.Err(err))
			responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
			return
		}

		numBytes = n
	}

	title := r.FormValue("title")

	advUpdated, err := h.Svc.UpdateAdvantageByID(context.Background(), uuidID, &advantage.AdvantageEntity{Title: title, Image: buf[:numBytes]})
	if err != nil {
		h.Log.Error("failed to update advantage", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, advUpdated)
}
