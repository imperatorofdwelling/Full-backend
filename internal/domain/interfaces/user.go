package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
)

type (
	UserRepository interface {
		CheckUserExists(ctx context.Context, email string) (bool, error)
		FindUserByID(ctx context.Context, id uuid.UUID) (user.User, error)
		UpdateUserByID(ctx context.Context, id uuid.UUID, user user.User) (user.User, error)
		DeleteUserByID(ctx context.Context, id uuid.UUID) error
	}

	UserService interface {
		UpdateUserByID(ctx context.Context, id uuid.UUID, user user.User) (user.User, error)
		DeleteUserByID(ctx context.Context, id uuid.UUID) error
	}

	UserHandler interface {
		UpdateUserByID(ctx context.Context, id uuid.UUID, user user.User) (user.User, error)
		DeleteUserByID(ctx context.Context, id uuid.UUID) error
	}
)
