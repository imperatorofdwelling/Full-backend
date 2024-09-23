package interfaces

import (
	"context"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/location"
	"net/http"
)

//go:generate mockery --name LocationRepository
type LocationRepository interface {
	FindByNameMatch(ctx context.Context, match string) (*[]location.Location, error)
}

//go:generate mockery --name LocationService
type LocationService interface {
	FindByNameMatch(ctx context.Context, match string) (*[]location.Location, error)
}

type LocationHandler interface {
	FindByNameMatch(http.ResponseWriter, *http.Request)
}
