package user

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/newPassword"
	model "github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
	"github.com/imperatorofdwelling/Full-backend/pkg/sendMail"
)

type Service struct {
	UserRepo         interfaces.UserRepository
	ConfirmEmailRepo interfaces.ConfirmEmailRepository
}

func (s *Service) GetUserByID(ctx context.Context, idStr string) (model.User, error) {
	const op = "service.user.GetUserByID"

	id, err := uuid.FromString(idStr)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	result, err := s.UserRepo.FindUserByID(ctx, id)
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

	err = s.UserRepo.DeleteUserByID(ctx, id)
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

	oldUser, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	user, err = s.compareUsers(ctx, oldUser, user)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}

	err = s.UserRepo.UpdateUserByID(ctx, id, user)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, err)
	}
	updatedUser, err := s.UserRepo.FindUserByID(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("%s: %w", op, service.ErrNotFound)
	}

	return updatedUser, nil
}

func (s *Service) UpdateUserPasswordByEmail(ctx context.Context, newPass newPassword.NewPassword) error {
	const op = "service.user.UpdateUserPasswordByEmail"

	if err := s.CheckUserPassword(ctx, newPass); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	userID, err := s.UserRepo.GetUserIDByEmail(ctx, newPass.Email)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.UserRepo.UpdateUserPasswordByID(ctx, uuid.FromStringOrNil(userID), newPass.Password)
	if err != nil {
		return fmt.Errorf("%s: failed to update password: %w", op, err)
	}

	err = s.ConfirmEmailRepo.UpdatePasswordOTPFalse(ctx, newPass.Email)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Service) CheckUserPassword(ctx context.Context, newPass newPassword.NewPassword) error {
	const op = "service.user.CheckUserPassword"

	isVerified, err := s.ConfirmEmailRepo.CheckPasswordOTPVerified(ctx, newPass.Email)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if !isVerified {
		return fmt.Errorf("%s: attempt to change password without verifying email", op)
	}

	tooLong, err := s.ConfirmEmailRepo.CheckPasswordOTPVerifiedForTooLong(ctx, newPass.Email)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if tooLong {
		err = s.ConfirmEmailRepo.ResetPasswordOTP(ctx, newPass.Email)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		newOTP, err := s.ConfirmEmailRepo.GetPasswordOTP(ctx, newPass.Email)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		err = sendMail.SimpleEmailSend(newPass.Email, newOTP, "Password Reset")
		if err != nil {
			return fmt.Errorf("%s : failed to send email to user: %w", op, err)
		}

		return fmt.Errorf("%s: previous code expired, we sent you a new one, please approve it again", op)
	}

	return nil
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
		if existEmail, _ := s.UserRepo.CheckUserExists(ctx, newUser.Email); existEmail {
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
