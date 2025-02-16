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
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/newPassword"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
	"time"
)

func TestUserHandler_NewUserHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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

func TestUserHandler_DeleteUserPfp(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	router := chi.NewRouter()

	fakeUUID, _ := uuid.NewV4()

	t.Run("should be no error", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteUserPfp", mock.Anything, fakeUUID).Return(nil).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/profile/picture/"+fakeUUID.String(), nil)

		router.HandleFunc("/user/profile/picture/{id}", hdl.DeleteUserPfp)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)
	})

	t.Run("should be parsing uuid error", func(t *testing.T) {
		r := httptest.NewRecorder()

		invalidUUID := "invalid"

		req := httptest.NewRequest(http.MethodDelete, "/user/profile/picture/"+invalidUUID, nil)

		router.HandleFunc("/user/profile/picture/{id}", hdl.DeleteUserPfp)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should be error deleting image", func(t *testing.T) {
		r := httptest.NewRecorder()

		svc.On("DeleteUserPfp", mock.Anything, fakeUUID).Return(errors.New("fail")).Once()

		req := httptest.NewRequest(http.MethodDelete, "/user/profile/picture/"+fakeUUID.String(), nil)

		router.HandleFunc("/user/profile/picture/{id}", hdl.DeleteUserPfp)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestUserHandler_GetUserByID(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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

		assert.Equal(t, http.StatusOK, r.Code)

		//assert.Equal(t, expected.ID, actual.ID)
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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

func TestUserHandler_UpdateUserPasswordByEmail_Payload_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	invalidJSON := `{"Email": "something@example.com", "Password": 123}`

	emptyPayload := newPassword.NewPassword{
		Email: "",
	}

	emptyPayloadBytes, err := json.Marshal(emptyPayload)
	if err != nil {
		t.Fatalf("Failed to marshal empty payload: %v", err)
	}

	router := chi.NewRouter()
	router.Put("/password", hdl.UpdateUserPasswordByEmail)

	t.Run("should return bad request for invalid JSON payload", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", bytes.NewReader([]byte(invalidJSON)))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request for nil payload", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", nil)
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return bad request for empty values in payload", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", bytes.NewReader(emptyPayloadBytes))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestUserHandler_UpdateUserPasswordByEmail_CheckUserPassword_SVC_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	payloadFull := newPassword.NewPassword{
		Email:    "something@example.com",
		Password: "something",
	}

	payloadBytesFull, err := json.Marshal(payloadFull)

	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	router := chi.NewRouter()
	router.Put("/password", hdl.UpdateUserPasswordByEmail)

	t.Run("should return error while checking CheckUserPassword svc", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", bytes.NewReader(payloadBytesFull))
		req.Header.Set("Content-Type", "application/json")

		svc.On("CheckUserPassword", mock.Anything, mock.Anything).Return(errors.New("CheckUserPassword_SVC_Error")).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestUserHandler_UpdateUserPasswordByEmail_UpdateUserPasswordByEmail_SVC_Error_Internal(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	payloadFull := newPassword.NewPassword{
		Email:    "something@example.com",
		Password: "something",
	}

	payloadBytesFull, err := json.Marshal(payloadFull)

	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	router := chi.NewRouter()
	router.Put("/password", hdl.UpdateUserPasswordByEmail)

	t.Run("should return internal server error on svc UpdateUserPasswordByEmail", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", bytes.NewReader(payloadBytesFull))
		req.Header.Set("Content-Type", "application/json")

		svc.On("CheckUserPassword", mock.Anything, mock.Anything).Return(nil).Once()
		svc.On("UpdateUserPasswordByEmail", mock.Anything, mock.Anything).Return(errors.New("update user password error")).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestUserHandler_UpdateUserPasswordByEmail_UpdateUserPasswordByEmail_SVC_Error_ErrNotFound(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	payloadFull := newPassword.NewPassword{
		Email:    "something@example.com",
		Password: "something",
	}

	payloadBytesFull, err := json.Marshal(payloadFull)

	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	router := chi.NewRouter()
	router.Put("/password", hdl.UpdateUserPasswordByEmail)

	t.Run("should return internal server error on svc UpdateUserPasswordByEmail", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", bytes.NewReader(payloadBytesFull))
		req.Header.Set("Content-Type", "application/json")

		svc.On("CheckUserPassword", mock.Anything, mock.Anything).Return(nil).Once()
		svc.On("UpdateUserPasswordByEmail", mock.Anything, mock.Anything).Return(service.ErrNotFound).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestUserHandler_UpdateUserPasswordByEmail_UpdateUserPasswordByEmail_SVC_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.UserService{}
	hdl := UserHandler{
		Log: log,
		Svc: &svc,
	}

	payloadFull := newPassword.NewPassword{
		Email:    "something@example.com",
		Password: "something",
	}

	payloadBytesFull, err := json.Marshal(payloadFull)

	if err != nil {
		t.Fatalf("Failed to marshal payload: %v", err)
	}

	router := chi.NewRouter()
	router.Put("/password", hdl.UpdateUserPasswordByEmail)

	t.Run("should return status ok", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPut, "/password", bytes.NewReader(payloadBytesFull))
		req.Header.Set("Content-Type", "application/json")

		svc.On("CheckUserPassword", mock.Anything, mock.Anything).Return(nil).Once()
		svc.On("UpdateUserPasswordByEmail", mock.Anything, mock.Anything).Return(nil).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestUserHandler_UpdateUserByIdUnauthorized(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
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

func TestUserHandler_GetUserPfp(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := new(mocks.UserService)

	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}
	testUserID, _ := uuid.NewV4()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID,
	})
	tokenString, _ := token.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(JWTMiddleware("your-secret-key", log))
	router.Get("/profile/picture", hdl.GetUserPfp)

	t.Run("should return 200 OK", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/profile/picture", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetUserPfp", mock.Anything, mock.Anything).Return(mock.Anything, nil).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})

	t.Run("should return 500 Internal Server Error", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/profile/picture", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("GetUserPfp", mock.Anything, mock.Anything).Return(mock.Anything, errors.New("error with GetUserPfp")).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestUserHandler_CreateUserPfp_Errors(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := new(mocks.UserService)

	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}
	testUserID, _ := uuid.NewV4()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID,
	})
	tokenString, _ := token.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(JWTMiddleware("your-secret-key", log))
	router.Post("/profile/picture", hdl.CreateUserPfp)

	t.Run("should return 400 Bad Request for invalid multipart body", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/profile/picture", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		req.Header.Set("Content-Type", "multipart/form-data")
		req.Body = ioutil.NopCloser(bytes.NewReader([]byte("invalid content")))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return 400 Bad Request for missing file in multipart body", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/profile/picture", body)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		req.Header.Set("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return 400 Bad Request with the incorrect image type", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/svg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/profile/picture", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})

	t.Run("should return 500 Internal Server Error with the blank image", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.svg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/profile/picture", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestUserHandler_CreateUserPfp_Svc(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := new(mocks.UserService)

	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}
	testUserID, _ := uuid.NewV4()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID,
	})
	tokenString, _ := token.SignedString([]byte("your-secret-key"))

	router := chi.NewRouter()
	router.Use(JWTMiddleware("your-secret-key", log))
	router.Post("/profile/picture", hdl.CreateUserPfp)

	t.Run("should return SVC error", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.jpg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/profile/picture", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUserPfp", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("error with CreateUserPfp")).Once()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})

	t.Run("should return SVC success", func(t *testing.T) {
		r := httptest.NewRecorder()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		boundary := writer.Boundary()

		partHeader := make(textproto.MIMEHeader)
		partHeader.Set("Content-Type", "image/jpeg")
		partHeader.Set("Content-Disposition", `form-data; name="image"; filename="test.jpg"`)
		part, err := writer.CreatePart(partHeader)
		if err != nil {
			t.Fatal(err)
		}

		jpegContent := []byte{
			0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46, 0x00, 0x01,
			0x01, 0x01, 0x00, 0x60, 0x00, 0x60, 0x00, 0x00, 0xFF, 0xD9,
		}
		part.Write(jpegContent)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/profile/picture", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		svc.On("CreateUserPfp", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusCreated, r.Code)
	})
}

func TestUserHandler_UserPfp_401(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := new(mocks.UserService)

	hdl := UserHandler{
		Log: log,
		Svc: svc,
	}

	router := chi.NewRouter()
	router.Get("/profile/picture", hdl.GetUserPfp)
	router.Post("/profile/picture", hdl.CreateUserPfp)

	tokenfake := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenStringFake, _ := tokenfake.SignedString([]byte("your-secret-key"))

	t.Run("should return 401 Not logger in - Get", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/profile/picture", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenStringFake,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})

	t.Run("should return 401 Not logger in - Post", func(t *testing.T) {
		r := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/profile/picture", nil)

		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenStringFake,
		}
		req.AddCookie(cookie)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}
