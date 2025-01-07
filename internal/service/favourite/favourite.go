package favourite

import (
	"context"
	"fmt"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	stays2 "github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
)

type Service struct {
	Repo interfaces.FavouriteRepo
}

func (s *Service) AddToFavourites(ctx context.Context, userId, stayID string) error {
	const op = "service.Favourite.AddFavourite"

	err := s.Repo.AddFavourite(ctx, userId, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) RemoveFromFavourites(ctx context.Context, userId, stayID string) error {
	const op = "service.Favourite.RemoveFavourite"

	err := s.Repo.RemoveFavourite(ctx, userId, stayID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) GetAllFavourites(ctx context.Context, userID string) ([]stays2.StayEntityFav, error) {
	const op = "service.Favourite.GetAllFavourites"

	favorites, err := s.Repo.GetAllFavourites(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return favorites, nil
}
