package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"net/http"
)

//go:generate mockery --name StaysRepo
type StaysRepo interface {
	CreateStay(context.Context, *stays.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*stays.Stay, error)
	GetStays(context.Context) ([]*stays.Stay, error)
}

//go:generate mockery --name StaysService
type StaysService interface {
	CreateStay(context.Context, *stays.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*stays.Stay, error)
	GetStays(context.Context) ([]*stays.Stay, error)
}

type StaysHandler interface {
	CreateStay(http.ResponseWriter, *http.Request)
	GetStayByID(http.ResponseWriter, *http.Request)
	GetStays(http.ResponseWriter, *http.Request)
}
