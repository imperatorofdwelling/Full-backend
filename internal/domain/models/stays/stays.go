package stays

import (
	"github.com/gofrs/uuid"
	"time"
)

var (
	Apartment StayType = "apartment"
	House     StayType = "house"
	Hotel     StayType = "hotel"
)

type (
	StayType string

	StayEntity struct {
		UserID              uuid.UUID `json:"user_id" validate:"required,uuid"`
		LocationID          uuid.UUID `json:"location_id" validate:"required,uuid"`
		Name                string    `json:"name" validate:"required"`
		Type                StayType  `json:"type" validate:"required"`
		NumberOfBedrooms    int       `json:"number_of_bedrooms" validate:"required"`
		NumberOfBeds        int       `json:"number_of_beds" validate:"required"`
		NumberOfBathrooms   int       `json:"number_of_bathrooms" validate:"required"`
		Guests              int       `json:"guests" validate:"required"`
		IsSmokingProhibited bool      `json:"is_smoking_prohibited,omitempty" default:"false"`
		Square              float32   `json:"square" validate:"required"`
		Street              string    `json:"street" validate:"required"`
		House               string    `json:"house" validate:"required"`
		Entrance            string    `json:"entrance,omitempty"`
		Floor               string    `json:"floor,omitempty"`
		Room                string    `json:"room,omitempty"`
		Price               float32   `json:"price" validate:"required"`
	}

	Stay struct {
		ID                  uuid.UUID `json:"id"`
		UserID              uuid.UUID `json:"user_id"`
		LocationID          uuid.UUID `json:"location_id"`
		Name                string    `json:"name"`
		Type                StayType  `json:"type"`
		NumberOfBedrooms    int       `json:"number_of_bedrooms"`
		NumberOfBeds        int       `json:"number_of_beds"`
		NumberOfBathrooms   int       `json:"number_of_bathrooms"`
		Guests              int       `json:"guests"`
		Rating              float32   `json:"rating"`
		IsSmokingProhibited bool      `json:"is_smoking_prohibited"`
		Square              float32   `json:"square"`
		Street              string    `json:"street"`
		House               string    `json:"house"`
		Entrance            string    `json:"entrance"`
		Floor               string    `json:"floor"`
		Room                string    `json:"room"`
		Price               float32   `json:"price"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
	}

	StayImagesEntity struct {
		Images []byte    `json:"images"`
		StayID uuid.UUID `json:"stay_id"`
	}

	StayImage struct {
		ID        uuid.UUID `json:"id"`
		StayID    uuid.UUID `json:"stay_id"`
		ImageName string    `json:"image_name"`
		IsMain    bool      `json:"is_main"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
