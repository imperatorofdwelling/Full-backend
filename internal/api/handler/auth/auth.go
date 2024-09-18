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
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
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
	r.Post("/registration", h.CreateUser)
	r.Post("/login", h.LoginUser)
}

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.Registration"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var userCurrent auth.Registration
	if err := render.DecodeJSON(r.Body, &userCurrent); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
	}

	userCreated, err := h.Svc.Registration(context.Background(), &userCurrent)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) || errors.Is(err, service.ErrUserNotFound) {
			responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusCreated, userCreated)
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	const op = "handler.user.LoginUser"
	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	var userCurrent auth.Login
	if err := render.DecodeJSON(r.Body, &userCurrent); err != nil {
		h.Log.Error("failed to decode request body", slogError.Err(err))
	}

	userID, err := h.Svc.Login(context.Background(), &userCurrent)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			responseApi.WriteError(w, r, http.StatusNotFound, slogError.Err(err))
			return
		}
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
	})
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	// Set token as a cookie
	cookie := &http.Cookie{
		Name:     "jwt-token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	responseApi.WriteJson(w, r, http.StatusOK, userID)
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
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
			return
		}
		userID, ok := claims["id"].(float64)
		if !ok {
			responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
			return
		}
		ctx := context.WithValue(r.Context(), "userID", int64(userID))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//
//// TODO создать отдельную структуру для возвращения ошибок через api: {error:""}
//func (h *Handler) Login(log *slog.Logger) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		ctx := r.Context()
//
//		var login account.Login
//		err := json.NewDecoder(r.Body).Decode(&login)
//		if err != nil {
//			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
//			log.Error("Ошибка декодирования данных", err)
//			return
//		}
//		result, err := h.service.Login(ctx, login)
//		if err != nil {
//			switch err.Error() {
//			case db.ErrNotExist.Error():
//				http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusNotFound)
//				return
//			default:
//				http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
//				return
//			}
//		}
//		w.Header().Set("Content-Type", "application/json")
//		// Generate JWT token
//		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//			"id":  result,
//			"exp": time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
//		})
//		tokenString, err := token.SignedString([]byte("your-secret-key"))
//		if err != nil {
//			http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
//			return
//		}
//
//		// Set token as a cookie
//		cookie := &http.Cookie{
//			Name:     "jwt-token",
//			Value:    tokenString,
//			Expires:  time.Now().Add(time.Hour * 24),
//			HttpOnly: true,
//		}
//		http.SetCookie(w, cookie)
//
//		response := struct {
//			ID string `json:"id"`
//		}{ID: strconv.FormatInt(result, 10)}
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//}
//
//func (h *Handler) Registration(log *slog.Logger) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		// Check if jwt-token cookie is present and contains a valid ID
//		cookie, err := r.Cookie("jwt-token")
//		if err == nil {
//			tokenString := cookie.Value
//			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//				}
//				return []byte("your-secret-key"), nil
//			})
//			if err != nil {
//				http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
//				return
//			}
//			claims, ok := token.Claims.(jwt.MapClaims)
//			if !ok || !token.Valid {
//				http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
//				return
//			}
//			_, ok = claims["id"].(float64)
//			if !ok {
//				http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
//				return
//			}
//			// If the token is valid, redirect to the dashboard or return an error
//			http.Redirect(w, r, "/dashboard", http.StatusFound)
//			return
//		}
//		ctx := r.Context()
//
//		var reg account.Registration
//		err = json.NewDecoder(r.Body).Decode(&reg)
//		if err != nil {
//			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
//			log.Error("Ошибка декодирования данных", err)
//			return
//		}
//		result, err := h.service.Registration(ctx, reg)
//		if err != nil {
//			switch err.Error() {
//			case db.ErrNotExist.Error():
//				http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusNotFound)
//				return
//			default:
//				http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
//				return
//			}
//		}
//
//		// Generate JWT token
//		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//			"id":  result,
//			"exp": time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
//		})
//		tokenString, err := token.SignedString([]byte("your-secret-key"))
//		if err != nil {
//			http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
//			return
//		}
//
//		// Set token as a cookie
//		cookie = &http.Cookie{
//			Name:     "jwt-token",
//			Value:    tokenString,
//			Expires:  time.Now().Add(time.Hour * 24),
//			HttpOnly: true,
//		}
//		http.SetCookie(w, cookie)
//
//		response := struct {
//			ID string `json:"id"`
//		}{ID: strconv.FormatInt(result.ID, 10)}
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//}
//
