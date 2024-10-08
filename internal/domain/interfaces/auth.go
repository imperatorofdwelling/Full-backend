package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/auth"
	"net/http"
)

type (
	AuthRepository interface {
		Register(ctx context.Context, user auth.Registration) (uuid.UUID, error)
		Login(ctx context.Context, user auth.Login) (uuid.UUID, error)
	}

	AuthService interface {
		Register(ctx context.Context, user auth.Registration) (uuid.UUID, error)
		Login(ctx context.Context, user auth.Login) (uuid.UUID, error)
	}

	AuthHandler interface {
		Registration(w http.ResponseWriter, r *http.Request)
		LoginUser(w http.ResponseWriter, r *http.Request)
	}
)
