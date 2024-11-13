package chat

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.ChatService
	Log *slog.Logger
}

func (h *Handler) NewChatHandler(r chi.Router) {
	r.Route("/chat", func(r chi.Router) {
		r.Get("/", h.GetChatsByUserID)
		r.Get("/{chatId}", h.GetMessagesByChatID)
		r.Post("/{ownerId}", h.SendMessage)
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

	responseApi.WriteJson(w, r, http.StatusOK, "Message sent!")
}
