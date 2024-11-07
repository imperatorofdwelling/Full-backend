package advantage

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/advantage"
	"github.com/imperatorofdwelling/Full-backend/internal/repo"
	errs "github.com/imperatorofdwelling/Full-backend/internal/service"
	service "github.com/imperatorofdwelling/Full-backend/internal/service/file"
)

type Service struct {
	Repo    interfaces.AdvantageRepo
	FileSvc interfaces.FileService
}

func (s *Service) CreateAdvantage(ctx context.Context, adv *advantage.AdvantageEntity) error {
	const op = "service.AdvantageService.CreateAdvantage"

	isExists, err := s.Repo.CheckAdvantageIfExists(ctx, adv.Title)
	if err != nil {
		return err
	}

	if isExists {
		return fmt.Errorf("%s: %s already exists", op, adv.Title)
	}

	fWithPath, err := s.FileSvc.UploadImage(adv.Image, service.SvgImageType, service.FilePathAdvantages)
	if err != nil {
		return err
	}

	err = s.Repo.CreateAdvantage(ctx, adv.Title, fWithPath)
	if err != nil {
		errF := s.FileSvc.RemoveFile(fWithPath)
		if errF != nil {
			return errF
		}
		return err
	}

	return nil
}

func (s *Service) RemoveAdvantage(ctx context.Context, advID uuid.UUID) error {
	const op = "service.AdvantageService.RemoveAdvantage"

	adv, err := s.Repo.FindAdvantageByID(ctx, advID)
	if err != nil {
		return err
	}

	if adv.Image != "" {
		err = s.FileSvc.RemoveFile(adv.Image)
		if err != nil {
			return err
		}
	}

	err = s.Repo.RemoveAdvantage(ctx, advID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllAdvantages(ctx context.Context) ([]advantage.Advantage, error) {
	const op = "service.advantage.GetAllAdvantages"

	adv, err := s.Repo.GetAllAdvantages(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return adv, nil
}

func (s *Service) UpdateAdvantageByID(ctx context.Context, id uuid.UUID, adv *advantage.AdvantageEntity) (advantage.Advantage, error) {
	const op = "service.advantage.UpdateAdvantageByID"

	advFound, err := s.Repo.FindAdvantageByID(ctx, id)
	if err != nil {
		return advantage.Advantage{}, fmt.Errorf("%s: %w", op, err)
	}

	if advFound.ID == uuid.Nil {
		return advantage.Advantage{}, fmt.Errorf("%s: %s", op, repo.ErrUserNotFound)
	}

	var newAdv advantage.Advantage

	if adv.Image != nil {
		image, err := s.FileSvc.UploadImage(adv.Image, service.SvgImageType, service.FilePathAdvantages)
		if err != nil {
			return advantage.Advantage{}, err
		}

		newAdv.Image = image

		err = s.FileSvc.RemoveFile(advFound.Image)
		if err != nil {
			return advantage.Advantage{}, err
		}
	} else {
		newAdv.Image = advFound.Image
	}

	if adv.Title != "" {
		newAdv.Title = adv.Title
	} else {
		newAdv.Title = advFound.Title
	}

	err = s.Repo.UpdateAdvantageByID(ctx, id, &newAdv)
	if err != nil {
		return advantage.Advantage{}, fmt.Errorf("%s: %w", op, err)
	}

	advUpdated, err := s.Repo.FindAdvantageByID(ctx, id)
	if err != nil {
		return advantage.Advantage{}, fmt.Errorf("%s: %w", op, err)
	}

	return *advUpdated, nil
}

func (s *Service) GetAdvantageByID(ctx context.Context, id uuid.UUID) (*advantage.Advantage, error) {
	const op = "service.advantage.GetAdvantageByID"

	adv, err := s.Repo.FindAdvantageByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if adv == nil {
		return nil, fmt.Errorf("%s: %w", op, errs.ErrAdvantageNotFound)
	}

	return adv, nil
}
