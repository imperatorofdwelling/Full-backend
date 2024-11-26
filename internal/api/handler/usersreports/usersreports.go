package usersreports

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
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
			r.Put("/{reportId}", h.UpdateUsersReports)
			r.Delete("/{reportId}", h.DeleteUsersReports)
		})
	})
}

// CreateUsersReports creates a report for a user
// @Summary Create User Report
// @Description Creates a new report on a user by another user
// @Tags UsersReports
// @Accept json
// @Produce json
// @Param toBlameId path string true "ID of the user being reported"
// @Param body body map[string]string true "Report content with title and description"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/report/create/{toBlameId} [post]
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

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(errors.Wrap(err, "decoding error")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	title, okTitle := reqBody["title"]
	description, okDescription := reqBody["description"]
	if !okTitle || !okDescription {
		h.Log.Error("body params errors", slogError.Err(errors.New("body errors")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	if err := h.Svc.CreateUsersReports(r.Context(), userID, toBlameID, title, description); err != nil {
		h.Log.Error("service failed to create user report", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]string{"message": "User report created successfully"})
}

// GetAllUsersReports retrieves all user reports
// @Summary Get All User Reports
// @Description Retrieves all reports created by a user
// @Tags UsersReports
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
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

// UpdateUsersReports updates a user report
// @Summary Update User Report
// @Description Updates an existing report for a specified user
// @Tags UsersReports
// @Accept json
// @Produce json
// @Param reportId path string true "ID of the report to update"
// @Param body body map[string]string true "Report content with title and description"
// @Success 200 {object} usersreports.UsersReportEntity "Updated user report object"
// @Failure 400 {object} responseApi.ResponseError "Invalid request"
// @Failure 401 {object} responseApi.ResponseError "Unauthorized"
// @Failure 500 {object} responseApi.ResponseError "Internal server error"
// @Router /user/report/{reportId} [put]
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
	toBlameID := chi.URLParam(r, "reportId")

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(errors.Wrap(err, "decoding error")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	title, okTitle := reqBody["title"]
	description, okDescription := reqBody["description"]
	if !okTitle || !okDescription {
		h.Log.Error("body params errors", slogError.Err(errors.New("body errors")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("title and description are required")))
		return
	}

	report, err := h.Svc.UpdateUsersReports(r.Context(), userID, toBlameID, title, description)
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
// @Tags UsersReports
// @Produce json
// @Param reportId path string true "ID of the report to delete"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
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
