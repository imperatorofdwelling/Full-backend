package image

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	createImage "github.com/imperatorofdwelling/Full-backend/internal/service/file"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"io/ioutil"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	Svc interfaces.FileService
	Log *slog.Logger
}

func (h *Handler) NewImageHandler(r chi.Router) {
	r.Route("/image", func(r chi.Router) {
		r.Post("/{category}", h.UploadImage)
		r.Delete("/{fileName}", h.RemoveImage)
	})
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// image size restrictions
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		h.Log.Error("failed to parse multipart form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, "Failed to parse form")
		return
	}

	category := chi.URLParam(r, "category")
	if category == "" {
		h.Log.Error("category is required")
		responseApi.WriteError(w, r, http.StatusBadRequest, "Category is required")
		return
	}

	file, hdl, err := r.FormFile("image")
	if err != nil {
		h.Log.Error("failed to get file from form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, "Failed to get file")
		return
	}
	defer file.Close()

	imgBytes, err := ioutil.ReadAll(file)
	if err != nil {
		h.Log.Error("failed to read file content", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, "Failed to read file content")
		return
	}

	var imgType string
	ext := hdl.Filename[strings.LastIndex(hdl.Filename, "."):]
	if ext == ".jpg" {
		imgType = ".jpg"
	} else if ext == ".png" {
		imgType = ".png"
	} else {
		h.Log.Error("unsupported file type", slog.String("filename", hdl.Filename))
		responseApi.WriteError(w, r, http.StatusBadRequest, "Unsupported file type")
		return
	}

	uploadedPath, err := h.Svc.UploadImage(imgBytes, createImage.ImageType(imgType), category)
	if err != nil {
		h.Log.Error("failed to upload an image", slogError.Err(err), slog.String("category", category), slog.String("filename", hdl.Filename))
		responseApi.WriteError(w, r, http.StatusInternalServerError, "Failed to save file")
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"image": uploadedPath})
}

func (h *Handler) RemoveImage(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "fileName")
	err := h.Svc.RemoveFile(fileName)
	if err != nil {
		h.Log.Error("failed to remove file", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "File removed successfully"})
}
