package mw

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	domain "github.com/imperatorofdwelling/Full-backend/internal/domain/models/role"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"net/http"
	"os"
	"strings"
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

		tokenClaims, err := getTokenClaims(token)
		if err != nil {
			permissionDenied(w, r, "cannot get token")
		}

		userID, err := getUserIDFromTClaims(tokenClaims)
		if err != nil {
			permissionDenied(w, r, "unable to get user ID from token")
		}

		userRole, err := getUserRoleFromClaims(tokenClaims)
		if err != nil {
			permissionDenied(w, r, "unable to get user role from token")
			return
		}

		// Получаем путь и метод запроса
		requestPath := r.URL.Path
		requestMethod := r.Method

		// Проверяем доступ в зависимости от роли
		if !permissionCheck(userRole, requestPath, requestMethod) {
			responseApi.WriteError(w, r, http.StatusForbidden, "forbidden")
			return
		}

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

func getTokenClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}

func getUserIDFromTClaims(claims jwt.MapClaims) (string, error) {
	userID, ok := claims["user_id"].(string) // ID сохраняется при LOGIN
	if !ok {
		return "", errors.New("invalid user ID in token")
	}

	return userID, nil
}

func getUserRoleFromClaims(claims jwt.MapClaims) (float64, error) {
	userRole, ok := claims["user_role"].(float64) // Роль сохраняется при LOGIN
	if !ok {
		return -1, errors.New("invalid user role in token")
	}

	return userRole, nil
}

func permissionCheck(role float64, path string, method string) bool {
	var routes []domain.Route

	// Определяем маршруты в зависимости от роли
	if role == domain.TenantRole {
		routes = domain.TenantRoutes
	} else if role == domain.LandlordRole {
		routes = domain.LandlordRoutes
	} else {
		return false // Неизвестная роль
	}

	// Проверяем, есть ли путь и метод в разрешенных маршрутах
	for _, route := range routes {
		if strings.Contains(path, route.Path) {
			for _, allowedMethod := range route.Methods {
				if strings.Contains(allowedMethod, method) {
					return true
				}
			}
		}
	}

	return false // Доступ запрещен
}
