package handler

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	modelPass "github.com/imperatorofdwelling/Full-backend/internal/domain/models/passwordOTP"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/jsonReader"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/imperatorofdwelling/Full-backend/pkg/validator"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
	"time"
)

type AuthHandler struct {
	Svc interfaces.AuthService
	Log *slog.Logger
}

func (h *AuthHandler) NewAuthHandler(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Post("/registration", h.Registration)
		r.Post("/login", h.LoginUser)
	})

	r.Group(func(r chi.Router) {
		r.Post("/confirm/password/otp", h.ConfirmPasswordOTP)
	})

	r.Group(func(r chi.Router) {
		r.Use(mw.WithAuth)
		r.Post("/confirm/email/otp/{otp}", h.ConfirmEmailOTP)
		r.Post("/confirm/email/change/otp/{otp}", h.ConfirmEmailChangeOTP)
	})
}

// Registration
//
// @Summary Register a new user
// @Description Creates a new user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   request  body     model.Registration  true  "Registration"
// @Success 201 {object} UUID
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
// @Router /registration [post]
func (h *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.Registration"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var userCurrent model.Registration
	if err := jsonReader.ReadJSON(w, r, &userCurrent); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("failed to decode request body")))
		return
	}

	if !userCurrent.IsHashed {
		// password hashing
		hashedPassword := sha256.Sum256([]byte(userCurrent.Password))
		userCurrent.Password = hex.EncodeToString(hashedPassword[:])
	}

	// creating a new validator for registration
	v := validator.New()
	auth.ValidateRegistration(v, &userCurrent)

	if !v.IsValid() {
		responseApi.WriteError(w, r, http.StatusBadRequest, v.Errors)
		return
	}

	userCreated, err := h.Svc.Register(context.Background(), userCurrent)
	if err != nil {
		h.Log.Error("Error during registration", slog.String("error", err.Error()))

		if errors.Is(err, service.ErrUserAlreadyExists) {
			responseApi.WriteError(w, r, http.StatusBadRequest, fmt.Sprintf("%v", service.ErrUserAlreadyExists))
			return
		}

		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusBadRequest, fmt.Sprintf("%v", service.ErrNotFound))
			return
		}

		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, userCreated)
}

// LoginUser
//
// @Summary Login an existing user
// @Description Authenticates an existing user and returns a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   request  body     model.Login  true  "Login"
// @Success 200 {object} UUID
// @Failure 401 {object} response.ResponseError
// @Failure 404 {object} response.ResponseError
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
// @Router /login [post]
func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.LoginUser"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	cookie, err := r.Cookie("jwt-token")
	if err == nil {
		responseApi.WriteError(w, r, http.StatusUnauthorized, errors.New("already logged in"))
		return
	}

	var userCurrent model.Login
	if err := jsonReader.ReadJSON(w, r, &userCurrent); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, errors.New("failed to decode request body"))
		return
	}

	userID, err := h.Svc.Login(context.Background(), userCurrent)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(service.ErrNotFound))
			return
		}
		if errors.Is(err, service.ErrValid) {
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(service.ErrValid))
			return
		}
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(err))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
		"user_id": userID,
	})
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	// Set token as a cookie
	cookie = &http.Cookie{
		Name:     "jwt-token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	responseApi.WriteJson(w, r, http.StatusOK, userID.String())
}

// ConfirmEmailOTP godoc
//
//	@Summary		Confirm One-Time Password (OTP)
//	@Description	Verify the one-time password (OTP) provided by the user for email confirmation
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			otp		path		string	true		"One-Time Password (OTP)"
//	@Success		200	{string}	string	"OTP confirmed successfully!"
//	@Failure		400	{object}	response.ResponseError	"Bad Request - invalid OTP"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized - user not logged in"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error - could not verify OTP"
//	@Router			/confirm/email/otp/{otp} [post]
func (h *AuthHandler) ConfirmEmailOTP(w http.ResponseWriter, r *http.Request) {
	const op = "handler.auth.ConfirmEmailOTP"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user id not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logger in")))
		return
	}

	otp := chi.URLParam(r, "otp")

	err := h.Svc.CheckEmailOTP(context.Background(), userID, otp)
	if err != nil {
		h.Log.Error("failed to check otp", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not check otp")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "otp confirmed!")
}

// ConfirmPasswordOTP godoc
//
//		@Summary		Confirm One-Time Password (OTP)
//		@Description	Verify the one-time password (OTP) provided by the user for password changing
//		@Tags			auth
//		@Accept			json
//		@Produce		json
//	 @Param   		request  body     modelPass.PasswordOTP  true  "Request body with email and otp"
//		@Success		200		{string}	string	"OTP confirmed successfully!"
//		@Failure		400		{object}	response.ResponseError	"Bad Request - invalid OTP or missing fields"
//		@Failure		500		{object}	response.ResponseError	"Internal Server Error - could not verify OTP"
//		@Router			/confirm/password/otp [post]
func (h *AuthHandler) ConfirmPasswordOTP(w http.ResponseWriter, r *http.Request) {
	const op = "handler.auth.ConfirmPasswordOTP"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var req modelPass.PasswordOTP
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("invalid request body")))
		return
	}

	if req.Email == "" || req.OTP == "" {
		h.Log.Error("missing email or otp in request")
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("email and otp are required")))
		return
	}

	err := h.Svc.CheckPasswordOTP(context.Background(), req.Email, req.OTP)
	if err != nil {
		h.Log.Error("failed to check otp", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not check otp")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "password otp confirmed!")
}

// ConfirmEmailChangeOTP godoc
//
//	@Summary		Confirm Email Change OTP
//	@Description	Verify the one-time password (OTP) provided by the user for email change confirmation
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			otp		path		string	true		"One-Time Password (OTP)"
//	@Success		200	{string}	string	"OTP confirmed successfully!"
//	@Failure		400	{object}	response.ResponseError	"Bad Request - invalid OTP"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized - user not logged in"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error - could not verify OTP"
//	@Router			/confirm/email/change/otp/{otp} [post]
func (h *AuthHandler) ConfirmEmailChangeOTP(w http.ResponseWriter, r *http.Request) {
	const op = "handler.auth.ConfirmEmailChangeOTP"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user id not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logger in")))
		return
	}

	otp := chi.URLParam(r, "otp")

	err := h.Svc.CheckEmailChangeOTP(context.Background(), userID, otp)
	if err != nil {
		h.Log.Error("failed to check otp", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not check otp")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "otp confirmed!")
}
