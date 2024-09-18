package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
)

type (
	AuthRepository interface {
		Registration(ctx context.Context, user *auth.Registration) (uuid.UUID, error)
		Login(ctx context.Context, user *auth.Login) (uuid.UUID, error)
	}

	AuthService interface {
		Registration(ctx context.Context, user *auth.Registration) (uuid.UUID, error)
		Login(ctx context.Context, user *auth.Login) (uuid.UUID, error)
	}

	AuthHandler interface {
	}
)
