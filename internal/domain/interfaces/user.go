package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"net/http"
)

type (
	UserRepository interface {
		CheckUserExists(ctx context.Context, email string) (bool, error)
		FindUserByID(ctx context.Context, id uuid.UUID) (user.User, error)
		UpdateUserByID(ctx context.Context, id uuid.UUID, user user.User) error
		DeleteUserByID(ctx context.Context, id uuid.UUID) error
	}

	UserService interface {
		GetUserByID(ctx context.Context, idStr string) (user.User, error)
		UpdateUserByID(ctx context.Context, idStr string, user user.User) (user.User, error)
		DeleteUserByID(ctx context.Context, idStr string) error
	}

	UserHandler interface {
		GetUserByID(w http.ResponseWriter, r *http.Request)
		UpdateUserByID(w http.ResponseWriter, r *http.Request)
		DeleteUserByID(w http.ResponseWriter, r *http.Request)
	}
)
