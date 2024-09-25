package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	models "github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysadvantage"
	"net/http"
)

//go:generate mockery --name StaysAdvantageRepo
type StaysAdvantageRepo interface {
	CreateStaysAdvantage(ctx context.Context, stayAdv *models.StayAdvantageEntity) error
	DeleteStaysAdvantageByID(context.Context, uuid.UUID) error
	CheckStaysAdvantageIfExists(context.Context, uuid.UUID) (bool, error)
}

//go:generate mockery --name StaysAdvantageService
type StaysAdvantageService interface {
	CreateStaysAdvantage(ctx context.Context, req *models.StayAdvantageCreateReq) error
	DeleteStaysAdvantageByID(ctx context.Context, id uuid.UUID) error
}

type StaysAdvantageHandler interface {
	CreateStaysAdvantage(http.ResponseWriter, *http.Request)
	DeleteStaysAdvantageByID(http.ResponseWriter, *http.Request)
}
