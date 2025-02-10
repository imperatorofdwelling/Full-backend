package message

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	_ "github.com/imperatorofdwelling/Full-backend/internal/domain/models/response"
	mw "github.com/imperatorofdwelling/Full-backend/internal/middleware"
	responseApi "github.com/imperatorofdwelling/Full-backend/internal/utils/response"
	"github.com/imperatorofdwelling/Full-backend/pkg/logger/slogError"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	Svc interfaces.MessageService
	Log *slog.Logger
}

func (h *Handler) NewMessageHandler(r chi.Router) {
	r.Route("/message", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(mw.WithAuth)
			r.Get("/", h.GetMessagesByUserID)
			r.Get("/{messageId}", h.GetMessageByMessageID)
			r.Put("/{messageId}", h.UpdateMessageByID)
			r.Delete("/{messageId}", h.DeleteMessageByID)
		})
	})
}

// GetMessagesByUserID godoc
//
//	@Summary		Get Messages
//	@Description	Retrieve all messages for a user by their user ID
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	    message.Entity	"List of messages for the user"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
//	@Router			/message [get]
func (h *Handler) GetMessagesByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.message.GetMessagesByUserID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	userID, ok := r.Context().Value(mw.UserIdKey).(string)
	if !ok {
		h.Log.Error("user ID not found in context")
		responseApi.WriteError(w, r, http.StatusUnauthorized, slogError.Err(errors.New("user not logged in")))
		return
	}

	messages, err := h.Svc.GetMessagesByUserID(context.Background(), userID)
	if err != nil {
		h.Log.Error("failed to get messages by user id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"messages": messages})
}

// GetMessageByMessageID godoc
//
//	@Summary		Get Message by Message ID
//	@Description	Retrieve a single message by its message ID
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			messageId	path		string	true	"The ID of the message"
//	@Success		200	{object}	message.Message	"The message details"
//	@Failure		404	{object}	response.ResponseError	"Message not found"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
//	@Router			/message/{messageId} [get]
func (h *Handler) GetMessageByMessageID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.message.GetMessageByMessageID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "messageId")

	msg, err := h.Svc.GetMessageByMessageID(context.Background(), id)
	if err != nil {
		h.Log.Error("failed to get message by messageId", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	if msg == nil {
		h.Log.Error("message not found")
		responseApi.WriteError(w, r, http.StatusNotFound, fmt.Errorf("message not found"))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"message": msg})
}

// UpdateMessageByID godoc
//
//	@Summary		Update an Existing Message
//	@Description	Update an existing message by message ID
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			messageId	path		string	true	"The ID of the message"
//	@Param			request	body		message.Entity	true	"Updated message content"
//	@Success		200	{object}	message.Entity	"Updated message"
//	@Failure		400	{object}	response.ResponseError	"Bad Request"
//	@Failure		401	{object}	response.ResponseError	"Unauthorized"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
//	@Router			/message/{messageId} [put]
func (h *Handler) UpdateMessageByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.message.UpdateMessageByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "messageId")

	var msg message.Entity
	err := render.DecodeJSON(r.Body, &msg)
	if err != nil {
		h.Log.Error("%s: %v", op, err)
		responseApi.WriteError(w, r, http.StatusBadRequest, slogError.Err(err))
		return
	}

	updatedMsg, err := h.Svc.UpdateMessageByID(context.Background(), id, msg)
	if err != nil {
		h.Log.Error("failed to update message", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusOK, map[string]interface{}{"message": updatedMsg})
}

// DeleteMessageByID godoc
//
//	@Summary		Delete Message by Message ID
//	@Description	Delete a message by its message ID
//	@Tags			messages
//	@Accept			json
//	@Produce		json
//	@Param			messageId	path		string	true	"The ID of the message"
//	@Success		204	{object}	string 	"Message deleted successfully"
//	@Failure		404	{object}	response.ResponseError	"Message not found"
//	@Failure		500	{object}	response.ResponseError	"Internal Server Error"
//	@Router			/message/{messageId} [delete]
func (h *Handler) DeleteMessageByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.message.DeleteMessageByID"

	h.Log = h.Log.With(
		slog.String("op", op),
		slog.String("request_id", middleware.GetReqID(r.Context())),
	)

	id := chi.URLParam(r, "messageId")

	err := h.Svc.DeleteMessageByID(context.Background(), id)
	if err != nil {
		h.Log.Error("failed to delete message by message id", slogError.Err(err))
		responseApi.WriteError(w, r, http.StatusInternalServerError, slogError.Err(err))
		return
	}

	responseApi.WriteJson(w, r, http.StatusNoContent, "message successfully deleted")
}
