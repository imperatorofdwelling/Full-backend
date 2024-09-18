package handler

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"github.com/imperatorofdwelling/Website-backend/internal/service"
	"github.com/imperatorofdwelling/Website-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Website-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	Svc interfaces.UserService
	Log *slog.Logger
}

func (h *UserHandler) NewUserHandler(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", h.CreateUser)
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.CreateUser"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var user models.UserEntity

	if err := render.DecodeJSON(r.Body, &user); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
	}

	userCreated, err := h.Svc.CreateUser(context.Background(), &user)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) || errors.Is(err, service.ErrUserNotFound) {
			responseApi.WriteJson(w, r, http.StatusBadRequest, slogError.Err(err))
			return
		}

		responseApi.WriteJson(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, userCreated)
}
