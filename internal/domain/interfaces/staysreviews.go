package interfaces

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreviews"
	"net/http"
)

//go:generate mockery --name StaysReviewsRepo
type StaysReviewsRepo interface {
	CreateStaysReview(context.Context, *staysreviews.StaysReviewEntity) error
	UpdateStaysReviewByID(context.Context, *staysreviews.StaysReviewEntity, uuid.UUID) error
	DeleteStaysReviewByID(context.Context, uuid.UUID) error
	FindOneStaysReviewByID(context.Context, uuid.UUID) (*staysreviews.StaysReview, error)
	FindAllStaysReviews(context.Context) ([]staysreviews.StaysReview, error)
	CheckIfExists(context.Context, uuid.UUID) (bool, error)
}

//go:generate mockery --name StaysReviewsService
type StaysReviewsService interface {
	CreateStaysReview(context.Context, *staysreviews.StaysReviewEntity) error
	UpdateStaysReview(context.Context, *staysreviews.StaysReviewEntity, uuid.UUID) (*staysreviews.StaysReview, error)
	DeleteStaysReview(context.Context, uuid.UUID) error
	FindOneStaysReview(context.Context, uuid.UUID) (*staysreviews.StaysReview, error)
	FindAllStaysReviews(context.Context) ([]staysreviews.StaysReview, error)
}

type StaysReviewsHandler interface {
	CreateStaysReview(w http.ResponseWriter, r *http.Request)
	UpdateStaysReview(w http.ResponseWriter, r *http.Request)
	DeleteStaysReview(w http.ResponseWriter, r *http.Request)
	FindOneStaysReview(w http.ResponseWriter, r *http.Request)
	FindAllStaysReviews(w http.ResponseWriter, r *http.Request)
}
