package contracts

import (
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/contracts"
	"golang.org/x/net/context"
	"time"
)

type Service struct {
	Repo interfaces.ContractsRepo
}

func (s *Service) AddContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) error {
	const op = "service.Contracts.AddContract"

	err := s.Repo.AddContract(ctx, userId, stayId, dateStart, dateEnd)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Service) UpdateContract(ctx context.Context, userId, stayId string, dateStart, dateEnd time.Time) error {
	const op = "service.Contracts.UpdateContract"

	err := s.Repo.UpdateContract(ctx, userId, stayId, dateStart, dateEnd)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Service) GetAllContracts(ctx context.Context, userId string) ([]contracts.ContractEntity, error) {
	const op = "service.Contracts.GetAllContracts"

	allContracts, err := s.Repo.GetAllContracts(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return allContracts, nil
}
