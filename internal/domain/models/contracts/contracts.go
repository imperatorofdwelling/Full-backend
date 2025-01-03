package contracts

import "time"

type (
	Contract struct {
		UserID    string    `json:"user_id"`
		StayID    string    `json:"stay_id"`
		Price     float64   `json:"price"`
		DateStart time.Time `json:"date_start"`
		DateEnd   time.Time `json:"date_end"`
		Square    float64   `json:"square"`
		Street    string    `json:"street"`
		House     string    `json:"house"`
		Entrance  string    `json:"entrance"`
		Floor     string    `json:"floor,omitempty"`
		Room      string    `json:"room,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} // @name Contract

	ContractEntity struct {
		UserName  string
		StayName  string
		Price     float64
		DateStart time.Time
		DateEnd   time.Time
	} // @name ContractEntity

	ContractResponse struct {
		DateStart string `json:"dateStart" example:"2023-01-01T00:00:00Z"`
		DateEnd   string `json:"dateEnd" example:"2023-01-10T00:00:00Z"`
	} // @name ContractResponse

)
