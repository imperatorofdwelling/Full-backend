package handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
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
	}

	result, err := h.Svc.UpdateUserByID(context.Background(), id, updateUser)
	if err != nil {
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
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
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}
	responseApi.WriteJson(w, r, http.StatusNoContent, nil)
}
