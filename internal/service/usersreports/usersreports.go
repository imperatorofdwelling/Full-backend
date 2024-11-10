package usersreports

import (
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	"golang.org/x/net/context"
)

type Service struct {
	Repo interfaces.UsersReportsRepo
}

func (s *Service) CreateUsersReports(ctx context.Context, userId, toBlameId, title, description string) error {
	const op = "service.UsersReports.CreateUsersReports"

	err := s.Repo.CreateUsersReports(ctx, userId, toBlameId, title, description)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetAllUsersReports(ctx context.Context, userId string) ([]usersreports.UsersReportEntity, error) {
	const op = "service.UsersReports.GetAllUsersReports"

	reports, err := s.Repo.GetAllUsersReports(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return reports, nil
}

func (s *Service) UpdateUsersReports(ctx context.Context, userId, toBlameId, title, description string) (*usersreports.UsersReportEntity, error) {
	const op = "service.UsersReports.UpdateUsersReports"

	report, err := s.Repo.UpdateUsersReports(ctx, userId, toBlameId, title, description)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return report, nil
}

func (s *Service) DeleteUsersReports(ctx context.Context, userId, reportId string) error {
	const op = "service.UsersReports.DeleteUsersReports"

	err := s.Repo.DeleteUsersReports(ctx, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
