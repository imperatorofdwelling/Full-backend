package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/advantage"
	"net/http"
)

//go:generate mockery --name AdvantageRepo
type AdvantageRepo interface {
	CreateAdvantage(ctx context.Context, advTitle string, imgPath string) error
	CheckAdvantageIfExists(ctx context.Context, advName string) (bool, error)
	RemoveAdvantage(ctx context.Context, id uuid.UUID) error
	FindAdvantageByID(ctx context.Context, id uuid.UUID) (*advantage.Advantage, error)
	GetAllAdvantages(context.Context) ([]advantage.Advantage, error)
	UpdateAdvantageByID(ctx context.Context, id uuid.UUID, adv *advantage.Advantage) error
}

//go:generate mockery --name AdvantageService
type AdvantageService interface {
	CreateAdvantage(context.Context, *advantage.AdvantageEntity) error
	RemoveAdvantage(context.Context, uuid.UUID) error
	GetAllAdvantages(context.Context) ([]advantage.Advantage, error)
	UpdateAdvantageByID(context.Context, uuid.UUID, *advantage.AdvantageEntity) (advantage.Advantage, error)
	GetAdvantageByID(context.Context, uuid.UUID) (*advantage.Advantage, error)
}

type AdvantageHandler interface {
	CreateAdvantage(http.ResponseWriter, *http.Request)
	RemoveAdvantage(http.ResponseWriter, *http.Request)
	GetAllAdvantages(http.ResponseWriter, *http.Request)
	UpdateAdvantage(http.ResponseWriter, *http.Request)
}
