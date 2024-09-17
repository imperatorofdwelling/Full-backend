package models

import "time"

type (
	LocationEntity struct {
		ID              int       `json:"id"`
		City            string    `json:"city"`
		FederalDistrict string    `json:"federal_district,omitempty"`
		FiasID          string    `json:"fias_id,omitempty"`
		KladrID         string    `json:"kladr_id,omitempty"`
		Lat             string    `json:"lat,omitempty"`
		Lon             string    `json:"lon,omitempty"`
		Okato           string    `json:"okato,omitempty"`
		Oktmo           string    `json:"oktmo,omitempty"`
		Population      float32   `json:"population,omitempty"`
		RegionIsoCode   string    `json:"region_iso_code,omitempty"`
		RegionName      string    `json:"region_name,omitempty"`
		CreatedAt       time.Time `json:"createdAt,omitempty"`
		UpdatedAt       time.Time `json:"updatedAt,omitempty"`
	}

	Location struct {
		ID              int       `json:"id"`
		City            string    `json:"city"`
		FederalDistrict string    `json:"federal_district"`
		FiasID          string    `json:"fias_id"`
		KladrID         string    `json:"kladr_id"`
		Lat             string    `json:"lat"`
		Lon             string    `json:"lon"`
		Okato           string    `json:"okato"`
		Oktmo           string    `json:"oktmo"`
		Population      float32   `json:"population"`
		RegionIsoCode   string    `json:"region_iso_code"`
		RegionName      string    `json:"region_name"`
		CreatedAt       time.Time `json:"createdAt"`
		UpdatedAt       time.Time `json:"updatedAt"`
	}
)
