package chat

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/config"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestChatHandler_NewChatHandler(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	t.Run("should be no errors", func(t *testing.T) {
		hdl.NewChatHandler(router)
	})
}

func TestChatHandler_GetChatsByUserID_UserID_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetChatsByUserID)

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestChatHandler_GetChatsByUserID_Svc_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetChatsByUserID)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		ctx := context.WithValue(req.Context(), "user_id", testUserID.String())
		req = req.WithContext(ctx)
		svc.On("GetChatsByUserID", mock.Anything, testUserID.String()).Return(nil, errors.New("service error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestChatHandler_GetChatsByUserID_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetChatsByUserID)

	testUserID, _ := uuid.NewV4()
	testToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"user_id": testUserID.String(),
	})
	tokenString, _ := testToken.SignedString([]byte("your-secret-key"))

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cookie := &http.Cookie{
			Name:  "jwt-token",
			Value: tokenString,
		}
		req.AddCookie(cookie)

		ctx := context.WithValue(req.Context(), "user_id", testUserID.String())
		req = req.WithContext(ctx)
		svc.On("GetChatsByUserID", mock.Anything, testUserID.String()).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestChatHandler_GetMessagesByChatID_ChatID_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessagesByChatID)

	t.Run("should be chat id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		svc.On("GetMessagesByChatID", mock.Anything, mock.Anything).Return(nil, errors.New("service error"))

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
	})
}

func TestChatHandler_GetMessagesByChatID_Succcess(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/", hdl.GetMessagesByChatID)

	t.Run("should be chat id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		svc.On("GetMessagesByChatID", mock.Anything, mock.Anything).Return(nil, nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}

func TestChatHandler_SendMessage_UserID_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Post("/", hdl.SendMessage)

	t.Run("should be user id error", func(t *testing.T) {
		r := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/", nil)

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusUnauthorized, r.Code)
	})
}

func TestChatHandler_SendMessage_Decode_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Post("/", hdl.SendMessage)

	t.Run("should return bad request on invalid JSON", func(t *testing.T) {
		invalidJSON := "invalid-json-format"
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(invalidJSON))

		ctx := context.WithValue(req.Context(), "user_id", "some-valid-user-id")
		req = req.WithContext(ctx)

		r := httptest.NewRecorder()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestChatHandler_SendMessage_Svc_Error(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Post("/", hdl.SendMessage)

	t.Run("should return internal server error on service failure", func(t *testing.T) {
		payload := `{"UserId": "some-valid-user-id", "text": "Hello, World!"}`

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))

		ctx := context.WithValue(req.Context(), "user_id", "some-valid-user-id")
		req = req.WithContext(ctx)

		svc.On("SendMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("internal server error"))

		r := httptest.NewRecorder()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusBadRequest, r.Code)
	})
}

func TestChatHandler_SendMessage_Success(t *testing.T) {
	config.GlobalEnv = config.LocalEnv

	log := logger.New()
	svc := mocks.ChatService{}
	hdl := Handler{
		Svc: &svc,
		Log: log,
	}
	router := chi.NewRouter()

	router.Post("/", hdl.SendMessage)

	t.Run("should return internal server error on service failure", func(t *testing.T) {
		payload := `{"UserId": "some-valid-user-id", "text": "Hello, World!"}`

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payload))

		ctx := context.WithValue(req.Context(), "user_id", "some-valid-user-id")
		req = req.WithContext(ctx)

		svc.On("SendMessage", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		r := httptest.NewRecorder()

		router.ServeHTTP(r, req)

		assert.Equal(t, http.StatusOK, r.Code)
	})
}
