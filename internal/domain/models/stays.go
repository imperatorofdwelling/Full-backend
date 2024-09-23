package models

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
		UserID              uuid.UUID `json:"user_id"`
		LocationID          uuid.UUID `json:"location_id"`
		Name                string    `json:"name"`
		ImageMain           string    `json:"image_main"`
		Images              []string  `json:"images"`
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
	}

	Stay struct {
		ID                  uuid.UUID `json:"id"`
		UserID              uuid.UUID `json:"user_id"`
		LocationID          uuid.UUID `json:"location_id"`
		Name                string    `json:"name"`
		ImageMain           string    `json:"image_main"`
		Images              []string  `json:"images"`
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
)
