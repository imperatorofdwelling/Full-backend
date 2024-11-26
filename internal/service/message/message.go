package message

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/message"
)

type Service struct {
	Repo interfaces.MessageRepository
}

func (s *Service) GetMessagesByUserID(ctx context.Context, userId string) ([]*message.Entity, error) {
	const op = "service.message.GetMessagesByUserID"

	messages, err := s.Repo.GetMessagesByUserID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return messages, nil
}

func (s *Service) GetMessageByMessageID(ctx context.Context, messageId string) (*message.Message, error) {
	const op = "service.message.GetMessageByMessageID"

	msg, err := s.Repo.GetMessageByMessageID(ctx, messageId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if msg == nil {
		return nil, fmt.Errorf("%s: message not found", op)
	}

	return msg, nil
}

func (s *Service) UpdateMessageByID(ctx context.Context, messageId string, msg message.Entity) (*message.Message, error) {
	const op = "service.message.UpdateMessageByID"

	err := s.Repo.UpdateMessageByID(ctx, messageId, msg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	updatedMsg, err := s.Repo.GetMessageByMessageID(ctx, messageId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if updatedMsg == nil {
		return nil, fmt.Errorf("%s: updated message not found", op)
	}

	return updatedMsg, nil
}

func (s *Service) DeleteMessageByID(ctx context.Context, messageId string) error {
	const op = "service.message.DeleteMessageByID"

	err := s.Repo.DeleteMessageByID(ctx, messageId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
