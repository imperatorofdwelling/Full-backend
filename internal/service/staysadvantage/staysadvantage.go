package staysadvantage

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysadvantage"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	serviceImg "github.com/imperatorofdwelling/Full-backend/internal/service/file"
)

type Service struct {
	Repo    interfaces.StaysAdvantageRepo
	StaySvc interfaces.StaysService
	AdvSvc  interfaces.AdvantageService
	FileSvc interfaces.FileService
}

func (s *Service) CreateStaysAdvantage(ctx context.Context, stayReq *models.StayAdvantageCreateReq) error {
	const op = "service.staysadvantage.CreateStaysAdvantage"

	stay, err := s.StaySvc.GetStayByID(ctx, stayReq.StayID)
	if err != nil {
		return err
	}

	if stay == nil {
		return fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	adv, err := s.AdvSvc.GetAdvantageByID(ctx, stayReq.AdvantageID)
	if err != nil {
		return err
	}

	if adv == nil {
		return fmt.Errorf("%s: %w", op, service.ErrAdvantageNotFound)
	}

	fWithPath, err := s.FileSvc.UploadImage([]byte(adv.Image), serviceImg.SvgImageType, "stays-advantage")
	if err != nil {
		return err
	}

	stayAdvantage := models.StayAdvantageEntity{
		StayID:      stay.ID,
		AdvantageID: adv.ID,
		Title:       adv.Title,
		Image:       fWithPath,
	}

	err = s.Repo.CreateStaysAdvantage(ctx, &stayAdvantage)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteStaysAdvantageByID(ctx context.Context, id uuid.UUID) error {
	const op = "service.staysadvantage.DeleteStaysAdvantageByID"

	isExists, err := s.Repo.CheckStaysAdvantageIfExists(ctx, id)
	if err != nil {
		return err
	}

	if !isExists {
		return fmt.Errorf("%s: %w", op, service.ErrAdvantageNotFound)
	}

	err = s.StaySvc.DeleteStayByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
