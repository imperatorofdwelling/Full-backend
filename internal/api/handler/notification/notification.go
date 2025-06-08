package notification

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/notification"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.NotificationService
	Log *slog.Logger
}

func (h *Handler) NewNotificationHandler(r chi.Router) {
	r.Route("/notification", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Get("/token/set", h.SetDeviceToken)
		})
	})
}

func (h *Handler) SetDeviceToken(w http.ResponseWriter, r *http.Request) {
	const op = "handler.notification.SetDeviceToken"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	var deviceInfo notification.DeviceInfo
	if err := json.NewDecoder(r.Body).Decode(&deviceInfo); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}
	defer r.Body.Close()

	// Базовая валидация токена
	if deviceInfo.Token == "" {
		h.Log.Error("empty device token")
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("device token cannot be empty")))
		return
	}

	updatedToken, err := h.Svc.SetDeviceTokenToUser(context.Background(), deviceInfo, userID)
	if err != nil {
		h.Log.Error("failed to get messages by user id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, updatedToken)
}

func (h *Handler) SendNotification(w http.ResponseWriter, r *http.Request) {
	const op = "handler.notification.SendNotification"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	var request struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.Log.Error("failed to decode request", "error", err)
		responseApi.WriteError(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.Svc.SendNotificationToUser(r.Context(), userID, request.Message); err != nil {
		h.Log.Error("failed to send notification", "error", err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, err)
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "Notification sent")
}
