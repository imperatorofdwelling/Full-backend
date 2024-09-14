package service

import (
	"context"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
)

type UserService struct {
	Repo interfaces.UserRepository
}

func (s *UserService) FetchByUsername(ctx context.Context, username string) (*models.User, error) {
	panic("implement me")
}
