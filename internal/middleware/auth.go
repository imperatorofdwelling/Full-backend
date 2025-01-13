package mw

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"log"
	"net/http"
	"os"
)

type contextKey string

// Константы для получения значений из куки
const (
	userIDKey   contextKey = "user_id"
	userRoleKey contextKey = "user_role"
	pathKey     contextKey = "request_path"
)

func WithAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := getTokenFromRequest(r)
		if err != nil {
			permissionDenied(w, r, "unable to get token from request")
		}

		token, err := validateToken(tokenString)
		if err != nil || !token.Valid {
			permissionDenied(w, r, "invalid token")
		}

		userID, err := getUserIDFromToken(token)
		if err != nil {
			permissionDenied(w, r, "unable to get user ID from token")
		}

		userRole, err := getUserRoleFromToken(token)
		if err != nil {
			permissionDenied(w, r, "unable to get user role from token")
			return
		}
		requestPath := r.URL.Path
		log.Println(requestPath)

		// Store the user ID in the request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		ctx = context.WithValue(r.Context(), userRoleKey, userRole)
		r = r.WithContext(ctx)

		handler.ServeHTTP(w, r)
	})
}

func permissionDenied(w http.ResponseWriter, r *http.Request, error string) {
	responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("permission denied: "+error)))
	return
}

func getTokenFromRequest(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		return "", err
	}

	tokenString := cookie.Value

	return tokenString, nil
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY_AUTH")), nil
	})
}

func getUserIDFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string) // ID сохраняется при LOGIN
	if !ok {
		return "", errors.New("invalid user ID in token")
	}

	return userID, nil
}

func getUserRoleFromToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userRole, ok := claims["role_id"].(string) // Роль сохраняется при LOGIN
	if !ok {
		return "", errors.New("invalid user role in token")
	}

	return userRole, nil
}
