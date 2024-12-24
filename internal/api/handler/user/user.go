package handler

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	modelPass "github.com/imperatorofdwelling/Full-backend/internal/domain/models/newPassword"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
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
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Put("/{id}", h.UpdateUserByID)
			r.Delete("/{id}", h.DeleteUserByID)
		})

		r.Group(func(r chi.Router) {
			r.Get("/{id}", h.GetUserByID)
			r.Put("/password", h.UpdateUserPasswordByEmail)
		})
	})
}

// GetUserByID
//
// @Summary Get a user by ID
// @Description Retrieves a user by the provided ID
// @ID getUserByID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path     string     true  "User  ID"
// @Success 200 {object} model.User
// @Failure 400 {object} responseApi.ResponseError "Invalid request"
// @Failure 404 {object} responseApi.ResponseError "User  not found"
// @Router /user/{id} [get]
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.GetUserByID"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var id = chi.URLParam(r, "id")
	_, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse UUID", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	result, err := h.Svc.GetUserByID(context.Background(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, result)

}

// UpdateUserByID
//
// @Summary Update a user by ID
// @Description Update a user with the provided ID
// @ID updateUserByID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security ApiKeyAuth
// @Param   model.User	body     model.User true  "User update data"
// @Success 200 {object} model.User
// @Failure 400 {object} responseApi.ResponseError "Invalid request"
// @Failure 404 {object} responseApi.ResponseError "User not found"
// @Failure 500 {object} responseApi.ResponseError "Internal server error"
// @Router /user/{id} [put]
func (h *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.LoginUser"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	_, ok := r.Context().Value("user_id").(string)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid user ID in context")))
		return
	}

	var id = chi.URLParam(r, "id")

	parsedID, err := uuid.FromString(id)
	if err != nil {
		h.Log.Error("failed to parse UUID", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	var updateUser model.User
	if err := render.DecodeJSON(r.Body, &updateUser); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	// Передача контекста
	result, err := h.Svc.UpdateUserByID(context.Background(), parsedID.String(), updateUser)
	if err != nil {
		// Обработка ошибок
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

// DeleteUserByID
//
// @Summary Delete a user by ID
// @Description Delete a user with the provided ID
// @ID deleteUserByID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security ApiKeyAuth
// @Success 204 {object} nil "User successfully deleted"
// @Failure 400 {object} responseApi.ResponseError "Invalid request"
// @Failure 401 {object} responseApi.ResponseError "Unauthorized"
// @Failure 404 {object} responseApi.ResponseError "User not found"
// @Failure 500 {object} responseApi.ResponseError "Internal server error"
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.DeleteUserByID"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	_, ok := r.Context().Value("user_id").(string)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid user ID in context")))
		return
	}

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

// UpdateUserPasswordByEmail
//
// @Summary Update user password by email
// @Description Updates the user's password after verifying the OTP and checking its expiration
// @ID updateUserPasswordByEmail
// @Tags users
// @Accept json
// @Produce json
// @Param userNewPassword body modelPass.NewPassword true "User's new password and OTP"
// @Success 200 {string} string "Password changed successfully"
// @Failure 400 {object} responseApi.ResponseError "Invalid request or OTP verification failed"
// @Failure 404 {object} responseApi.ResponseError "User not found"
// @Failure 500 {object} responseApi.ResponseError "Internal server error"
// @Router /user/password [put]
func (h *UserHandler) UpdateUserPasswordByEmail(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.UpdateUserPasswordByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var userNewPassword modelPass.NewPassword
	if err := render.DecodeJSON(r.Body, &userNewPassword); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	if userNewPassword.Password == "" || userNewPassword.Email == "" {
		h.Log.Error("invalid request body", slogError.Err(errors.New("invalid request body")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("invalid request body")))
		return
	}

	err := h.Svc.CheckUserPassword(context.Background(), userNewPassword)
	if err != nil {
		h.Log.Error("failed to check user password", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.UpdateUserPasswordByEmail(context.Background(), userNewPassword)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
		}
	}

	responseApi.WriteJson(w, r, http.StatusOK, "password changed")
}
