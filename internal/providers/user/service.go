package user

import (
	"context"
	"github.com/imperatorofdwelling/Website-backend/internal/domain"
)

type service struct {
	repo domain.UserRepository
}

func (s *service) FetchByUsername(ctx context.Context, username string) (*domain.User, error) {
	panic("implement me")
}
