package usersreports

import (
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/usersreports"
	service "github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"golang.org/x/net/context"
)

type Service struct {
	Repo    interfaces.UsersReportsRepo
	FileSvc interfaces.FileService
}

func (s *Service) CreateUsersReports(ctx context.Context, userId, toBlameId, title, description string, image []byte) error {
	const op = "service.UsersReports.CreateUsersReports"

	imageType, err := checkers.DetectImageType(image)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fWithPath, err := s.FileSvc.UploadImage(image, imageType, service.FilePathUsersReportsImages)
	if err != nil {
		return err
	}

	err = s.Repo.CreateUsersReports(ctx, userId, toBlameId, title, description, fWithPath)
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

func (s *Service) GetUsersReportById(ctx context.Context, userId, id string) (*usersreports.UsersReport, error) {
	const op = "service.UsersReports.GetUsersReportById"

	report, err := s.Repo.GetUsersReportById(ctx, userId, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return report, nil
}

func (s *Service) UpdateUsersReports(ctx context.Context, userId, reportId, title, description string, imageData []byte) (*usersreports.UsersReportEntity, error) {
	const op = "service.UsersReports.UpdateUsersReports"

	reportFound, err := s.Repo.GetUsersReportById(ctx, userId, reportId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if reportFound == nil {
		return nil, fmt.Errorf("%s: report not found with id %s", op, reportId)
	}

	var updatedImagePath string
	if len(imageData) > 0 {
		imageType, err := checkers.DetectImageType(imageData)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		updatedImagePath, err = s.FileSvc.UploadImage(imageData, imageType, service.FilePathUsersReportsImages)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if reportFound.ReportAttach != nil {
			err = s.FileSvc.RemoveFile(*reportFound.ReportAttach)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", op, err)
			}
		}
	} else {
		updatedImagePath = *reportFound.ReportAttach
	}

	report, err := s.Repo.UpdateUsersReports(ctx, userId, reportId, title, description, updatedImagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return report, nil
}

func (s *Service) DeleteUsersReports(ctx context.Context, userId, reportId string) error {
	const op = "service.UsersReports.DeleteUsersReports"

	reportFound, err := s.Repo.GetUsersReportById(ctx, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if reportFound == nil {
		return fmt.Errorf("%s: report not found with id %s", op, reportId)
	}

	if reportFound.ReportAttach != nil {
		err = s.FileSvc.RemoveFile(*reportFound.ReportAttach)
		if err != nil {
			return err
		}
	}

	err = s.Repo.DeleteUsersReports(ctx, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
