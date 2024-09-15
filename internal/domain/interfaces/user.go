package interfaces

import (
	"context"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"net/http"
)

type (
	UserRepository interface {
		FetchByUsername(ctx context.Context, username string) (*models.UserEntity, error)
	}

	UserService interface {
		FetchByUsername(ctx context.Context, username string) (*models.User, error)
	}

	UserHandler interface {
		CreateUser(http.ResponseWriter, *http.Request)
		FetchByUsername() http.HandlerFunc
	}
)
