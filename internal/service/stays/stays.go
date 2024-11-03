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

func (s *Service) DeleteStayByID(ctx context.Context, id uuid.UUID) error {
	const op = "service.stays.DeleteStay"

	exists, err := s.Repo.CheckStayIfExistsByID(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	err = s.Repo.DeleteStayByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateStayByID(ctx context.Context, stay *stays.StayEntity, id uuid.UUID) (*stays.Stay, error) {
	const op = "service.stays.UpdateStayByID"

	foundLocation, err := s.LocSvc.GetByID(ctx, stay.LocationID)
	if err != nil {
		return &stays.Stay{}, err
	}

	if foundLocation == nil {
		return &stays.Stay{}, fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	err = s.Repo.UpdateStayByID(ctx, stay, id)
	if err != nil {
		return &stays.Stay{}, err
	}

	updatedStay, err := s.Repo.GetStayByID(ctx, id)
	if err != nil {
		return &stays.Stay{}, err
	}

	return updatedStay, nil
}

func (s *Service) GetStaysByUserID(ctx context.Context, userId uuid.UUID) ([]*stays.Stay, error) {
	const op = "service.stays.GetStaysByUserID"

	// TODO Check if user exists

	usrStays, err := s.Repo.GetStaysByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return usrStays, nil
}

func (s *Service) GetImagesByStayID(ctx context.Context, id uuid.UUID) ([]stays.StayImage, error) {
	const op = "service.stays.GetImagesByStayID"

	foundStay, err := s.GetStayByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if foundStay == nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	images, err := s.Repo.GetImagesByStayID(ctx, id)
	if err != nil {
		return nil, err
	}

	return images, nil

}

func (s *Service) GetMainImageByStayID(ctx context.Context, id uuid.UUID) (stays.StayImage, error) {
	const op = "service.stays.GetMainImageByStayID"

	foundStay, err := s.GetStayByID(ctx, id)
	if err != nil {
		return stays.StayImage{}, err
	}

	if foundStay == nil {
		return stays.StayImage{}, fmt.Errorf("%s: %w", op, service.ErrStayNotFound)
	}

	image, err := s.Repo.GetMainImageByStayID(ctx, id)
	if err != nil {
		return stays.StayImage{}, err
	}

	return image, nil
}
