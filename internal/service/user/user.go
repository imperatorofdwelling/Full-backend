package service

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Website-backend/internal/service"
)

type UserService struct {
	Repo interfaces.UserRepository
}

func (s *UserService) CreateUser(ctx context.Context, user *user.Registration) (*user.User, error) {
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

func (s *UserService) Login(ctx context.Context, user *user.Login) (*user.User, error) {
	const op = "service.user.Login"
	userExists, err := s.Repo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if !userExists {
		return nil, fmt.Errorf("%s: %w", op, service.ErrUserNotFound)
	}

	id, err := s.Repo.Login(ctx, user)
}
