package interfaces

import (
	"context"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
	"net/http"
)

//go:generate mockery --name MessageRepository
type (
	MessageRepository interface {
		GetMessagesByUserID(ctx context.Context, userId string) ([]*message.Entity, error)
		GetMessageByMessageID(ctx context.Context, messageId string) (*message.Message, error)
		UpdateMessageByID(ctx context.Context, messageId string, msg message.Entity) error
		DeleteMessageByID(ctx context.Context, messageId string) error
	}
)

//go:generate mockery --name MessageService
type (
	MessageService interface {
		GetMessagesByUserID(ctx context.Context, userId string) ([]*message.Entity, error)
		GetMessageByMessageID(ctx context.Context, messageId string) (*message.Message, error)
		UpdateMessageByID(ctx context.Context, messageId string, msg message.Entity) (*message.Message, error)
		DeleteMessageByID(ctx context.Context, messageId string) error
	}
)

type (
	MessageHandler interface {
		GetMessagesByUserID(http.ResponseWriter, *http.Request)
		GetMessageByMessageID(http.ResponseWriter, *http.Request)
		UpdateMessageByID(http.ResponseWriter, *http.Request)
		DeleteMessageByID(http.ResponseWriter, *http.Request)
	}
)
