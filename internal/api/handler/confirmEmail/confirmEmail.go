package confirmEmail

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
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
	r.Route("/otp", func(r chi.Router) {
		r.Use(mw.WithAuth)
		r.Get("/", h.CreateOTP)
	})
}

// CreateOTP godoc
//
//	@Summary		Create One-Time Password (OTP)
//	@Description	Generate a one-time password (OTP) for email confirmation by user ID
//	@Tags			email-confirmation
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"success in creating otp!"
//	@Failure		401	{object}	responseApi.ResponseError	"Unauthorized - user not logged in"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error - could not generate OTP"
//	@Router			/otp [get]
func (h *Handler) CreateOTP(w http.ResponseWriter, r *http.Request) {
	const op = "handler.confirmEmail.CreateOTP"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("unauthorized: user not logged in")))
		return
	}

	err := h.Svc.CreateOTP(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to generate one-time password", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(errors.Wrap(err, "could not generate OTP")))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "otp created!")
}
