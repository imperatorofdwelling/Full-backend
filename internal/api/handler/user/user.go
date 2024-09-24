package handler

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	Svc interfaces.UserService
	Log *slog.Logger
}

func (h *UserHandler) NewUserHandler(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Put("/{id}", h.UpdateUserByID)
		r.Delete("/{id}", h.DeleteUserByID)
	})
}

func (h *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.LoginUser"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var id = chi.URLParam(r, "id")

	var updateUser model.User
	if err := render.DecodeJSON(r.Body, &updateUser); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := h.Svc.UpdateUserByID(context.Background(), id, updateUser)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
			return
		}
		if errors.Is(err, service.ErrUpdateFailed) {
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
			return
		}
		if errors.Is(err, service.ErrEmailAlreadyExists) {
			responseApi.WriteError(w, r, http.StatusConflict, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}
	responseApi.WriteJson(w, r, http.StatusOK, result)
}

func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.DeleteUserByID"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var id = chi.URLParam(r, "id")

	err := h.Svc.DeleteUserByID(context.Background(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}
	responseApi.WriteJson(w, r, http.StatusNoContent, nil)
}
