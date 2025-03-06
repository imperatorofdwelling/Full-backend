package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays"
	"mime/multipart"
	"net/http"
)

//go:generate mockery --name StaysRepo
type StaysRepo interface {
	CreateStay(context.Context, *stays.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*stays.Stay, error)
	GetStays(context.Context) ([]stays.StayResponse, error)
	GetStaysByUserID(context.Context, uuid.UUID) ([]*stays.Stay, error)
	DeleteStayByID(context.Context, uuid.UUID) error
	UpdateStayByID(context.Context, *stays.StayEntity, uuid.UUID) error
	CheckStayIfExistsByID(context.Context, uuid.UUID) (bool, error)
	GetImagesByStayID(context.Context, uuid.UUID) ([]stays.StayImage, error)
	GetMainImageByStayID(context.Context, uuid.UUID) (stays.StayImage, error)
	GetStayImageByID(context.Context, uuid.UUID) (stays.StayImage, error)
	CreateStayImage(ctx context.Context, fileName string, isMain bool, stayID uuid.UUID) error
	DeleteStayImage(context.Context, uuid.UUID) error
	GetStaysByLocationID(context.Context, uuid.UUID) (*[]stays.Stay, error)
	Filtration(ctx context.Context, search stays.Filtration, locationIDS []uuid.UUID) ([]stays.Stay, error)
}

//go:generate mockery --name StaysService
type StaysService interface {
	CreateStay(context.Context, *stays.StayEntity) error
	GetStayByID(context.Context, uuid.UUID) (*stays.Stay, error)
	GetStays(context.Context) ([]stays.StayResponse, error)
	GetStaysByUserID(context.Context, uuid.UUID) ([]*stays.Stay, error)
	DeleteStayByID(context.Context, uuid.UUID) error
	UpdateStayByID(context.Context, *stays.StayEntity, uuid.UUID) (*stays.Stay, error)
	GetImagesByStayID(context.Context, uuid.UUID) ([]stays.StayImage, error)
	GetMainImageByStayID(context.Context, uuid.UUID) (stays.StayImage, error)
	CreateImages(context.Context, []*multipart.FileHeader, uuid.UUID) error
	CreateMainImage(context.Context, *multipart.FileHeader, uuid.UUID) error
	DeleteStayImage(context.Context, uuid.UUID) error
	GetStaysByLocationID(context.Context, uuid.UUID) (*[]stays.Stay, error)
	Filtration(ctx context.Context, search stays.Filtration) ([]stays.Stay, error)
}

type StaysHandler interface {
	CreateStay(http.ResponseWriter, *http.Request)
	GetStayByID(http.ResponseWriter, *http.Request)
	GetStays(http.ResponseWriter, *http.Request)
	GetStaysByUserID(http.ResponseWriter, *http.Request)
	DeleteStayByID(http.ResponseWriter, *http.Request)
	UpdateStayByID(http.ResponseWriter, *http.Request)
	GetStayImagesByStayID(http.ResponseWriter, *http.Request)
	GetMainImageByStayID(http.ResponseWriter, *http.Request)
	CreateImages(http.ResponseWriter, *http.Request)
	CreateMainImage(http.ResponseWriter, *http.Request)
	DeleteStayImage(http.ResponseWriter, *http.Request)
	GetStaysByLocationID(http.ResponseWriter, *http.Request)
	Filtration(http.ResponseWriter, *http.Request)
}
