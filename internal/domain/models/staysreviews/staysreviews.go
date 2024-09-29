package staysreviews

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	StaysReviewEntity struct {
		StayID      string  `json:"stay_id"`
		UserID      string  `json:"user_id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Rating      float32 `json:"rating"`
	}

	StaysReview struct {
		ID          uuid.UUID `json:"id"`
		StayID      string    `json:"stay_id"`
		UserID      string    `json:"user_id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Rating      float32   `json:"rating"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
