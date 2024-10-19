package user

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
)

type Service struct {
	Repo interfaces.UserRepository
}

func (s *Service) GetUserByID(ctx context.Context, idStr string) (model.User, error) {
	const op = "service.user.GetUserByID"

	id, err := uuid.FromString(idStr)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	result, err := s.Repo.FindUserByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return result, nil
}

func (s *Service) DeleteUserByID(ctx context.Context, idStr string) error {
	const op = "service.user.DeleteUserByID"

	id, err := uuid.FromString(idStr)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.Repo.DeleteUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) UpdateUserByID(ctx context.Context, idStr string, user model.User) (model.User, error) {
	const op = "service.user.UpdateUserByID"
	id, err := uuid.FromString(idStr)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	oldUser, err := s.Repo.FindUserByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	user, err = s.compareUsers(ctx, oldUser, user)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	err = s.Repo.UpdateUserByID(ctx, id, user)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}
	updatedUser, err := s.Repo.FindUserByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	return updatedUser, nil
}

func (s *Service) stringToUUID(id string) (uuid.UUID, error) {
	u, err := uuid.FromString(id)
	if err != nil {
		return uuid.Nil, err
	}
	return u, nil
}

func (s *Service) compareUsers(ctx context.Context, oldUser model.User, newUser model.User) (model.User, error) {
	if newUser.ID == uuid.Nil {
		newUser.ID = oldUser.ID
	}
	if newUser.Name == "" {
		newUser.Name = oldUser.Name
	}
	if newUser.Email == "" {
		newUser.Email = oldUser.Email
	} else {
		if existEmail, _ := s.Repo.CheckUserExists(ctx, newUser.Email); existEmail {
			return model.User{}, service.ErrEmailAlreadyExists
		}
	}
	if newUser.Phone == "" {
		newUser.Phone = oldUser.Phone
	}
	if len(newUser.Avatar) == 0 {
		newUser.Avatar = oldUser.Avatar
	}
	if newUser.BirthDate.Valid {
		newUser.BirthDate = oldUser.BirthDate
	}
	if newUser.National == "" {
		newUser.National = oldUser.National
	}
	if newUser.Gender == "" {
		newUser.Gender = oldUser.Gender
	}
	if newUser.Country == "" {
		newUser.Country = oldUser.Country
	}
	if newUser.City == "" {
		newUser.City = oldUser.City
	}
	if newUser.CreatedAt.IsZero() {
		newUser.CreatedAt = oldUser.CreatedAt
	}
	if newUser.UpdatedAt.IsZero() {
		newUser.UpdatedAt = oldUser.UpdatedAt
	}
	return newUser, nil
}
