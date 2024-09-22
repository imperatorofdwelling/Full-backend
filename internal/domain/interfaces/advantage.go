package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Website-backend/internal/domain/models"
	"net/http"
)

//go:generate mockery --name AdvantageRepo
type AdvantageRepo interface {
	CreateAdvantage(ctx context.Context, advTitle string, imgPath string) error
	CheckAdvantageIfExists(ctx context.Context, advName string) (bool, error)
	RemoveAdvantage(ctx context.Context, id uuid.UUID) error
	FindAdvantageByID(ctx context.Context, id uuid.UUID) (*models.Advantage, error)
	GetAllAdvantages(context.Context) ([]models.Advantage, error)
	UpdateAdvantageByID(ctx context.Context, id uuid.UUID, adv *models.Advantage) error
}

//go:generate mockery --name AdvantageService
type AdvantageService interface {
	CreateAdvantage(context.Context, *models.AdvantageEntity) error
	RemoveAdvantage(context.Context, uuid.UUID) error
	GetAllAdvantages(context.Context) ([]models.Advantage, error)
	UpdateAdvantageByID(context.Context, uuid.UUID, *models.AdvantageEntity) (models.Advantage, error)
}

type AdvantageHandler interface {
	CreateAdvantage(http.ResponseWriter, *http.Request)
	RemoveAdvantage(http.ResponseWriter, *http.Request)
	GetAllAdvantages(http.ResponseWriter, *http.Request)
	UpdateAdvantage(http.ResponseWriter, *http.Request)
}
