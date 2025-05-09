package stays

import (
	"fmt"
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
		UserID             uuid.UUID                `json:"user_id" validate:"required,uuid"`
		LocationID         uuid.UUID                `json:"location_id" validate:"required,uuid"`
		Name               string                   `json:"name" validate:"required"`
		Type               StayType                 `json:"type" validate:"required"`
		Guests             int                      `json:"guests" validate:"required"`
		Amenities          map[amenity.Amenity]bool `json:"amenities" validate:"required"`
		House              string                   `json:"house" validate:"required"`
		Entrance           string                   `json:"entrance,omitempty"`
		Address            string                   `json:"address" validate:"required"`
		RoomsCount         string                   `json:"rooms_count" validate:"required"`
		BedsCount          string                   `json:"beds_count" validate:"required"`
		Price              string                   `json:"price" validate:"required"`
		Period             string                   `json:"period" validate:"required"`
		OwnersRules        string                   `json:"owners_rules" validate:"required"`
		CancellationPolicy string                   `json:"cancellation_policy" validate:"required"`
		DescribeProperty   string                   `json:"describe_property" validate:"required"`
		CreatedAt          time.Time                `json:"created_at"`
		UpdatedAt          time.Time                `json:"updated_at"`
	} // @name StayEntity

	Statistics struct {
		StayTotal    int `json:"total_stays"`
		StayFree     int `json:"stay_free"`
		StayOccupied int `json:"stay_occupied"`
	} // @name StayStatistics

	StayEntityFav struct {
		ID                 uuid.UUID                `json:"id"`
		UserID             uuid.UUID                `json:"user_id" validate:"required,uuid"`
		LocationID         uuid.UUID                `json:"location_id" validate:"required,uuid"`
		Name               string                   `json:"name" validate:"required"`
		Type               StayType                 `json:"type" validate:"required"`
		Guests             int                      `json:"guests" validate:"required"`
		Rating             float64                  `json:"rating"`
		Amenities          map[amenity.Amenity]bool `json:"amenities" validate:"required"`
		House              string                   `json:"house" validate:"required"`
		Entrance           string                   `json:"entrance,omitempty"`
		Address            string                   `json:"address" validate:"required"`
		RoomsCount         string                   `json:"rooms_count" validate:"required"`
		BedsCount          string                   `json:"beds_count" validate:"required"`
		Price              string                   `json:"price" validate:"required"`
		Period             string                   `json:"period" validate:"required"`
		OwnersRules        string                   `json:"owners_rules" validate:"required"`
		CancellationPolicy string                   `json:"cancellation_policy" validate:"required"`
		DescribeProperty   string                   `json:"describe_property" validate:"required"`
		CreatedAt          time.Time                `json:"created_at"`
		UpdatedAt          time.Time                `json:"updated_at"`
	} // @name StayEntityFav

	Stay struct {
		ID                 uuid.UUID                `json:"id"`
		UserID             uuid.UUID                `json:"user_id"`
		LocationID         uuid.UUID                `json:"location_id"`
		Name               string                   `json:"name"`
		Type               StayType                 `json:"type"`
		Guests             int                      `json:"guests"`
		Rating             float64                  `json:"rating"`
		Amenities          map[amenity.Amenity]bool `json:"amenities"`
		House              string                   `json:"house"`
		Entrance           string                   `json:"entrance"`
		Address            string                   `json:"address"`
		RoomsCount         string                   `json:"rooms_count"`
		BedsCount          string                   `json:"beds_count"`
		Price              string                   `json:"price"`
		Period             string                   `json:"period"`
		OwnersRules        string                   `json:"owners_rules"`
		CancellationPolicy string                   `json:"cancellation_policy"`
		DescribeProperty   string                   `json:"describe_property"`
		CreatedAt          time.Time                `json:"created_at"`
		UpdatedAt          time.Time                `json:"updated_at"`
	} // @name Stay

	StayOccupied struct {
		ID                 uuid.UUID                `json:"id"`
		UserID             uuid.UUID                `json:"user_id"`
		LocationID         uuid.UUID                `json:"location_id"`
		Name               string                   `json:"name"`
		Type               StayType                 `json:"type"`
		Guests             int                      `json:"guests"`
		Rating             float64                  `json:"rating"`
		Amenities          map[amenity.Amenity]bool `json:"amenities" validate:"required"`
		House              string                   `json:"house"`
		Entrance           string                   `json:"entrance"`
		Address            string                   `json:"address"`
		RoomsCount         string                   `json:"rooms_count"`
		BedsCount          string                   `json:"beds_count"`
		Price              string                   `json:"price"`
		Period             string                   `json:"period"`
		OwnersRules        string                   `json:"owners_rules"`
		CancellationPolicy string                   `json:"cancellation_policy"`
		DescribeProperty   string                   `json:"describe_property"`
		CreatedAt          time.Time                `json:"created_at"`
		UpdatedAt          time.Time                `json:"updated_at"`
		ArrivedAt          time.Time                `json:"arrived_at"`
		DepartureAt        time.Time                `json:"departure_at"`
	} // @name StayOccupied

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

	// Filtration represents the filtering options for stays.
	Filtration struct {
		// LocationID is the UUID of the location to filter stays by. Required value.
		// @Param location_id query string true "Location ID" Example: "550e8400-e29b-41d4-a716-446655440001"
		LocationID uuid.UUID `json:"location_id" validate:"required"`

		// SortBy specifies the sorting order for the results. Omitempty value.
		// @Param sort_by query string false "Sort by options: Nil, Old, New, Highly Recommended, Lowly Recommended" Example: "New"
		SortBy sort.Sort `json:"sort_by" validate:"omitempty"`

		// PriceMin is the minimum price for filtering stays. Omitempty value.
		// Need both min and max values if you use it.
		// @Param price_min query float true "Minimum price" Example: 50.0
		PriceMin float32 `json:"price_min" validate:"omitempty"`

		// PriceMax is the maximum price for filtering stays. Omitempty value.
		// Need both min and max values if you use it.
		// @Param price_max query float true "Maximum price" Example: 200.0
		PriceMax float32 `json:"price_max" validate:"omitempty"`

		// NumberOfBedrooms specifies the number of bedrooms to filter stays. Omitempty value.
		// @Param number_of_bedrooms query int32 false "Number of bedrooms" Example: [1, 2]
		NumberOfBedrooms []int32 `json:"number_of_bedrooms" validate:"omitempty"`

		// Amenities is a map of amenities to filter stays by. Omitempty value.
		// Example: "amenities": {"Wi-fi": true, "Air conditioner": false}
		// @Param amenities query map[string]bool false "Amenities filter" Example: {"Wi-fi": true, "Air conditioner": false}
		Amenities map[amenity.Amenity]bool `json:"amenities" validate:"omitempty"`

		// Rating specifies the rating range for filtering stays. Omitempty value.
		// Need an array with a minimum length of 2. Example: [5, 4] or [5, 4, 3]
		// @Param rating query float false "Rating range" Example: [5, 4]
		Rating []float64 `json:"rating" validate:"omitempty"`
	} // @name Filtration
)

func (f *Filtration) SetDefaults() error {
	countPrice := 0
	if f.PriceMin == 0 {
		f.PriceMin = -1
		countPrice++
	}
	if f.PriceMax == 0 {
		f.PriceMax = -1
		countPrice++
	}
	if countPrice == 1 {
		return fmt.Errorf("filtration only supports both of price_min and price_max values")
	}
	return nil
}
