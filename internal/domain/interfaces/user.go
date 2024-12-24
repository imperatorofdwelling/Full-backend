package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/user"
	"net/http"
)

//go:generate mockery --name UserRepository
type (
	UserRepository interface {
		CheckUserExists(ctx context.Context, email string) (bool, error)
		GetUserIDByEmail(ctx context.Context, email string) (string, error)
		FindUserByID(ctx context.Context, id uuid.UUID) (user.User, error)
		UpdateUserByID(ctx context.Context, id uuid.UUID, user user.User) error
		DeleteUserByID(ctx context.Context, id uuid.UUID) error
	}
)

//go:generate mockery --name UserService
type (
	UserService interface {
		GetUserByID(ctx context.Context, idStr string) (user.User, error)
		UpdateUserByID(ctx context.Context, idStr string, user user.User) (user.User, error)
		DeleteUserByID(ctx context.Context, idStr string) error
	}
)

type (
	UserHandler interface {
		GetUserByID(w http.ResponseWriter, r *http.Request)
		UpdateUserByID(w http.ResponseWriter, r *http.Request)
		DeleteUserByID(w http.ResponseWriter, r *http.Request)
	}
)
