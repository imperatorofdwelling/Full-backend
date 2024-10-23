package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func TestAuthHandler_Registration(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	fakeUUID, _ := uuid.NewV4()

	router := chi.NewRouter()
	router.Post("/register", hdl.Registration)

	tests := []struct {
		name         string
		body         string
		expectErr    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "should create a new user successfully",
			body:         `{"name": "test user", "email": "testuser@example.com", "password": "password123"}`,
			expectErr:    nil,
			expectedCode: http.StatusCreated,
			expectedBody: "",
		},
		{
			name:         "should return error if user already exists",
			body:         `{"name": "test user", "email": "testuser@example.com", "password": "password123"}`,
			expectErr:    service.ErrUserAlreadyExists,
			expectedCode: http.StatusBadRequest,
			expectedBody: "user already exists",
		},
		{
			name:         "should be error for unknown field",
			body:         `{"name":"test user", "unknown_field":"value"}`,
			expectErr:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "{\"error\":\"error=failed to decode request body\"}\n",
		},

		{
			name:         "should return error for invalid registration",
			body:         `{"name":"", "email":"invalid-email", "password":""}`,
			expectErr:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "{\"email\":\"must be in correct form\",\"name\":\"length of the name must be greater than 5\",\"password\":\"length of the password must be greater than 5\"}\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRecorder()

			if tt.expectErr == nil && tt.name != "should return error if user already exists" {
				var userToRegister model.Registration
				json.Unmarshal([]byte(tt.body), &userToRegister)
				svc.On("Register", mock.Anything, userToRegister).Return(fakeUUID, nil).Once()
			} else if tt.name == "should return error if user already exists" {
				var userToRegister model.Registration
				json.Unmarshal([]byte(tt.body), &userToRegister)

				svc.On("Register", mock.Anything, userToRegister).Return(uuid.Nil, tt.expectErr).Once()
			}

			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(r, req)

			assert.Equal(t, tt.expectedCode, r.Code)
			if tt.expectedBody != "" {
				assert.Contains(t, r.Body.String(), tt.expectedBody)
			}

		})
	}
}

func TestAuthHandler_RegistrationInternalError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := &mocks.AuthService{}
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/registration", hdl.Registration)

	t.Run("should return internal server error after validation", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := []byte(`{"name":"test user", "email":"testuser@example.com", "password":"password123"}`)
		req := httptest.NewRequest(http.MethodPost, "/registration", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		var userToRegister model.Registration
		json.Unmarshal(body, &userToRegister)
		svc.On("Register", mock.Anything, userToRegister).Return(uuid.Nil, errors.New("internal server error")).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
		assert.Contains(t, r.Body.String(), "{\"error\":\"error=internal server error\"}\n")
	})
}

func TestAuthHandler_Login(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.AuthService)
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	fakeUUID, _ := uuid.NewV4()

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return ok", func(t *testing.T) {
		r := httptest.NewRecorder()

		pBody := `{"email": "testuser@example.com", "password": "password123"}`

		obj := model.Login{
			Email:    "testuser@example.com",
			Password: "password123",
		}

		svc.On("Login", mock.Anything, obj).Return(fakeUUID, nil)

		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(pBody))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should return unauthorized if already logged in", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		req.AddCookie(&http.Cookie{Name: "jwt-token", Value: "some-token"})

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
		assert.Contains(t, r.Body.String(), "already logged in")
	})

	t.Run("should return bad request if login fails", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := []byte(`{"email":"testuser@example.com", "password":"wrong2password"}`)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		svc.On("Login", mock.Anything, mock.Anything).Return(nil, service.ErrValid)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
		assert.Contains(t, r.Body.String(), "{\"error\":\"error=invalid data\"}\n")
	})

	t.Run("should return bad request if JSON is invalid", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := []byte(`{"email":"testuser@example.com", "password":}`)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

}

func TestAuthHandler_LoginBadRequest(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.AuthService)
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return bad request if JSON is invalid", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := []byte(`{"email":"testuser@example.com", "password"}`)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

}

func TestAuthHandler_LoginNotFound(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.AuthService)
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return not found", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := []byte(`{"email":"testuser@example.com", "password":"wrong2password"}`)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		svc.On("Login", mock.Anything, mock.Anything).Return(nil, service.ErrNotFound)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
		assert.Contains(t, r.Body.String(), "{\"error\":\"error=not found\"}\n")
	})
}

func TestAuthHandler_LoginInternalServerError(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.AuthService)
	hdl := AuthHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Post("/login", hdl.LoginUser)

	t.Run("should return internal server error", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := []byte(`{"email":"testuser@example.com", "password":"wrong2password"}`)
		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		svc.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("internal server error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
		assert.Contains(t, r.Body.String(), "{\"error\":\"error=internal server error\"}\n")
	})

}
