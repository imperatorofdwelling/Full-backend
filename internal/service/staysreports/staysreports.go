package staysreports

import (
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	"golang.org/x/net/context"
)

type Service struct {
	Repo interfaces.StaysReportsRepo
}

func (s *Service) CreateStaysReports(ctx context.Context, userId, stayId, title, description string) error {
	const op = "service.StaysReports.CreateStaysReports"

	err := s.Repo.CreateStaysReports(ctx, userId, stayId, title, description)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetAllStaysReports(ctx context.Context, userId string) ([]staysreports.StaysReportEntity, error) {
	const op = "service.StaysReports.GetAllStaysReports"

	reports, err := s.Repo.GetAllStaysReports(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return reports, nil
}

func (s *Service) UpdateStaysReports(ctx context.Context, userId, reportId, title, description string) (*staysreports.StaysReportEntity, error) {
	const op = "service.StaysReports.UpdateStaysReports"

	report, err := s.Repo.UpdateStaysReports(ctx, userId, reportId, title, description)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return report, nil
}

func (s *Service) DeleteStaysReports(ctx context.Context, userId, reportId string) error {
	const op = "service.StaysReports.DeleteStaysReports"

	err := s.Repo.DeleteStaysReports(ctx, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
