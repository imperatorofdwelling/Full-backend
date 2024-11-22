package chat

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/chat"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
)

type Service struct {
	Repo interfaces.ChatRepository
}

func (s *Service) GetChatsByUserID(ctx context.Context, userID string) ([]*chat.Chat, error) {
	const op = "service.chat.GetChatsByUserID"

	chats, err := s.Repo.GetChatsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return chats, nil
}

func (s *Service) GetOrCreateChatID(ctx context.Context, userID, otherUserID string) (*string, error) {
	const op = "service.chat.GetOrCreateChatID"

	chats, err := s.Repo.GetOrCreateChatID(ctx, userID, otherUserID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return chats, nil
}

func (s *Service) GetMessagesByChatID(ctx context.Context, chatID string) ([]*message.Entity, error) {
	const op = "service.chat.GetMessagesByChatID"

	messages, err := s.Repo.GetMessagesByChatID(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return messages, nil
}

func (s *Service) SendMessage(ctx context.Context, senderId string, receiverId string, msg message.Entity) error {
	const op = "service.chat.SendMessage"

	err := s.Repo.SendMessage(ctx, senderId, receiverId, msg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
