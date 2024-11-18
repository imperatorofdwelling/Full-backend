package mw

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"net/http"
	"os"
)

func WithAuth(handler http.Handler) http.Handler {
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
			return []byte(os.Getenv("SECRET_KEY_AUTH")), nil
		})
		if err != nil || !token.Valid {
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

		handler.ServeHTTP(w, r)
	})
}
