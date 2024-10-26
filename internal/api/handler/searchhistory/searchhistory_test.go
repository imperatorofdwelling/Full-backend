package searchhistory

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	handler "github.com/imperatorofdwelling/Full-backend/internal/api/handler/user"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSearchHistory_NewHistoryHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewHistorySearchHandler(router)
	})
}

func TestSearchHistory_GetAllHistoryByUserId(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Get("/history", hdl.GetAllHistoryByUserId)

	t.Run("should return error checking user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/history", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)

	})
}

func TestSearchHistory_GetAllHistoryByUserIdError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Get("/history", hdl.GetAllHistoryByUserId)

	t.Run("should return error while getting data by user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/history", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllHistoryByUserId", mock.Anything, testUserID.String()).Return(nil, errors.New("could not fetch history"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)

	})
}

func TestSearchHistory_GetAllHistoryByUserIdSuccess(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Get("/history", hdl.GetAllHistoryByUserId)

	t.Run("should return error while getting data by user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/history", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetAllHistoryByUserId", mock.Anything, testUserID.String()).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

	})
}

func TestSearchHistory_AddHistory(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Post("/history", hdl.AddHistory)

	t.Run("should return error checking user id", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/history", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestSearchHistory_AddHistoryBodyEmpty(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/history", hdl.AddHistory)

	t.Run("should return error checking json body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/history", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestSearchHistory_AddHistoryBodyNameNotFound(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/history", hdl.AddHistory)

	t.Run("should return error checking json body 'name' field", func(t *testing.T) {
		r := httptest.NewRecorder()
		body := `{"name": ""}`

		req := httptest.NewRequest(http.MethodPost, "/history", strings.NewReader(body))

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestSearchHistory_AddHistoryBodyNameFoundButError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/history", hdl.AddHistory)

	t.Run("should return error while running AddHistory func", func(t *testing.T) {
		r := httptest.NewRecorder()
		body := `{"name": "sdf"}`

		req := httptest.NewRequest(http.MethodPost, "/history", strings.NewReader(body))
		svc.On("AddHistory", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error adding a history"))

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestSearchHistory_AddHistoryBodyNameSuccessr(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.SearchHistoryService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(handler.JWTMiddleware("your-secret-key", log))
	router.Post("/history", hdl.AddHistory)

	t.Run("should return error while running AddHistory func", func(t *testing.T) {
		r := httptest.NewRecorder()
		body := `{"name": "sdf"}`

		req := httptest.NewRequest(http.MethodPost, "/history", strings.NewReader(body))
		svc.On("AddHistory", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)
		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
}
