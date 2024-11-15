package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	"net/http"
)

//go:generate mockery --name LocationRepo
type LocationRepo interface {
	FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	GetAll(ctx context.Context) (*[]models.Location, error)
	GetOneByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	UpdateByID(ctx context.Context, id uuid.UUID, location models.LocationEntity) error
}

//go:generate mockery --name LocationService
type LocationService interface {
	FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	GetAll(ctx context.Context) (*[]models.Location, error)
	GetOneByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
	UpdateByID(ctx context.Context, id uuid.UUID, location models.LocationEntity) error
}

type LocationHandler interface {
	FindByNameMatch(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	GetOneByID(http.ResponseWriter, *http.Request)
	DeleteByID(http.ResponseWriter, *http.Request)
	UpdateByID(http.ResponseWriter, *http.Request)
}
