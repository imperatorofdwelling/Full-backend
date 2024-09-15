package user

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"github.com/imperatorofdwelling/Website-backend/internal/service"
)

type Service struct {
	Repo interfaces.UserRepository
}

func (s *Service) CreateUser(ctx context.Context, user *models.UserEntity) (*models.User, error) {
	const op = "service.user.CreateUser"

	userExists, err := s.Repo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, fmt.Errorf("%s: %w", op, service.ErrUserAlreadyExists)
	}

	id, err := s.Repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	userFound, err := s.Repo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if userFound.ID == uuid.Nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrUserNotFound)
	}

	return userFound, nil
}
