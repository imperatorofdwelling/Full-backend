package usersreports

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type Handler struct {
	Svc interfaces.UsersReportsService
	Log *slog.Logger
}

func (h *Handler) NewUsersReportsHandler(r chi.Router) {
	r.Route("/user/report", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Post("/create/{toBlameId}", h.CreateUsersReports)
			r.Get("/", h.GetAllUsersReports)
			r.Get("/{reportId}", h.GetUsersReportById)
			r.Patch("/{reportId}", h.UpdateUsersReports)
			r.Delete("/{reportId}", h.DeleteUsersReports)
		})
	})
}

// CreateUsersReports creates a new user report.
//
// @Summary      Create a new user report
// @Description  Creates a user report with an optional image and necessary details such as title and description.
// @Tags         usersReports
// @Accept       multipart/form-data
// @Produce      json
// @Param        toBlameId  path      string                true   "ID of the user being reported"
// @Param        title      formData  string                true   "Title of the report"
// @Param        description formData string                true   "Description of the report"
// @Param        image      formData  file                  true  "Optional image file (JPEG or PNG)"
// @Success      201        {object}  string    "Message confirming successful creation"
// @Failure      400        {object}  response.ResponseError    "Invalid input, missing fields, or unsupported image type"
// @Failure      401        {object}  response.ResponseError    "Unauthorized, user not logged in"
// @Failure      500        {object}  response.ResponseError    "Internal server error"
// @Router       /user/report/{toBlameId} [post]
func (h *Handler) CreateUsersReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UsersReports.CreateUsersReports"

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
	toBlameID := chi.URLParam(r, "toBlameId")

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

	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		h.Log.Error("title and description are required", slogError.Err(errors.New("missing fields")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	if err = h.Svc.CreateUsersReports(r.Context(), userID, toBlameID, title, description, buf[:n]); err != nil {
		h.Log.Error("service failed to create user report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]string{"message": "User report created successfully"})
}

// GetAllUsersReports retrieves all user reports
// @Summary Get All User Reports
// @Description Retrieves all reports created by a user
// @Tags usersReports
// @Produce json
// @Success 200 {array} usersreports.UsersReportEntity
// @Failure 401 {object} response.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} response.ResponseError "{"error": "could not get reports"}"
// @Router /user/report/ [get]
func (h *Handler) GetAllUsersReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UsersReports.GetAllUsersReports"

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

	reports, err := h.Svc.GetAllUsersReports(r.Context(), userID)
	if err != nil {
		h.Log.Error("service failed to fetch user reports", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "failed to fetch user reports")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, reports)
}

// GetUsersReportById handles fetching a user's report by its ID.
//
// @Summary      Get a user's report by ID
// @Description  Retrieve detailed information about a specific user's report by its unique ID.
// @Tags         usersReports
// @Accept       json
// @Produce      json
// @Param        reportId   path      string  true  "Report ID"
// @Success      200        {object} usersreports.UsersReportEntity "Successful response containing user report data"
// @Failure      401        {object} response.ResponseError  "Unauthorized: User not logged in"
// @Failure      500        {object} response.ResponseError  "Internal Server Error"
// @Security     ApiKeyAuth
// @Router       /users/report/{reportId} [get]
func (h *Handler) GetUsersReportById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UsersReports.GetUsersReportById"

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

	reportId := chi.URLParam(r, "reportId")

	reports, err := h.Svc.GetUsersReportById(r.Context(), userID, reportId)
	if err != nil {
		h.Log.Error("service failed to fetch user report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "failed to fetch user reports")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, reports)
}

// UpdateUsersReports updates a user report
// @Summary Update User Report
// @Description Updates specific fields (title, description, or image) of an existing user report
// @Tags usersReports
// @Accept multipart/form-data
// @Produce json
// @Param reportId path string true "ID of the report to update"
// @Param title formData string false "New title for the report"
// @Param description formData string false "New description for the report"
// @Param image formData file true "New image file (JPEG or PNG)"
// @Success 200 {object} usersreports.UsersReportEntity "Updated user report object"
// @Failure 400 {object} response.ResponseError "Invalid request"
// @Failure 401 {object} response.ResponseError "Unauthorized"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/report/{reportId} [patch]
func (h *Handler) UpdateUsersReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UsersReports.UpdateUsersReports"

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
	reportId := chi.URLParam(r, "reportId")

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

	title := r.FormValue("title")
	description := r.FormValue("description")
	if title == "" || description == "" {
		h.Log.Error("body params errors", slogError.Err(errors.New("title and description are required")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	report, err := h.Svc.UpdateUsersReports(r.Context(), userID, reportId, title, description, imageData)
	if err != nil {
		h.Log.Error("service failed to update user report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"Updated user report": report})
}

// DeleteUsersReports deletes a user report
// @Summary Delete User Report
// @Description Deletes a specified report created by a user
// @Tags usersReports
// @Produce json
// @Param reportId path string true "ID of the report to delete"
// @Success 200 {object} string "User report was deleted"
// @Failure 401 {object} response.ResponseError "Unauthorized"
// @Failure 500 {object} response.ResponseError
// @Router /user/report/{reportId} [delete]
func (h *Handler) DeleteUsersReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.UsersReports.DeleteUsersReports"

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

	if err := h.Svc.DeleteUsersReports(r.Context(), userID, reportID); err != nil {
		h.Log.Error("service failed to delete user report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "User report was deleted"})
}
