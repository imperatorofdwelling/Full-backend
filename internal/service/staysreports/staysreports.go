package staysreports

import (
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreports"
	service "github.com/imperatorofdwelling/Full-backend/internal/service/file"
	"github.com/imperatorofdwelling/Full-backend/pkg/checkers"
	"golang.org/x/net/context"
)

type Service struct {
	Repo    interfaces.StaysReportsRepo
	FileSvc interfaces.FileService
}

func (s *Service) CreateStaysReports(ctx context.Context, userId, stayId, title, description string, image []byte) error {
	const op = "service.StaysReports.CreateStaysReports"

	fWithPath, err := s.FileSvc.UploadImage(image, service.PngImageType, service.FilePathStaysReportsImages)
	if err != nil {
		return err
	}

	err = s.Repo.CreateStaysReports(ctx, userId, stayId, title, description, fWithPath)
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

func (s *Service) GetStaysReportById(ctx context.Context, userId, id string) (*staysreports.StayReport, error) {
	const op = "service.StaysReports.GetStaysReportById"

	report, err := s.Repo.GetStaysReportById(ctx, userId, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return report, nil
}

func (s *Service) UpdateStaysReports(ctx context.Context, userId, reportId, title, description string, image []byte) (*staysreports.StaysReportEntity, error) {
	const op = "service.StaysReports.UpdateStaysReports"

	reportFound, err := s.Repo.GetStaysReportById(ctx, userId, reportId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if reportFound == nil {
		return nil, fmt.Errorf("%s: report not found with id %s", op, reportId)
	}

	var updatedImagePath string
	if len(image) > 0 {
		imageType, err := checkers.DetectImageType(image)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		updatedImagePath, err = s.FileSvc.UploadImage(image, imageType, service.FilePathStaysReportsImages)
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

	report, err := s.Repo.UpdateStaysReports(ctx, userId, reportId, title, description, updatedImagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return report, nil
}

func (s *Service) DeleteStaysReports(ctx context.Context, userId, reportId string) error {
	const op = "service.StaysReports.DeleteStaysReports"

	reportFound, err := s.Repo.GetStaysReportById(ctx, userId, reportId)
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

	err = s.Repo.DeleteStaysReports(ctx, userId, reportId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
