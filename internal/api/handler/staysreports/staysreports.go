package staysreports

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.StaysReportsService
	Log *slog.Logger
}

func (h *Handler) NewStaysReportsHandler(r chi.Router) {
	r.Route("/report", func(r chi.Router) {
		r.Post("/create/{stayId}", h.CreateStaysReports)
		r.Get("/", h.GetAllStaysReports)
		r.Put("/{reportId}", h.UpdateStaysReports)
		r.Delete("/{reportId}", h.DeleteStaysReports)
	})
}

// CreateStaysReports handles the creation of a stay report
// @Summary Create a stay report
// @Description Creates a new stay report for a specific stay
// @Tags stays-reports
// @Accept json
// @Produce json
// @Param stayId path string true "Stay ID"
// @Param body body map[string]string true "Report data"
// @Success 201 {object} map[string]string "{"message": "Stay report created successfully"}"
// @Failure 400 {object} responseApi.ResponseError "{"error": "message"}"
// @Failure 401 {object} responseApi.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} responseApi.ResponseError "{"error": "message"}"
// @Router /report/create/{stayId} [post]
func (h *Handler) CreateStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.CreateStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	// Getting userID from ctx
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", nil)
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}
	stayID := chi.URLParam(r, "stayId")

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "failed to decode request body")))
		return
	}

	title, okTitle := reqBody["title"]
	description, okDescription := reqBody["description"]
	if !okTitle || !okDescription {
		h.Log.Error("missing required fields in request body", nil)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	err := h.Svc.CreateStaysReports(context.Background(), userID, stayID, title, description)
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
// @Tags stays-reports
// @Produce json
// @Success 200 {array} staysreports.StaysReportEntity
// @Failure 401 {object} responseApi.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} responseApi.ResponseError "{"error": "message"}"
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
		h.Log.Error("user not logged in", nil)
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

// UpdateStaysReports handles updating a stay report
// @Summary Update a stay report
// @Description Updates a specific stay report
// @Tags stays-reports
// @Accept json
// @Produce json
// @Param reportId path string true "Report ID"
// @Param body body map[string]string true "Updated report data"
// @Success 200 {object} map[string]string "{"message": "Stay report updated successfully"}"
// @Failure 400 {object} responseApi.ResponseError "{"error": "message"}"
// @Failure 401 {object} responseApi.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} responseApi.ResponseError "{"error": "message"}"
// @Router /report/{reportId} [put]
func (h *Handler) UpdateStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.UpdateStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", nil)
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}
	reportID := chi.URLParam(r, "reportId")

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "failed to decode request body")))
		return
	}

	title, okTitle := reqBody["title"]
	description, okDescription := reqBody["description"]
	if !okTitle || !okDescription {
		h.Log.Error("missing required fields in request body", nil)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	err := h.Svc.UpdateStaysReports(context.Background(), userID, reportID, title, description)
	if err != nil {
		h.Log.Error("failed to update stays report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "Stay report updated successfully"})
}

// DeleteStaysReports handles deleting a stay report by report ID
// @Summary Delete a stay report
// @Description Deletes a specific stay report by report ID
// @Tags stays-reports
// @Param reportId path string true "Report ID"
// @Success 200 {object} map[string]string "{"message": "Stay report was deleted"}"
// @Failure 401 {object} responseApi.ResponseError "{"error": "user not logged in"}"
// @Failure 500 {object} responseApi.ResponseError "{"error": "message"}"
// @Router /report/{reportId} [delete]
func (h *Handler) DeleteStaysReports(w http.ResponseWriter, r *http.Request) {
	const op = "handler.StaysReports.DeleteStaysReports"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user not logged in", nil)
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
