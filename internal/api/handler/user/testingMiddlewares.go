package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log/slog"
	"net/http"
)

func JWTMiddleware(secretKey string, log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt-token")
			if err != nil {
				responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("missing JWT token")))
				return
			}
			tokenString := cookie.Value

			// Verify the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
				return
			}
			if !token.Valid {
				responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
				return
			}

			// Extract user_id from the token claims
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

			log.Info("User ID extracted from token: ", slog.String("user_id", userID))

			// Store the user ID in the request context
			ctx := context.WithValue(r.Context(), "user_id", userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
