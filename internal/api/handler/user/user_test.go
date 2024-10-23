package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserHandler_NewUserHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	t.Run("should be no error", func(t *testing.T) {
		hdl.NewUserHandler(router)
	})
}

func TestUserHandler_NewPublicUserHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	t.Run("should be no error", func(t *testing.T) {
		hdl.NewPublicUserHandler(router)
	})
}

func TestUserHandler_GetUserByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	testUserID, _ := uuid.NewV4()
	expected := user.User{
		ID:        testUserID,
		Name:      "John Doe",
		Email:     "johndoe@mail.ru",
		Phone:     "123456789",
		Avatar:    nil,
		BirthDate: sql.NullTime{},
		National:  "",
		Gender:    "",
		Country:   "",
		City:      "",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetUserByID", mock.Anything, mock.Anything).Return(expected, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/user/"+testUserID.String(), nil)

		router.HandleFunc("/user/{id}", hdl.GetUserByID)

		router.ServeHTTP(r, req)

		var actual user.User

		_ = render.DecodeJSON(r.Body, &actual)

		assert.Equal(t, http.StatusNoContent, r.Code)

		assert.Equal(t, expected.ID, actual.ID)
	})

	t.Run("should return bad request for invalid UUID", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodGet, "/user/"+invalidUUID, nil)

		router.HandleFunc("/user/{id}", hdl.GetUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return not found", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetUserByID", mock.Anything, testUserID.String()).Return(user.User{}, service.ErrNotFound).Once()

		req := httptest.NewRequest(http.MethodGet, "/user/"+testUserID.String(), nil)

		router.HandleFunc("/user/{id}", hdl.GetUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestUserHandler_GetUserByIDBadRequest(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.UserService)
	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.HandleFunc("/user/{id}", hdl.GetUserByID)

	testUserID, _ := uuid.NewV4()

	t.Run("should return bad request for invalid user ID", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("GetUserByID", mock.Anything, mock.Anything).Return(user.User{}, fmt.Errorf("invalid id"))

		req := httptest.NewRequest(http.MethodGet, "/user/"+testUserID.String(), nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
		svc.AssertExpectations(t)
	})
}

func TestUserHandler_UpdateUserByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})

	tokenString, err := testToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	payload := user.User{
		Name:    "Updated Name",
		Email:   "updated@example.com",
		Phone:   "1234567890",
		Country: "Updated Country",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	t.Run("should return unauthorized without token", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(payloadBytes))
		req.Header.Set("Content-Type", "application/json")

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

	t.Run("should update user successfully", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		svc.On("UpdateUserByID", mock.Anything, testUserID.String(), payload).Return(user.User{}, nil).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
		svc.AssertExpectations(t)
	})

	t.Run("should return 404 when user not found", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		svc.On("UpdateUserByID", mock.Anything, testUserID.String(), payload).Return(user.User{}, service.ErrNotFound).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)

		svc.AssertExpectations(t)
	})

	t.Run("should return 400 when update fails", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		svc.On("UpdateUserByID", mock.Anything, testUserID.String(), payload).Return(user.User{}, service.ErrUpdateFailed).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
		svc.AssertExpectations(t)
	})
	t.Run("should return 409 when email already exists", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		svc.On("UpdateUserByID", mock.Anything, testUserID.String(), payload).Return(user.User{}, service.ErrEmailAlreadyExists).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusConflict, r.Code)
		svc.AssertExpectations(t)
	})

	t.Run("should return 500 on unknown error", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		svc.On("UpdateUserByID", mock.Anything, testUserID.String(), payload).Return(user.User{}, errors.New("unknown error")).Once()

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
		svc.AssertExpectations(t)
	})
}

func TestUserHandler_UpdateUserByIdUnauthorized(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.UserService)
	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}

	testUserID, _ := uuid.NewV4()

	payload := user.User{
		Name:    "Updated Name",
		Email:   "updated@example.com",
		Phone:   "1234567890",
		Country: "Updated Country",
	}

	t.Run("should return unauthorized when user_id is not present in context", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		router := chi.NewRouter()

		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)

		svc.AssertExpectations(t)
	})
}

func TestUserHandler_UpdateUserByIdUUID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.UserService)
	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}

	testUserID, _ := uuid.NewV4()

	payload := user.User{
		Name:    "Updated Name",
		Email:   "updated@example.com",
		Phone:   "1234567890",
		Country: "Updated Country",
	}

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})

	tokenString, err := testToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	t.Run("should return error parsing uuid", func(t *testing.T) {
		r := httptest.NewRecorder()
		reqBody, _ := json.Marshal(payload)

		req := httptest.NewRequest(http.MethodPut, "/user/"+"12312", bytes.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))

		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)

		svc.AssertExpectations(t)
	})
}

func TestUserHandler_UpdateUserByIdInvalidJSON(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.UserService)
	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}

	testUserID, _ := uuid.NewV4()

	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})

	tokenString, err := testToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	t.Run("should return bad request when JSON body is invalid", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidJSON := []byte(`{"name": "Updated Name", "email":}`)

		req := httptest.NewRequest(http.MethodPut, "/user/"+testUserID.String(), bytes.NewReader(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))

		router.Put("/user/{id}", hdl.UpdateUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)

		svc.AssertExpectations(t)
	})
}

func TestUserHandler_DeleteUserByID(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})

	tokenString, err := testToken.SignedString([]byte("your-secret-key"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	t.Run("should delete user successfully", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteUserByID", mock.Anything, testUserID.String()).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Delete("/user/{id}", hdl.DeleteUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)
		svc.AssertExpectations(t)
	})

	t.Run("should return 404 when user not found", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteUserByID", mock.Anything, testUserID.String()).Return(service.ErrNotFound).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Delete("/user/{id}", hdl.DeleteUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
		svc.AssertExpectations(t)
	})

	t.Run("should return 500 on internal server error", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteUserByID", mock.Anything, testUserID.String()).Return(errors.New("internal server error")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/"+testUserID.String(), nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router := chi.NewRouter()
		router.Use(JWTMiddleware("your-secret-key", log))
		router.Delete("/user/{id}", hdl.DeleteUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
		svc.AssertExpectations(t)
	})
}

func TestUserHandler_DeleteUserByIdUnauthorized(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := new(mocks.UserService)
	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}

	testUserID, _ := uuid.NewV4()

	t.Run("should return unauthorized when user_id is not present in context", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/user/"+testUserID.String(), nil)
		req.Header.Set("Content-Type", "application/json")

		router := chi.NewRouter()

		router.Delete("/user/{id}", hdl.DeleteUserByID)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)

		svc.AssertExpectations(t)
	})
}

func TestJWTMiddlewareUnexpectedSigningMethod(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()
	router.Use(JWTMiddleware("your-secret-key", log))
	router.Delete("/user/{id}", hdl.DeleteUserByID)

	testUserID, _ := uuid.NewV4()

	testToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})

	testTokenFake := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		// "user_id" отсутствует
	})
	tokenStringFake, err := testTokenFake.SignedString([]byte("your-secret-key"))
	if err != nil {
		t.Fatalf("Failed to sign token: %v", err)
	}

	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return error with invalid token", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodDelete, "/user/"+testUserID.String(), nil)
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
		svc.AssertExpectations(t)
	})

	t.Run("should return invalid user ID in token", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/user/"+testUserID.String(), nil)
		req.Header.Set("Content-Type", "application/json")

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenStringFake,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
		svc.AssertExpectations(t)
	})
}
