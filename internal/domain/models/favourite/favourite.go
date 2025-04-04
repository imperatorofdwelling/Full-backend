package favourite

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	FavouriteEntity struct {
		UserID      uuid.UUID `json:"user_id"`
		StayID      uuid.UUID `json:"stay_id"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} // @name FavouriteEntity

	Favourite struct {
		UserID uuid.UUID `json:"user_id"`
		StayID uuid.UUID `json:"stay_id"`
		City   string    `json:"city"`
	} // @name Favourite
)
