package stays

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	staysInterface "github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
)

type Service struct {
	Repo   staysInterface.StaysRepo
	LocSvc staysInterface.LocationService
}

func (s *Service) CreateStay(ctx context.Context, stay *stays.StayEntity) error {
	const op = "service.stays.CreateStay"

	foundLocation, err := s.LocSvc.GetByID(ctx, stay.LocationID)
	if err != nil {
		return err
	}
	if foundLocation.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	//TODO check user if exists with user svc

	err = s.Repo.CreateStay(ctx, stay)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetStayByID(ctx context.Context, id uuid.UUID) (*stays.Stay, error) {
	const op = "service.stays.GetStayByID"

	stay, err := s.Repo.GetStayByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if stay == nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	return stay, nil
}

func (s *Service) GetStays(ctx context.Context) ([]*stays.Stay, error) {
	const op = "service.stays.GetStays"

	stays, err := s.Repo.GetStays(ctx)
	if err != nil {
		return nil, err
	}

	return stays, nil
}
