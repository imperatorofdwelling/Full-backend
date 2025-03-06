package stays

import (
	"github.com/gofrs/uuid"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays/amenity"
	"github.com/imperatorofdwelling/Full-backend/internal/domain/models/stays/sort"
	"time"
)

var (
	Apartment StayType = "apartment"
	House     StayType = "house"
	Hotel     StayType = "hotel"
)

type (
	StayType   string
	StayEntity struct {
		UserID              uuid.UUID       `json:"user_id" validate:"required,uuid"`
		LocationID          uuid.UUID       `json:"location_id" validate:"required,uuid"`
		Name                string          `json:"name" validate:"required"`
		Type                StayType        `json:"type" validate:"required"`
		NumberOfBedrooms    int             `json:"number_of_bedrooms" validate:"required"`
		NumberOfBeds        int             `json:"number_of_beds" validate:"required"`
		NumberOfBathrooms   int             `json:"number_of_bathrooms" validate:"required"`
		Guests              int             `json:"guests" validate:"required"`
		Amenities           map[string]bool `json:"amenities" validate:"required"`
		IsSmokingProhibited bool            `json:"is_smoking_prohibited,omitempty" default:"false"`
		Square              float32         `json:"square" validate:"required"`
		Street              string          `json:"street" validate:"required"`
		House               string          `json:"house" validate:"required"`
		Entrance            string          `json:"entrance,omitempty"`
		Floor               string          `json:"floor,omitempty"`
		Room                string          `json:"room,omitempty"`
		Price               float32         `json:"price" validate:"required"`
	} // @name StayEntity

	StayEntityFav struct {
		ID                  uuid.UUID       `json:"id"`
		UserID              uuid.UUID       `json:"user_id" validate:"required,uuid"`
		LocationID          uuid.UUID       `json:"location_id" validate:"required,uuid"`
		Name                string          `json:"name" validate:"required"`
		Type                StayType        `json:"type" validate:"required"`
		NumberOfBedrooms    int             `json:"number_of_bedrooms" validate:"required"`
		NumberOfBeds        int             `json:"number_of_beds" validate:"required"`
		NumberOfBathrooms   int             `json:"number_of_bathrooms" validate:"required"`
		Guests              int             `json:"guests" validate:"required"`
		Amenities           map[string]bool `json:"amenities" validate:"required"`
		IsSmokingProhibited bool            `json:"is_smoking_prohibited,omitempty" default:"false"`
		Square              float32         `json:"square" validate:"required"`
		Street              string          `json:"street" validate:"required"`
		House               string          `json:"house" validate:"required"`
		Entrance            string          `json:"entrance,omitempty"`
		Floor               string          `json:"floor,omitempty"`
		Room                string          `json:"room,omitempty"`
		Price               float32         `json:"price" validate:"required"`
		City                string          `json:"city" validate:"required"`
	} // @name StayEntityFav

	Stay struct {
		ID                  uuid.UUID       `json:"id"`
		UserID              uuid.UUID       `json:"user_id"`
		LocationID          uuid.UUID       `json:"location_id"`
		Name                string          `json:"name"`
		Type                StayType        `json:"type"`
		NumberOfBedrooms    int             `json:"number_of_bedrooms"`
		NumberOfBeds        int             `json:"number_of_beds"`
		NumberOfBathrooms   int             `json:"number_of_bathrooms"`
		Guests              int             `json:"guests"`
		Rating              float64         `json:"rating"`
		Amenities           map[string]bool `json:"amenities" validate:"required"`
		IsSmokingProhibited bool            `json:"is_smoking_prohibited"`
		Square              float32         `json:"square"`
		Street              string          `json:"street"`
		House               string          `json:"house"`
		Entrance            string          `json:"entrance"`
		Floor               string          `json:"floor"`
		Room                string          `json:"room"`
		Price               float32         `json:"price"`
		CreatedAt           time.Time       `json:"created_at"`
		UpdatedAt           time.Time       `json:"updated_at"`
	} // @name Stay

	StayImagesEntity struct {
		Images []byte    `json:"images"`
		StayID uuid.UUID `json:"stay_id"`
	} // @name StayImagesEntity

	StayImage struct {
		ID        uuid.UUID `json:"id"`
		StayID    uuid.UUID `json:"stay_id"`
		ImageName string    `json:"image_name"`
		IsMain    bool      `json:"is_main"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} // @name StayImage

	StayResponse struct {
		Stay
		Images []StayImage `json:"images"`
	} // @name StayResponse

	Filtration struct {
		Location         string                   `json:"locationName" validate:"required"`
		SortBy           sort.Sort                `json:"sort_by" validate:"omitempty"`
		PriceMin         float32                  `json:"price_min" validate:"omitempty"`
		PriceMax         float32                  `json:"price_max" validate:"omitempty"`
		NumberOfBedrooms []int32                  `json:"number_of_bedrooms" validate:"omitempty"`
		Amenities        map[amenity.Amenity]bool `json:"amenities" validate:"omitempty"`
		Rating           []float64                `json:"rating" validate:"omitempty"`
	} // @name Filtration
)

func (f *Filtration) SetDefaults() {
	if f.PriceMin == 0 {
		f.PriceMin = -1
	}
	if f.PriceMax == 0 {
		f.PriceMax = -1
	}
}
