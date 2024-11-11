package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAuthHandler_NewAuthHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewAuthHandler(router)
	})
}

func TestAuthHandler_Registration_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/register", hdl.Registration)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"name": "test user", "email": "testuser@example.com", "password": "password123"}`

		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(payload))

		var registration auth.Registration
		err := json.Unmarshal([]byte(payload), &registration)
		assert.NoError(t, err)

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), 12)
		assert.NoError(t, err)
		registration.Password = string(hashedPassword)

		svc.On("Register", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})

	t.Run("should be an error with body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/register", nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestAuthHandler_Registration_Errors(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/register", hdl.Registration)

	t.Run("should return error when user already exists", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"name": "test user", "email": "testuser@example.com", "password": "password123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(payload))

		var registration auth.Registration
		err := json.Unmarshal([]byte(payload), &registration)
		assert.NoError(t, err)

		// Мокаем ошибку, когда пользователь уже существует
		svc.On("Register", mock.Anything, mock.Anything).Return(nil, service.ErrUserAlreadyExists)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "user already exists")
	})

	t.Run("should return bad request when fields are invalid", func(t *testing.T) {
		r := httptest.NewRecorder()

		// Тело запроса с некорректными данными (например, без email или с коротким паролем)
		payload := `{"name": "test user", "email": "invalid-email", "password": "123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(payload))

		router.ServeHTTP(r, req)

		// Проверка, что сервер возвращает ошибку 400 (Bad Request)
		assert.Equal(t, http.StatusBadRequest, r.Code)

		assert.Contains(t, r.Body.String(), "email")
	})
}

func TestAuthHandler_Registration_Errors_Already_Exists(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/register", hdl.Registration)

	t.Run("should return error when resource not found", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"name": "test user", "email": "testuser@example.com", "password": "password123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(payload))

		var registration auth.Registration
		err := json.Unmarshal([]byte(payload), &registration)
		assert.NoError(t, err)

		svc.On("Register", mock.Anything, mock.Anything).Return(nil, service.ErrNotFound)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "not found")
	})

}

func TestAuthHandler_Registration_Errors_Internal_Server_Error(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/register", hdl.Registration)

	t.Run("should return internal server error for unknown errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"name": "test user", "email": "testuser@example.com", "password": "password123"}`
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(payload))

		var registration auth.Registration
		err := json.Unmarshal([]byte(payload), &registration)
		assert.NoError(t, err)

		// Мокаем ошибку, которая не обрабатывается явно
		svc.On("Register", mock.Anything, mock.Anything).Return(nil, errors.New("unexpected error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

}

func TestAuthHandler_LoginUser_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should login successfully", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"email": "testuser@example.com", "password": "password123"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))

		svc.On("Login", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)

		cookies := r.Result().Cookies()
		assert.Len(t, cookies, 1)
		assert.Equal(t, "jwt-token", cookies[0].Name)
	})
}

func TestAuthHandler_LoginUser_LoggedIn(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return error if already logged in", func(t *testing.T) {
		r := httptest.NewRecorder()

		// Создаем request с действующим jwt-токеном в cookie
		testUserID, _ := uuid.NewV4()
		testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
			"user_id": testUserID.String(),
		})
		tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("Login", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
		assert.Contains(t, r.Body.String(), "already logged in")
	})
}

func TestAuthHandler_LoginUser_NoBody(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return error if no body was provided", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/login", nil)

		svc.On("Login", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

}

func TestAuthHandler_LoginUser_ErrorHandling(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return not found error", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"email": "testuser@example.com", "password": "password123"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))

		// Мокаем ошибку ErrNotFound
		svc.On("Login", mock.Anything, mock.Anything).Return(nil, service.ErrNotFound)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestAuthHandler_LoginUser_ErrorHandling_Request_Error(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return bad request error", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"email": "testuser@example.com", "password": "invalidpassword"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))

		// Мокаем ошибку ErrValid
		svc.On("Login", mock.Anything, mock.Anything).Return(nil, service.ErrValid)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestAuthHandler_LoginUser_ErrorHandling_Internal_Error(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return internal server error", func(t *testing.T) {
		r := httptest.NewRecorder()

		payload := `{"email": "testuser@example.com", "password": "password123"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(payload))

		// Мокаем общую ошибку
		svc.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("unexpected error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}