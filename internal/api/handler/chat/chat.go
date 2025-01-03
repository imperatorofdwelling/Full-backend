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
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/chat"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/connectionmanager"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
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
	Cm  *connectionmanager.ConnectionManager
}

func (h *Handler) NewChatHandler(r chi.Router) {
	r.Route("/chat", func(r chi.Router) {
		r.Use(mw.WithAuth)
		r.Get("/", h.GetChatsByUserID)
		r.Get("/{chatId}", h.GetMessagesByChatID)
		r.Post("/{ownerId}", h.SendMessage)
		r.Handle("/ws/{chatId}", http.HandlerFunc(h.HandleWebSocket))
	})
}

// GetChatsByUserID godoc
//
//	@Summary		Get Chats
//	@Description	Retrieve all chats for a user by their user ID
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		chat.Chat	"List of chats for the user"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
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
//	@Success		200	{object}	message.Entity	"List of messages for the chat"
//	@Failure		404	{object}	response.ResponseError	"Chat not found"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
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
//	@Summary		Create a message
//	@Description	Send a message to a specified chat by its owner ID and user ID
//	@Tags			chats
//	@Accept			json
//	@Produce		json
//	@Param			ownerId	path	string	true	"The ID of the chat owner"
//	@Param			request	body	message.Entity	true	"Message content to send"
//	@Success		200	{string}	string	"Message sent!"
//	@Failure		400	{object}	response.ResponseError	"Bad Request"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
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

// HandleWebSocket godoc
//
//	@Summary		Establishes a WebSocket connection for chat
//	@Description	Handles WebSocket connections, retrieves chat history, and supports real-time messaging.
//	@Tags			webSocket
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"JWT Token for authentication"
//	@Param			chatId	path		string	true	"Chat ID to retrieve messages from"
//	@Success		101	{string}	string	"WebSocket connection established"
//	@Failure		400	{object}	response.ResponseError	"Bad Request - Missing or invalid token"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized - Invalid token or claims"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
//	@Router			/chat/ws/{chatId} [get]
func (h *Handler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Getting and checking token
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

	if ownerID == "" {
		h.Log.Error("ownerId or userId missing", slog.String("owner_id", ownerID))
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(errors.New("ownerId or userId missing")))
		return
	}

	// Upgrading websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Log.Error("WebSocket upgrade failed", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("connection upgrade failed")))
		return
	}
	defer conn.Close()

	// Connection manager
	h.Cm.AddConnection(ownerID, conn)
	h.Log.Info("Owner added to Connection Manager", slog.String("owner_id", ownerID))
	h.Cm.AllConnections()

	// Getting message history
	chatId := chi.URLParam(r, "chatId")
	h.Log.Info("id", "id", chatId)
	messages, err := h.Svc.GetMessagesByChatID(context.Background(), chatId)
	if err != nil {
		h.Log.Error("Error while getting messages by chat id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("failed to fetch messages")))
		return
	}

	for _, msg := range messages {
		msgData := []byte(msg.Text)
		if err := conn.WriteMessage(websocket.TextMessage, msgData); err != nil {
			h.Log.Error("Failed to send message history", slogError.Err(err))
			h.Cm.RemoveConnection(ownerID)
			return
		}
	}

	// Infinite loop for websocket
	for {
		_, messageData, err := conn.ReadMessage()
		if err != nil {
			h.Log.Warn("Owner disconnected", slog.String("owner_id", ownerID), slogError.Err(err))
			h.Cm.RemoveConnection(ownerID)
			break
		}

		messageText := string(messageData)
		h.Log.Info("Message received", slog.String("owner_id", ownerID), slog.String("message", messageText))

		// Creating a message entity
		msg := message.Entity{
			UserID: uuid.Must(uuid.Parse(ownerID)),
			Text:   messageText,
			Media:  nil,
		}

		h.Log.Info("Created message", slog.String("user_id", msg.UserID.String()), slog.String("message_text", msg.Text))

		// Saving the message to the database
		err = h.Svc.SendMessageInChat(context.Background(), chatId, ownerID, msg)
		if err != nil {
			h.Log.Error("Failed to save message in chat", slog.String("chat_id", chatId), slogError.Err(err))
			continue
		}

		// Broadcasting the message to all clients
		h.Cm.BroadcastMessage(ownerID, messageData)
	}
}
