package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"net/http"
)

type (
	UserRepository interface {
		CheckUserExists(ctx context.Context, email string) (bool, error)
		CreateUser(ctx context.Context, user *models.UserEntity) (uuid.UUID, error)
		FindUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	}

	UserService interface {
		CreateUser(ctx context.Context, user *models.UserEntity) (*models.User, error)
	}

	UserHandler interface {
		CreateUser(http.ResponseWriter, *http.Request)
	}
)
