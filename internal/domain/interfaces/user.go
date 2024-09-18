package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
)

type (
	UserRepository interface {
		CheckUserExists(ctx context.Context, email string) (bool, error)
		FindUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	}

	UserService interface {
	}

	UserHandler interface {
	}
)
