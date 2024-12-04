package message

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	uuid2 "github.com/google/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestMessageHandler_NewMessageHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewMessageHandler(router)
	})
}

func TestMessageHandler_GetMessagesByMessageID_UserIdError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessageByMessageID)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		svc.On("GetMessageByMessageID", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNotFound, r.Code)
	})
}

func TestMessageHandler_GetMessagesByUserID_UserIdError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessagesByUserID)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestMessageHandler_GetMessagesByUserID_SvcError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessagesByUserID)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return internal server error on service failure", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		ctx := context.WithValue(req.Context(), "user_id", testUserID.String())
		req = req.WithContext(ctx)

		svc.On("GetMessagesByUserID", mock.Anything, testUserID.String()).Return(nil, errors.New("service error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestMessageHandler_GetMessagesByUserID_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessagesByUserID)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should return internal server error on service failure", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		ctx := context.WithValue(req.Context(), "user_id", testUserID.String())
		req = req.WithContext(ctx)

		svc.On("GetMessagesByUserID", mock.Anything, testUserID.String()).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestMessageHandler_GetMessagesByMessageID_InternalServerError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessageByMessageID)

	t.Run("should be error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		svc.On("GetMessageByMessageID", mock.Anything, mock.Anything).Return(nil, errors.New("not authorized"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestMessageHandler_GetMessagesByMessageID_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/{messageId}", hdl.GetMessageByMessageID)

	t.Run("should be no errors", func(t *testing.T) {
		r := httptest.NewRecorder()

		random, _ := uuid.NewV4()
		req := httptest.NewRequest(http.MethodGet, "/"+random.String(), nil)

		expectedMessage := &message.Message{ID: uuid2.UUID(random), Text: "Hello World"}
		svc.On("GetMessageByMessageID", mock.Anything, random.String()).Return(expectedMessage, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestMessageHandler_UpdateMessageByID_DecodeError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Put("/{messageId}", hdl.UpdateMessageByID)

	t.Run("should return bad request on decode error", func(t *testing.T) {
		r := httptest.NewRecorder()
		testId, _ := uuid.NewV4()

		invalidJSON := "invalid-json-format"
		req := httptest.NewRequest(http.MethodPut, "/"+testId.String(), strings.NewReader(invalidJSON))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestMessageHandler_UpdateMessageByID_SvcError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Put("/{messageId}", hdl.UpdateMessageByID)

	t.Run("should return bad request on decode error", func(t *testing.T) {
		r := httptest.NewRecorder()
		testId, _ := uuid.NewV4()

		payload := `{"messageId": "123456789123456"}`
		req := httptest.NewRequest(http.MethodPut, "/"+testId.String(), strings.NewReader(payload))
		svc.On("UpdateMessageByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("svc error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestMessageHandler_UpdateMessageByID_SvcSuccess(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Put("/{messageId}", hdl.UpdateMessageByID)

	t.Run("should return bad request on decode error", func(t *testing.T) {
		r := httptest.NewRecorder()
		testId, _ := uuid.NewV4()

		payload := `{"messageId": "123456789123456"}`
		req := httptest.NewRequest(http.MethodPut, "/"+testId.String(), strings.NewReader(payload))
		svc.On("UpdateMessageByID", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestMessageHandler_DeleteMessageByID_MessageIdError(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Delete("/{messageId}", hdl.DeleteMessageByID)

	t.Run("should return not found on invalid message ID", func(t *testing.T) {
		r := httptest.NewRecorder()

		testId, _ := uuid.NewV4()
		req := httptest.NewRequest(http.MethodDelete, "/"+testId.String(), nil)

		svc.On("DeleteMessageByID", mock.Anything, testId.String()).Return(errors.New("message not found"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestMessageHandler_DeleteMessageByID_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.MessageService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Delete("/{messageId}", hdl.DeleteMessageByID)

	t.Run("should return not found on invalid message ID", func(t *testing.T) {
		r := httptest.NewRecorder()

		testId, _ := uuid.NewV4()
		req := httptest.NewRequest(http.MethodDelete, "/"+testId.String(), nil)

		svc.On("DeleteMessageByID", mock.Anything, testId.String()).Return(nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusNoContent, r.Code)
	})
}
