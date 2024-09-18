package auth

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
)

type Service struct {
	AuthRepo interfaces.AuthRepository
	UserRepo interfaces.UserRepository
}

func (s *Service) Registration(ctx context.Context, user *auth.Registration) (uuid.UUID, error) {
	const op = "service.user.Registration"

	userExists, err := s.UserRepo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if userExists {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrUserAlreadyExists)
	}

	id, err := s.AuthRepo.Registration(ctx, user)
	if err != nil {
		return uuid.Nil, err
	}

	userFound, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		return uuid.Nil, err
	}

	if userFound.ID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrUserNotFound)
	}

	return userFound.ID, nil
}

func (s *Service) Login(ctx context.Context, user *auth.Login) (uuid.UUID, error) {
	const op = "service.user.Login"
	userExists, err := s.UserRepo.CheckUserExists(ctx, user.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if !userExists {
		return uuid.Nil, fmt.Errorf("%s: %w", op, service.ErrUserNotFound)
	}

	id, err := s.AuthRepo.Login(ctx, user)
	if err != nil {
		return id, err
	}

	return id, err
}
