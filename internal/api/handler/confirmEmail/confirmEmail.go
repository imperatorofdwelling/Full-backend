package confirmEmail

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.ConfirmEmailService
	Log *slog.Logger
}

func (h *Handler) NewConfirmEmailHandler(r chi.Router) {
	r.Route("/email/otp", func(r chi.Router) {
		r.Use(mw.WithAuth)
		r.Get("/", h.CreateOTPEmail)
	})

	r.Route("/email/change/otp", func(r chi.Router) {
		r.Use(mw.WithAuth)
		r.Get("/", h.SendOtpForEmailChange)
	})

	r.Route("/password/otp", func(r chi.Router) {
		r.Get("/{email}", h.CreateOTPPassword)
	})
}

// CreateOTPEmail godoc
//
//	@Summary		Create One-Time Password (OTP)
//	@Description	Generate a one-time password (OTP) for email confirmation by user ID
//	@Tags			emailConfirmation
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"success in creating otp for email verification!"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized - user not logged in"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error - could not generate OTP"
//	@Router			/email/otp [get]
func (h *Handler) CreateOTPEmail(w http.ResponseWriter, r *http.Request) {
	const op = "handler.confirmEmail.CreateOTP"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logged in")))
		return
	}

	h.Log.Info(userID)
	fmt.Println(userID)

	err := h.Svc.CreateOTPEmail(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to generate one-time password for email verification", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not generate OTP")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "otp for email verification created!")
}

// CreateOTPPassword godoc
//
//	@Summary		Create One-Time Password (OTP) for Password Reset
//	@Description	Generate a one-time password (OTP) for resetting the password using the provided email
//	@Tags			emailConfirmation
//	@Accept			json
//	@Produce		json
//	@Param			email	path		string	true	"Email of the user"
//	@Success		200		{string}	string	"success in creating otp for password reset!"
//	@Failure		500		{object}	response.ResponseError	"Internal Server Error - could not generate OTP"
//	@Router			/password/otp/{email} [get]
func (h *Handler) CreateOTPPassword(w http.ResponseWriter, r *http.Request) {
	const op = "handler.confirmEmail.CreateOTPPassword"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	email := chi.URLParam(r, "email")

	err := h.Svc.CreateOTPPassword(context.Background(), email)
	if err != nil {
		h.Log.Error("failed to generate one-time password for password reset", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not generate OTP")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "otp for password reset created!")
}

// SendOtpForEmailChange godoc
//
//	@Summary		Send OTP for Email Change
//	@Description	Generate and send a one-time password (OTP) to the user's current email for changing their email address
//	@Tags			emailConfirmation
//	@Accept			json
//	@Produce		json
//	@Success		200		{string}	string	"otp for email change created!"
//	@Failure		401		{object}	response.ResponseError	"Unauthorized - invalid user ID in context"
//	@Failure		500		{object}	response.ResponseError	"Internal Server Error - failed to send OTP for email change"
//	@Router			/email/change/otp [get]
func (h *Handler) SendOtpForEmailChange(w http.ResponseWriter, r *http.Request) {
	const op = "handler.confirmEmail.SendOtpForEmailChange"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userId, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid user ID in context")))
		return
	}

	err := h.Svc.SendOtpForEmailChange(context.Background(), userId)
	if err != nil {
		h.Log.Error("failed to send otp for email change", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "otp for email change created!")
}
