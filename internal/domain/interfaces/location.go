package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"net/http"
)

//go:generate mockery --name LocationRepo
type LocationRepo interface {
	FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
}

//go:generate mockery --name LocationService
type LocationService interface {
	FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
}

type LocationHandler interface {
	FindByNameMatch(http.ResponseWriter, *http.Request)
}
