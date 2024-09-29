package staysreviews

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/interfaces"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/staysreviews"
	"github.com/imperatorofdwelling/Full-backend/internal/service"
)

type Service struct {
	Repo interfaces.StaysReviewsRepo
}

func (s *Service) CreateStaysReview(ctx context.Context, stayReview *staysreviews.StaysReviewEntity) error {
	const op = "service.staysreviews.CreateStaysReview"

	err := s.Repo.CreateStaysReview(ctx, stayReview)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateStaysReview(ctx context.Context, stayReview *staysreviews.StaysReviewEntity, id uuid.UUID) error {
	const op = "service.staysreviews.UpdateStaysReview"

	isExists, err := s.Repo.CheckIfExists(ctx, id)
	if err != nil {
		return err
	}

	if !isExists {
		return fmt.Errorf("%s: %w", op, service.ErrStaysReviewNotFound)
	}

	err = s.Repo.UpdateStaysReviewByID(ctx, stayReview, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindOneStaysReview(ctx context.Context, id uuid.UUID) (*staysreviews.StaysReview, error) {
	const op = "service.staysreviews.FindOneStaysReview"

	foundStayReview, err := s.Repo.FindOneStaysReviewByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if foundStayReview.ID == uuid.Nil {
		return nil, fmt.Errorf("%s: %w", op, service.ErrStaysReviewNotFound)
	}

	return foundStayReview, nil
}

func (s *Service) DeleteStaysReview(ctx context.Context, id uuid.UUID) error {
	const op = "service.staysreviews.DeleteStaysReview"

	isExists, err := s.Repo.CheckIfExists(ctx, id)
	if err != nil {
		return err
	}

	if !isExists {
		return fmt.Errorf("%s: %w", op, service.ErrStaysReviewNotFound)
	}

	err = s.Repo.DeleteStaysReviewByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindAllStaysReviews(ctx context.Context) ([]staysreviews.StaysReview, error) {
	const op = "service.staysreviews.FindAllStaysReviews"

	reviews, err := s.Repo.FindAllStaysReviews(ctx)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
