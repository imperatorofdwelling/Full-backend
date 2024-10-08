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
	GetStaysByUserID(context.Context, uuid.UUID) ([]*stays.Stay, error)
	DeleteStayByID(context.Context, uuid.UUID) error
	UpdateStayByID(context.Context, *stays.StayEntity, uuid.UUID) error
	CheckStayIfExistsByID(context.Context, uuid.UUID) (bool, error)
}

//go:generate mockery --name StaysService
type StaysService interface {
	CreateStay(context.Context, *stays.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*stays.Stay, error)
	GetStays(context.Context) ([]*stays.Stay, error)
	GetStaysByUserID(context.Context, uuid.UUID) ([]*stays.Stay, error)
	DeleteStayByID(context.Context, uuid.UUID) error
	UpdateStayByID(context.Context, *stays.StayEntity, uuid.UUID) (*stays.Stay, error)
}

type StaysHandler interface {
	CreateStay(http.ResponseWriter, *http.Request)
	GetStayByID(http.ResponseWriter, *http.Request)
	GetStays(http.ResponseWriter, *http.Request)
	GetStaysByUserID(http.ResponseWriter, *http.Request)
	DeleteStayByID(http.ResponseWriter, *http.Request)
	UpdateStayByID(http.ResponseWriter, *http.Request)
}
