package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models/user"
	"net/http"
)

type (
	UserRepository interface {
		CheckUserExists(ctx context.Context, email string) (bool, error)
		CreateUser(ctx context.Context, user *user.Entity) (uuid.UUID, error)
		FindUserByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	}

	UserService interface {
		CreateUser(ctx context.Context, user *user.Entity) (*user.User, error)
	}

	UserHandler interface {
		CreateUser(http.ResponseWriter, *http.Request)
	}
)
