package reservation

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	ReservationEntity struct {
		StayID    uuid.UUID `json:"stay_id"`
		UserID    uuid.UUID `json:"user_id"`
		Arrived   time.Time `json:"arrived"`
		Departure time.Time `json:"departure"`
	}

	Reservation struct {
		ID        uuid.UUID `json:"id"`
		StayID    uuid.UUID `json:"stay_id"`
		UserID    uuid.UUID `json:"user_id"`
		Arrived   time.Time `json:"arrived"`
		Departure time.Time `json:"departure"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
