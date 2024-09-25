package location

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
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
	if err != nil {
		return nil, err
	}

	return loc, nil
}
