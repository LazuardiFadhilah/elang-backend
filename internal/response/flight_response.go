package response

import "github.com/LazuardiFadhilah/elang-backend/internal/domain"

type AirlineFlightResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DepatureFlightResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type ArrivalFlightResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Flight struct {
	ID                 string                 `json:"id"`
	Flight_code        string                 `json:"flight_code"`
	Airline            AirlineFlightResponse  `json:"airline"`
	Depature           DepatureFlightResponse `json:"depature_airport"`
	Arrival_airport_id ArrivalFlightResponse  `json:"arrival_airport"`
	Depature_time      string                 `json:"depature_time"`
	Arrival_time       string                 `json:"arrival_time"`
	Duration           string                 `json:"duration"`
	Is_transit         bool                   `json:"is_transit"`
	Transit_airport_id *domain.Airport        `json:"transit_airport_id"`
	Base_price         int                    `json:"base_price"`
}

type FlightResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Flight `json:"data"`
}
