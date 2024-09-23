package location

import (
	"context"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	"strings"
)

type Service struct {
	Repo interfaces.LocationRepository
}

func (s *Service) FindByNameMatch(ctx context.Context, match string) (*[]location.Location, error) {
	m := strings.TrimSpace(match)

	locations, err := s.Repo.FindByNameMatch(ctx, m)
	if err != nil {
		return nil, err
	}

	return locations, err
}
