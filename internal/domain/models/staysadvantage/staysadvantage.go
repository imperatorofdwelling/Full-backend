package staysadvantage

import (
	"github.com/gofrs/uuid"
	"time"
)

type (
	StayAdvantageCreateReq struct {
		StayID      uuid.UUID `json:"stay_id"`
		AdvantageID uuid.UUID `json:"advantage_id"`
	} // @name StayAdvantageCreateReq

	StayAdvantageEntity struct {
		StayID      uuid.UUID `json:"stay_id"`
		AdvantageID uuid.UUID `json:"advantage_id"`
		Title       string    `json:"title"`
		Image       string    `json:"image"`
	} // @name StayAdvantageEntity

	StayAdvantage struct {
		ID          uuid.UUID `json:"id"`
		StayID      uuid.UUID `json:"stay_id"`
		AdvantageID uuid.UUID `json:"advantage_id"`
		Title       string    `json:"title"`
		Image       string    `json:"image"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} // @name StayAdvantage
)
