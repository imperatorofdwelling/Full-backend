package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/api/handler"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	modelPass "github.com/imperatorofdwelling/Full-backend/internal/domain/models/newPassword"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/internal/service/file"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
	"strings"
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
			r.Post("/profile/picture", h.CreateUserPfp)
			r.Get("/profile/picture", h.GetUserPfp)
			r.Put("/email/change", h.UpdateUserEmailById)
		})

		r.Group(func(r chi.Router) {
			r.Get("/profile/picture/{id}", h.GetUserPfpByUserID)
			r.Get("/{id}", h.GetUserByID)
			r.Put("/password", h.UpdateUserPasswordByEmail)
			r.Patch("/profile/picture/{id}", h.UpdateUserPfp)
			r.Delete("/profile/picture/{id}", h.DeleteUserPfp)
		})
	})
}

// UpdateUserPfp
//
// @Summary Change user avatar
// @Description Change user avatar by id
// @ID changeUserPfp
// @Tags users
// @Accept multipart/form-data
// @Produce  json
// @Security ApiKeyAuth
// @Param   id   path     string     true  "User  ID"
// @Param image formData file true "User's profile picture (JPEG or PNG)"
// @Success 200 {object} string "Successfully updated"
// @Failure 400 {object} response.ResponseError "Invalid request"
// @Failure 404 {object} response.ResponseError "User  not found"
// @Router /user/profile/picture/{id} [patch]
func (h *UserHandler) UpdateUserPfp(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.UpdateUserPfp"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID := chi.URLParam(r, "id")

	uuidID, err := uuid.FromString(userID)
	if err != nil {
		h.Log.Error("failed to parse user id: ", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, "failed to parse user id")
		return
	}

	err = r.ParseMultipartForm(file.MaxImageMemorySize)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	formFiles := r.MultipartForm.File

	image := formFiles["image"][0]

	err = h.Svc.ChangeUserPfp(r.Context(), uuidID, image)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "successfully updated")
}

// DeleteUserPfp
//
// @Summary Delete user avatar
// @Description Delete user avatar by user id
// @ID deleteUserPfp
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security ApiKeyAuth
// @Success 204 {object} string "no content"
// @Failure 400 {object} response.ResponseError "Invalid request"
// @Failure 401 {object} response.ResponseError "Unauthorized"
// @Failure 404 {object} response.ResponseError "User not found"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/profile/picture/{id} [delete]
func (h *UserHandler) DeleteUserPfp(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.DeleteUserPfp"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID := chi.URLParam(r, "id")

	uuidID, err := uuid.FromString(userID)
	if err != nil {
		h.Log.Error("failed to parse user id: ", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, "failed to parse user id")
		return
	}

	err = h.Svc.DeleteUserPfp(r.Context(), uuidID)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "successfully deleted")
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
// @Failure 400 {object} response.ResponseError "Invalid request"
// @Failure 404 {object} response.ResponseError "User  not found"
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
// @Failure 400 {object} response.ResponseError "Invalid request"
// @Failure 404 {object} response.ResponseError "User not found"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/{id} [put]
func (h *UserHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.LoginUser"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	_, ok := r.Context().Value(mw.UserIdKey).(string)
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
// @Failure 400 {object} response.ResponseError "Invalid request"
// @Failure 401 {object} response.ResponseError "Unauthorized"
// @Failure 404 {object} response.ResponseError "User not found"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.DeleteUserByID"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)
	_, ok := r.Context().Value(mw.UserIdKey).(string)
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
// @Failure 400 {object} response.ResponseError "Invalid request or OTP verification failed"
// @Failure 404 {object} response.ResponseError "User not found"
// @Failure 500 {object} response.ResponseError "Internal server error"
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
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "password changed")
}

// UpdateUserEmailById
//
// @Summary Update user email by ID
// @Description Updates the user's email after validating the request
// @ID updateUserEmailByID
// @Tags users
// @Accept json
// @Produce json
// @Param userEmail body map[string]string true "User's new email"
// @Success 200 {object} map[string]string "Email changed successfully"
// @Failure 400 {object} response.ResponseError "Invalid request or email validation failed"
// @Failure 401 {object} response.ResponseError "User not logged in"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/email/change [put]
func (h *UserHandler) UpdateUserEmailById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.UpdateUserEmailByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.Log.Error("failed to decode JSON body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("invalid JSON body")))
		return
	}

	newEmail, ok := reqBody["email"]
	if !ok || newEmail == "" {
		h.Log.Error("email is missing or empty", slogError.Err(errors.New("email field is required")))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("email field is required")))
		return
	}

	err := h.Svc.CheckUserEmail(context.Background(), userID, newEmail)
	if err != nil {
		h.Log.Error("failed to check user email", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.UpdateUserEmailByID(context.Background(), userID, newEmail)
	if err != nil {
		h.Log.Error("failed to update user email", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]string{"message": "email changed successfully"})
}

// CreateUserPfp
//
// @Summary Create user profile picture
// @Description Uploads a new profile picture for the authenticated user
// @ID createUserPfp
// @Tags users
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "User's profile picture (JPEG or PNG)"
// @Success 201 {object} map[string]string "User pfp added successfully"
// @Failure 400 {object} response.ResponseError "Invalid request or unsupported content type"
// @Failure 401 {object} response.ResponseError "User not logged in"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/profile/picture [post]
func (h *UserHandler) CreateUserPfp(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.CreateUserPfp"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, file.MaxImageMemorySize)

	err := r.ParseMultipartForm(file.MaxImageMemorySize)
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	image, hdl, err := r.FormFile("image")
	if err != nil {
		h.Log.Error("failed to parse form", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}
	defer image.Close()

	contentType := hdl.Header.Get("Content-Type")
	if !(strings.Contains(contentType, "image/jpeg") || strings.Contains(contentType, "image/png")) {
		h.Log.Error("unsupported content type", slogError.Err(handler.ErrInvalidImageType))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(handler.ErrInvalidImageType))
		return
	}

	buf := make([]byte, hdl.Size)
	n, err := image.Read(buf)
	if err != nil {
		h.Log.Error("failed to read image", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	if err = h.Svc.CreateUserPfp(r.Context(), userID, buf[:n]); err != nil {
		h.Log.Error("service failed to create user profile picture", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, map[string]string{"message": "User pfp added successfully"})
}

// GetUserPfp
//
// @Summary Get user profile picture
// @Description Retrieves the profile picture path for the authenticated user
// @ID getUserPfp
// @Tags users
// @Produce json
// @Success 200 {string} string "Path to the user's profile picture"
// @Failure 401 {object} response.ResponseError "User not logged in"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/profile/picture [get]
func (h *UserHandler) GetUserPfp(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.GetUserPfp"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user not logged in", slogError.Err(errors.New("user not logged in")))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	pfpPath, err := h.Svc.GetUserPfp(context.Background(), userID)
	if err != nil {
		h.Log.Error("service failed to get user profile picture", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, pfpPath)
}

// GetUserPfpByUserID
//
// @Summary Get user profile picture by ID
// @Description Retrieves the profile picture path for the specified user
// @ID getUserPfpByUserID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "Path to the user's profile picture"
// @Failure 400 {object} response.ResponseError "Invalid user ID"
// @Failure 500 {object} response.ResponseError "Internal server error"
// @Router /user/profile/picture/{id} [get]
func (h *UserHandler) GetUserPfpByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.GetUserPfpByUserID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var id = chi.URLParam(r, "id")

	pfpPath, err := h.Svc.GetUserPfp(context.Background(), id)
	if err != nil {
		h.Log.Error("service failed to get user profile picture", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, pfpPath)
}
