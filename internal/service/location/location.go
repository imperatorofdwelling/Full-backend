package service

import (
	"context"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models"
	"strings"
)

type LocationService struct {
	Repo interfaces.LocationRepository
}

func (s *LocationService) FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error) {
	m := strings.TrimSpace(match)

	locations, err := s.Repo.FindByNameMatch(ctx, m)
	if err != nil {
		return nil, err
	}

	return locations, err
}
