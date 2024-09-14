package handler

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/dgrijalva/jwt-go/v4"
//	"github.com/imperatorofdwelling/Website-backend/internal/db"
//	"github.com/imperatorofdwelling/Website-backend/internal/domain/account"
//	"log/slog"
//	"net/http"
//	"strconv"
//	"time"
//)
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
//func JWTMiddleware(log *slog.Logger) func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			cookie, err := r.Cookie("jwt-token")
//			if err != nil {
//				http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
//				return
//			}
//			tokenString := cookie.Value
//			// Verify the token as before
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
//			userID, ok := claims["id"].(float64)
//			if !ok {
//				http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
//				return
//			}
//			ctx := context.WithValue(r.Context(), "userID", int64(userID))
//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}
//}
