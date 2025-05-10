package response

import (
	"github.com/lib/pq"
)

// type FlightonFlightTier struct {
// 	ID                 string                 `json:"id"`
// 	Flight_code        string                 `json:"flight_code"`
// 	Airline            AirlineFlightResponse  `json:"airline"`
// 	Depature           DepatureFlightResponse `json:"depature_airport"`
// 	Arrival_airport_id ArrivalFlightResponse  `json:"arrival_airport"`
// 	Depature_time      string                 `json:"depature_time"`
// 	Arrival_time       string                 `json:"arrival_time"`
// 	Duration           string                 `json:"duration"`
// 	Is_transit         bool                   `json:"is_transit"`
// 	Transit_airport_id *TransitFlightResponse `json:"transit_airport"`
// 	Base_price         int                    `json:"base_price"`
// }

type FlightTier struct {
	ID         string         `json:"id"`
	Flight     string         `json:"flight_id"`
	Tier       string         `json:"tier"`
	Price      int            `json:"price"`
	Facilities pq.StringArray `json:"facilities"`
}

type FlightTierResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    FlightTier `json:"data"`
}
