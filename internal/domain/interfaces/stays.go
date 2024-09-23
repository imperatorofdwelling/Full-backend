package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"net/http"
)

//go:generate mockery --name StaysRepo
type StaysRepo interface {
	CreateStay(context.Context, *models.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*models.Stay, error)
	GetStays(context.Context) ([]*models.Stay, error)
}

//go:generate mockery --name StaysService
type StaysService interface {
	CreateStay(context.Context, *models.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*models.Stay, error)
	GetStays(context.Context) ([]*models.Stay, error)
}

type StaysHandler interface {
	CreateStay(http.ResponseWriter, *http.Request)
	GetStayByID(http.ResponseWriter, *http.Request)
	GetStays(http.ResponseWriter, *http.Request)
}
