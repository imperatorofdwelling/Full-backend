package searchhistory

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/searchhistory"
)

type Service struct {
	Repo interfaces.SearchHistoryRepo
}

func (s *Service) AddHistory(ctx context.Context, userId, name string) error {
	const op = "service.SearchHistory.AddHistory"

	err := s.Repo.AddHistory(ctx, userId, name)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetAllHistoryByUserId(ctx context.Context, userId string) ([]searchhistory.SearchHistory, error) {
	const op = "service.SearchHistory.GetAllHistoryByUserId"

	history, err := s.Repo.GetAllHistoryByUserId(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return history, nil
}
