package staysreports

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	Svc interfaces.StaysReportsService
	Log *slog.Logger
}

func (h *Handler) NewStaysReportsHandler(r chi.Router) {
	r.Route("/report", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Get("/", h.GetAllStaysReports)
			r.Get("/{stayId}", h.GetStaysReportById)
			r.Put("/{reportId}", h.UpdateStaysReports)
			r.Delete("/{reportId}", h.DeleteStaysReports)
		})
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Post("/create/{stayId}", h.CreateStaysReports)
		})
	})
}

// CreateStaysReports creates a new stay report.
//
// @Summary      Create a new stay report
// @Description  Creates a report for a specific stay, including an optional image and required details like title and description.
// @Tags         staysReports
// @Accept       multipart/form-data
// @Produce      json
// @Param        stayId       path      string                true   "ID of the stay being reported"
// @Param        title        formData  string                true   "Title of the report"
// @Param        description  formData  string                true   "Description of the report"
// @Param        image        formData  file                  true  "image file (JPEG or PNG)"
// @Success      201          {object}  string   "Confirmation message"
// @Failure      400          {object}  map[string]string     "Error message for invalid input or unsupported image type"
// @Failure      401          {object}  map[string]string     "Error message for unauthorized access"
// @Failure      500          {object}  map[string]string     "Error message for internal server error"
// @Security     ApiKeyAuth
// @Router       /report/create/{stayId} [post]
func (h *Handler) CreateStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.CreateStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	// Getting userID from ctx
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}
	stayID := chi.URLParam(r, "stayId")

	// Restrict request body size
	r.Body = http.MaxBytesReader(w, r.Body, file.MaxImageMemorySize)

	// Parse multipart form
	err := r.ParseMultipartForm(file.MaxImageMemorySize)
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	// Extracting image file
	image, hdl, err := r.FormFile("image")
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}
	defer image.Close()

	// Validate content type
	contentType := hdl.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "image/jpeg") || strings.Contains(contentType, "image/png")) {
		h.Log.Error("unsupported content type", slogError.Err(handler.ErrInvalidImageType))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrInvalidImageType))
		return
	}

	// Read image into buffer
	buf := make([]byte, hdl.Size)
	n, err := image.Read(buf)
	if err != nil {
		h.Log.Error("failed to read image", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	// Parse JSON part of the body
	title := r.FormValue("title")
	description := r.FormValue("description")

	if title == "" || description == "" {
		h.Log.Error("title and description are required", slogError.Err(errors.New("missing fields")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	err = h.Svc.CreateStaysReports(context.Background(), userID, stayID, title, description, buf[:n])
	if err != nil {
		h.Log.Error("failed to create stays report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]string{"message": "Stay report created successfully"})
}

// GetAllStaysReports handles fetching all stay reports
// @Summary Get all stay reports
// @Description Retrieves all stay reports for the authenticated user
// @Tags staysReports
// @Produce json
// @Success 200 {array} staysreports.StaysReportEntity
// @Failure 401 {object} response.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} response.ResponseError "{"error": "message"}"
// @Router /report [get]
func (h *Handler) GetAllStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.GetAllStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	// Getting userID from ctx
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	reports, err := h.Svc.GetAllStaysReports(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to fetch reports", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not fetch history")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, reports)
}

// GetStaysReportById retrieves a stay report by user ID.
// @Summary Retrieve stay report by user ID
// @Description Fetches a specific stay report associated with the logged-in user.
// @Tags staysReports
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} staysreports.StayReport "Retrieved stay report object"
// @Failure 401 {object} response.ResponseError "{"error": "user not logged in"}"
// @Failure 404 {object} response.ResponseError "{"error": "report not found"}"
// @Failure 500 {object} response.ResponseError "{"error": "could not fetch report"}"
// @Router /report/{stayId} [get]
func (h *Handler) GetStaysReportById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.GetStaysReportById"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	stayReportId := chi.URLParam(r, "stayId")

	h.Log = h.Log.With(slog.String("user_id", userID))

	report, err := h.Svc.GetStaysReportById(r.Context(), userID, stayReportId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.Log.Warn("report not found", slogError.Err(err))
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(errors.New("report not found")))
		} else {
			h.Log.Error("failed to fetch report", slogError.Err(err))
			responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not fetch report")))
		}
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, report)
}

// UpdateStaysReports handles partially updating a stay report
// @Summary Partially update a stay report
// @Description Updates specific fields of a stay report, such as title, description, or image
// @Tags staysReports
// @Accept multipart/form-data
// @Produce json
// @Param reportId path string true "Report ID"
// @Param title formData string false "Updated title"
// @Param description formData string false "Updated description"
// @Param image formData file false "Image file (JPEG or PNG)"
// @Success 200 {object} staysreports.StaysReportEntity "Updated stays report object"
// @Failure 400 {object} response.ResponseError "Bad Request"
// @Failure 401 {object} response.ResponseError "Unauthorized"
// @Failure 500 {object} response.ResponseError "Internal Server Error"
// @Router /report/{reportId} [patch]
func (h *Handler) UpdateStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.UpdateStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())))

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, file.MaxImageMemorySize)

	if err := r.ParseMultipartForm(file.MaxImageMemorySize); err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var imageData []byte
	image, hdl, err := r.FormFile("image")
	if err == nil {
		defer image.Close()

		imgContentType := hdl.Header.Get("Content-Type")
		if !(strings.Contains(imgContentType, "image/jpeg") || strings.Contains(imgContentType, "image/png")) {
			h.Log.Error("unsupported content type", slogError.Err(handler.ErrInvalidImageType))
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrInvalidImageType))
			return
		}

		imageData, err = io.ReadAll(image)
		if err != nil {
			h.Log.Error("failed to read image", slogError.Err(err))
			responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
			return
		}
	} else if !errors.Is(err, http.ErrMissingFile) {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	reportID := chi.URLParam(r, "reportId")

	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		h.Log.Error("body params errors", slogError.Err(errors.New("title and description are required")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	report, err := h.Svc.UpdateStaysReports(context.Background(), userID, reportID, title, description, imageData)
	if err != nil {
		h.Log.Error("failed to update stays report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"Updated stay report": report})
}

// DeleteStaysReports handles deleting a stay report by report ID
// @Summary Delete a stay report
// @Description Deletes a specific stay report by report ID
// @Tags staysReports
// @Param reportId path string true "Report ID"
// @Success 200 {object} string "{"message": "Stay report was deleted"}"
// @Failure 401 {object} response.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} response.ResponseError "{"error": "message"}"
// @Router /report/{reportId} [delete]
func (h *Handler) DeleteStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.DeleteStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	reportID := chi.URLParam(r, "reportId")

	err := h.Svc.DeleteStaysReports(context.Background(), userID, reportID)
	if err != nil {
		h.Log.Error("failed to delete stay report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "Stay report was deleted"})
}
