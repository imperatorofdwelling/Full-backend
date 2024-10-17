package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
	"time"
)

type AuthHandler struct {
	Svc interfaces.AuthService
	Log *slog.Logger
}

func (h *AuthHandler) NewAuthHandler(r chi.Router) {
	r.Post("/registration", h.Registration)
	r.Post("/login", h.LoginUser)
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
// @Failure 400 {object} responseApi.ResponseError
// @Failure 500 {object} responseApi.ResponseError
// @Router /registration [post]
func (h *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.Registration"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var userCurrent model.Registration
	if err := render.DecodeJSON(r.Body, &userCurrent); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("failed to decode request body")))
		return
	}

	userCreated, err := h.Svc.Register(context.Background(), userCurrent)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) || errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
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
// @Failure 401 {object} responseApi.ResponseError
// @Failure 404 {object} responseApi.ResponseError
// @Failure 400 {object} responseApi.ResponseError
// @Failure 500 {object} responseApi.ResponseError
// @Router /login [post]
func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.LoginUser"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	cookie, err := r.Cookie("jwt-token")
	if err == nil {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("already logged in")))
		return
	}

	var userCurrent model.Login
	if err := render.DecodeJSON(r.Body, &userCurrent); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
	}

	userID, err := h.Svc.Login(context.Background(), userCurrent)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
			return
		}
		if errors.Is(err, service.ErrValid) {
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
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

func (h *AuthHandler) JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt-token")
		if err != nil {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(err))
			return
		}
		tokenString := cookie.Value
		// Verify the token as before
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("your-secret-key"), nil
		})
		if err != nil {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
			return
		}
		if !token.Valid {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
			return
		}
		// Extract the user ID from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token claims")))
			return
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid user ID in token")))
			return
		}
		// Store the user ID in the request context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
