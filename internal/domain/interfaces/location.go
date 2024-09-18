package interfaces

import (
	"context"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"net/http"
)

//go:generate mockery --name LocationRepository
type LocationRepository interface {
	FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error)
}

//go:generate mockery --name LocationService
type LocationService interface {
	FindByNameMatch(ctx context.Context, match string) (*[]models.Location, error)
}

type LocationHandler interface {
	FindByNameMatch(http.ResponseWriter, *http.Request)
}
