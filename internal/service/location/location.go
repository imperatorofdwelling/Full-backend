package location

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"strings"
)

type Service struct {
	Repo interfaces.LocationRepo
}

func (s *Service) FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error) {
	m := strings.TrimSpace(match)

	locations, err := s.Repo.FindByNameMatch(ctx, m)
	if err != nil {
		return nil, err
	}

	return locations, err
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*models.Location, error) {
	const op = "service.location.GetByID"

	loc, err := s.Repo.GetByID(ctx, id)
	if loc.ID == uuid.Nil {
		return &models.Location{}, fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	if err != nil {
		return nil, err
	}

	return loc, nil
}

func (s *Service) GetAll(ctx context.Context) (*[]models.Location, error) {
	const op = "service.location.GetAll"

	locs, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return locs, nil
}

func (s *Service) DeleteByID(ctx context.Context, id uuid.UUID) error {
	const op = "service.location.DeleteByID"

	loc, err := s.Repo.GetByID(ctx, id)
	if loc.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	if err != nil {
		return err
	}

	err = s.Repo.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateByID(ctx context.Context, id uuid.UUID, location models.LocationEntity) error {
	const op = "service.location.UpdateByID"

	loc, err := s.Repo.GetByID(ctx, id)
	if loc.ID == uuid.Nil {
		return fmt.Errorf("%s: %w", op, service.ErrLocationNotFound)
	}

	if err != nil {
		return err
	}

	err = s.Repo.UpdateByID(ctx, id, location)
	if err != nil {
		return err
	}

	return nil
}
