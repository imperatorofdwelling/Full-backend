package chat

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces/mocks"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/connectionmanager"
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

func TestChatHandler_NewChatHandler(t *testing.T) {
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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
	log := logger.New(logger.EnvLocal)
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

func TestHandleWebSocket_(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	mockService := mocks.ChatService{}
	cm := connectionmanager.NewConnectionManager()
	hdl := Handler{
		Svc: &mockService,
		Cm:  cm,
		Log: log,
	}
	router := chi.NewRouter()

	router.Get("/chat/ws/{chatId}", hdl.HandleWebSocket)

	t.Run("should successfully handle websocket connection", func(t *testing.T) {
		validToken := "your-valid-token"
		req := httptest.NewRequest(http.MethodGet, "/chat/ws/123?token="+validToken, nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}

func TestHandleWebSocket_ValidToken(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	mockService := mocks.ChatService{}
	cm := connectionmanager.NewConnectionManager()
	hdl := Handler{
		Svc: &mockService,
		Cm:  cm,
		Log: log,
	}
	router := chi.NewRouter()

	token := generateValidToken(t)

	invalidTokenWithoutUserID := generateTokenWithoutUserID()

	router.Get("/chat/ws/{chatId}", hdl.HandleWebSocket)

	t.Run("should successfully upgrade http to websocket", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/chat/ws/123?token="+token, nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return error with empty token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/chat/ws/123?token=", nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should return error without empty token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/chat/ws/123?token="+invalidTokenWithoutUserID, nil)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})

}

func TestHandleWebSocket_Success(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	mockService := mocks.ChatService{}
	cm := connectionmanager.NewConnectionManager()
	hdl := Handler{
		Svc: &mockService,
		Cm:  cm,
		Log: log,
	}
	router := chi.NewRouter()

	token := generateValidToken(t)

	router.Get("/chat/ws/{chatId}", hdl.HandleWebSocket)

	t.Run("should successfully handle websocket connection", func(t *testing.T) {

		chatId := "asdasdasdasd"
		messages := []*message.Entity{
			{Text: "Hello World"},
			{Text: "How are you?"},
		}

		req := httptest.NewRequest(http.MethodGet, "/chat/ws/"+chatId+"?token="+token, nil)
		req.Header.Set("Connection", "Upgrade")
		req.Header.Set("Upgrade", "websocket")

		rr := httptest.NewRecorder()

		server := httptest.NewServer(router)
		defer server.Close()

		url := "ws://" + server.Listener.Addr().String() + "/chat/ws/" + chatId + "?token=" + token
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		defer conn.Close()

		assert.Equal(t, http.StatusOK, rr.Code)

		mockService.On("GetMessagesByChatID", context.Background(), chatId).Return(messages, nil).Once()

	})

	t.Run("should successfully handle websocket connection", func(t *testing.T) {

		chatId := "asdasdasdasd"

		req := httptest.NewRequest(http.MethodGet, "/chat/ws/"+chatId+"?token="+token, nil)
		req.Header.Set("Connection", "Upgrade")
		req.Header.Set("Upgrade", "websocket")

		rr := httptest.NewRecorder()

		server := httptest.NewServer(router)
		defer server.Close()

		url := "ws://" + server.Listener.Addr().String() + "/chat/ws/" + chatId + "?token=" + token
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		defer conn.Close()

		assert.Equal(t, http.StatusOK, rr.Code)

		mockService.On("GetMessagesByChatID", context.Background(), chatId).Return(nil, errors.New("messages error")).Once()

	})
}

func TestHandleWebSocket_MessageHandling_RealManager(t *testing.T) {
	log := logger.New(logger.EnvLocal)
	mockService := &mocks.ChatService{}
	realManager := connectionmanager.NewConnectionManager()
	hdl := Handler{
		Svc: mockService,
		Cm:  realManager,
		Log: log,
	}

	router := chi.NewRouter()
	router.Get("/chat/ws/{chatId}", hdl.HandleWebSocket)

	mockService.On("SendMessageInChat", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	mockService.On("SendMessageInChat", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("Error")).Once()

	t.Run("should successfully process a message", func(t *testing.T) {
		chatId := "chat-id"
		token := generateValidToken(t)
		req := httptest.NewRequest(http.MethodGet, "/chat/ws/"+chatId+"?token="+token, nil)
		req.Header.Set("Connection", "Upgrade")
		req.Header.Set("Upgrade", "websocket")

		server := httptest.NewServer(router)
		defer server.Close()

		url := "ws://" + server.Listener.Addr().String() + "/chat/ws/" + chatId + "?token=" + token
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		defer conn.Close()

		mockService.On("GetMessagesByChatID", mock.Anything, mock.Anything).
			Return([]*message.Entity{}, nil).
			Once()

		message := "Test message"
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	})

	t.Run("should throw error processing a message", func(t *testing.T) {
		chatId := "chat-id"
		token := generateValidToken(t)
		req := httptest.NewRequest(http.MethodGet, "/chat/ws/"+chatId+"?token="+token, nil)
		req.Header.Set("Connection", "Upgrade")
		req.Header.Set("Upgrade", "websocket")

		server := httptest.NewServer(router)
		defer server.Close()

		url := "ws://" + server.Listener.Addr().String() + "/chat/ws/" + chatId + "?token=" + token
		conn, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("Failed to upgrade connection: %v", err)
		}
		defer conn.Close()

		mockService.On("GetMessagesByChatID", mock.Anything, mock.Anything).
			Return([]*message.Entity{}, nil).
			Once()

		message := "Test message"
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))

	})
}

func generateValidToken(t *testing.T) string {
	claims := jwt.MapClaims{
		"user_id": "61f0c404-5cb3-11e7-907b-a6006ad3dba0",
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte("your-secret-key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	return tokenString
}

func generateTokenWithoutUserID() string {
	claims := jwt.MapClaims{
		"some_other_claim": "value",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("your-secret-key"))
	return signedToken
}
