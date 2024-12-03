package chat

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/connectionmanager"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allows you to determine whether the server should compress messages.
	EnableCompression: true,
}

type Handler struct {
	Svc interfaces.ChatService
	Log *slog.Logger
}

func (h *Handler) NewChatHandler(r chi.Router) {
	r.Route("/chat", func(r chi.Router) {
		r.Get("/", h.GetChatsByUserID)
		r.Get("/{chatId}", h.GetMessagesByChatID)
		r.Post("/{ownerId}", h.SendMessage)
		r.Handle("/ws/{userId}", http.HandlerFunc(h.HandleWebSocket))
	})
}

// GetChatsByUserID godoc
//
//	@Summary		Get Chats by User ID
//	@Description	Retrieve all chats for a user by their user ID
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"List of chats for the user"
//	@Failure		401	{object}	responseApi.ResponseError	"Unauthorized"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/chat [get]
func (h *Handler) GetChatsByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.chat.GetChatsByUserID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	chats, err := h.Svc.GetChatsByUserID(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to get chats by user id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"chats": chats})
}

// GetMessagesByChatID godoc
//
//	@Summary		Get Messages by Chat ID
//	@Description	Retrieve all messages for a chat by its chat ID
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Param			chatId	path		string	true	"The ID of the chat"
//	@Success		200	{object}	map[string]interface{}	"List of messages for the chat"
//	@Failure		404	{object}	responseApi.ResponseError	"Chat not found"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/chat/{chatId} [get]
func (h *Handler) GetMessagesByChatID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.chat.GetMessagesByChatID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "chatId")

	messages, err := h.Svc.GetMessagesByChatID(context.Background(), id)
	if err != nil {
		h.Log.Error("failed to get messages by chatId", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"messages": messages})
}

// SendMessage godoc
//
//	@Summary		Send a message to a chat
//	@Description	Send a message to a specified chat by its owner ID and user ID
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Param			ownerId	path		string	true	"The ID of the chat owner"
//	@Param			request	body		message.Entity	true	"Message content to send"
//	@Success		200	{object}	map[string]interface{}	"Message sent successfully"
//	@Failure		400	{object}	responseApi.ResponseError	"Bad Request"
//	@Failure		401	{object}	responseApi.ResponseError	"Unauthorized"
//	@Failure		500	{object}	responseApi.ResponseError	"Internal Server Error"
//	@Router			/chat/{ownerId} [post]
func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	const op = "handler.chat.GetMessagesByChatID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	id := chi.URLParam(r, "ownerId")

	var msg message.Entity
	err := render.DecodeJSON(r.Body, &msg)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	err = h.Svc.SendMessage(context.Background(), id, userID, msg)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, "Message sent!")
}

func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("your-secret-key"), nil
	})
	if err != nil || !token.Valid {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token")))
		return
	}

	// Extract the user ID from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid token claims")))
		return
	}

	ownerID, ok := claims["user_id"].(string)
	if !ok {
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("invalid user ID in token")))
		return
	}

	userID := chi.URLParam(r, "userId")

	if ownerID == "" || userID == "" {
		h.Log.Error("ownerId or userId missing", slog.String("owner_id", ownerID), slog.String("user_id", userID))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("ownerId or userId missing")))
		return
	}

	h.Log.Info("New WebSocket connection request", slog.String("owner_id", ownerID), slog.String("user_id", userID))

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Log.Error("WebSocket upgrade failed", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("connection upgrade failed")))
		return
	}
	defer conn.Close()

	cm := connectionmanager.NewConnectionManager()

	cm.AddConnection(ownerID, conn)

	h.Log.Info("Owner added to Connection Manager", slog.String("owner_id", ownerID))
	h.Log.Info("smh", userID)

	cm.AllConnections()

	for {
		messageType, messageData, err := conn.ReadMessage()
		if err != nil {
			h.Log.Warn("Owner disconnected", slog.String("owner_id", ownerID), slogError.Err(err))
			cm.RemoveConnection(ownerID)
			break
		}

		messageText := string(messageData)
		h.Log.Info("Message received", slog.String("owner_id", ownerID), slog.String("message", messageText))

		msg := message.Entity{
			UserID: uuid.Must(uuid.Parse(userID)),
			Text:   messageText,
			Media:  nil,
		}

		h.Log.Info("Created message", slog.String("user_id", msg.UserID.String()), slog.String("message_text", msg.Text))

		if err := conn.WriteMessage(messageType, messageData); err != nil {
			h.Log.Error("Failed to send message", slog.String("owner_id", ownerID), slogError.Err(err))
			cm.RemoveConnection(ownerID)
			break
		}

		h.Log.Info("Message sent to user", slog.String("user_id", userID), slog.String("owner_id", ownerID))
	}

}
