package interfaces

import (
	"context"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/chat"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	"net/http"
)

//go:generate mockery --name ChatRepository
type (
	ChatRepository interface {
		GetChatsByUserID(ctx context.Context, userID string) ([]*chat.Chat, error)
		GetOrCreateChatID(ctx context.Context, userID, otherUserID string) (*string, error)
		GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Entity, error)
		SendMessage(ctx context.Context, senderId, receiverId string, msg message.Entity) error
	}
)

//go:generate mockery --name ChatService
type (
	ChatService interface {
		GetChatsByUserID(ctx context.Context, userID string) ([]*chat.Chat, error)
		GetOrCreateChatID(ctx context.Context, userID, otherUserID string) (*string, error)
		GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Entity, error)
		SendMessage(ctx context.Context, senderId, receiverId string, msg message.Entity) error
	}
)

type (
	ChatHandler interface {
		GetChatsByUserID(w http.ResponseWriter, r *http.Request)
		GetMessagesByChatID(w http.ResponseWriter, r *http.Request)
		SendMessage(w http.ResponseWriter, r *http.Request)
	}
)
