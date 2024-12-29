package searchhistory

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/searchhistory"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"

	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.SearchHistoryService
	Log *slog.Logger
}

func (h *Handler) NewHistorySearchHandler(r chi.Router) {
	r.Route("/history", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Get("/", h.GetAllHistoryByUserId)
			r.Post("/", h.AddHistory)
		})
	})
}

// GetAllHistoryByUserId godoc
//
//	@Summary		Get Search History
//	@Description	Get all search history for a user by ID
//	@Tags			searchHistory
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		searchhistory.SearchHistory	"ok"
//	@Failure		401	{object}	responseApi.ResponseError	"Error"
//	@Failure		500	{object}	responseApi.ResponseError	"Error"
//	@Router			/history [get]
func (h *Handler) GetAllHistoryByUserId(w http.ResponseWriter, r *http.Request) {
	const op = "handler.searchHistory.GetAllHistoryByUserId"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	hist, err := h.Svc.GetAllHistoryByUserId(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to fetch history", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not fetch history")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, hist)
}

// AddHistory godoc
//
//	@Summary		Create Search History
//	@Description	Add a new entry to the search history for a user
//	@Tags			searchHistory
//	@Accept			json
//	@Produce		json
//	@Param			name	body		string	true	"Name of the search history entry"
//	@Success		201	{object}	string	"message"
//	@Failure		401	{object}	responseApi.ResponseError	"Error"
//	@Failure		400	{object}	responseApi.ResponseError	"Error"
//	@Failure		500	{object}	responseApi.ResponseError	"Error"
//	@Router			/history [post]
func (h *Handler) AddHistory(w http.ResponseWriter, r *http.Request) {
	const op = "handler.searchHistory.AddHistory"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.Wrap(err, "failed to decode request body")))
		return
	}

	name, exists := reqBody["name"]
	if !exists || name == "" {
		h.Log.Error("name parameter is required")
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("name parameter is required")))
		return
	}

	err := h.Svc.AddHistory(context.Background(), userID, name)
	if err != nil {
		h.Log.Error("could not add history", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not add history")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]string{"message": "History created successfully"})
}
