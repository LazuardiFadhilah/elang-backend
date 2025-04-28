package handler

import (
	"fmt"
	"net/http"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/response"
	"github.com/LazuardiFadhilah/elang-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type FlightHandler struct {
	service        service.FlightService
	airlineService service.AirlineService
	airportService service.AirportService
}

func NewFlightHandler(service service.FlightService, airlineService service.AirlineService, airportService service.AirportService) *FlightHandler {
	return &FlightHandler{service: service, airlineService: airlineService, airportService: airportService}
}

func (h *FlightHandler) CreateFlight(c *gin.Context) {
	var input domain.Flight
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_JSON_BODY",
			"status":  http.StatusBadRequest,
		})
		return
	}
	if input.Flight_code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "FLIGHT_CODE_REQUIRED",
			"status":  http.StatusBadRequest,
		})
		return
	}
	if input.Airline_id.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "AIRLINE_ID_REQUIRED",
			"status":  http.StatusBadRequest,
		})
		return
	}
	if input.Depature_airport_id.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "DEPATURE_AIRPORT_ID_REQUIRED",
			"status":  http.StatusBadRequest,
		})
		return
	}
	if input.Arrival_airport_id.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ARRIVAL_AIRPORT_ID_REQUIRED",
			"status":  http.StatusBadRequest,
		})
		return
	}
	if input.Depature_time.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "DEPATURE_TIME_REQUIRED",
			"status":  http.StatusBadRequest,
		})
		return
	}
	if input.Arrival_time.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "ARRIVAL_TIME_REQUIRED",
			"status":  http.StatusBadRequest,
		})
		return
	}

	duration := input.Arrival_time.Sub(input.Depature_time)
	if duration < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_DEPATURE_ARRIVAL_TIME",
			"status":  http.StatusBadRequest,
		})
		return
	}
	input.Duration = duration

	if input.Is_transit {
		if input.Transit_airport_id == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "TRANSIT_AIRPORT_ID_REQUIRED",
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	createdFlight, err := h.service.CreateFlight(&input)
	if err != nil {
		fmt.Println("Error creating flight:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "INTERNAL_SERVER_ERROR",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	existingAirline, err := h.airlineService.GetAirlineByID(input.Airline_id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "INTERNAL_SERVER_ERROR",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	depatureAirport, err := h.airportService.GetAirportByID(input.Depature_airport_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "INTERNAL_SERVER_ERROR",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	arrivalAirport, err := h.airportService.GetAirportByID(input.Arrival_airport_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "INTERNAL_SERVER_ERROR",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	response := response.FlightResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_CREATE_FLIGHT",
		Data: response.Flight{
			ID:          createdFlight.ID.String(),
			Flight_code: createdFlight.Flight_code,
			Airline: response.AirlineFlightResponse{
				ID:   createdFlight.Airline_id.String(),
				Name: existingAirline.Name},
			Depature: response.DepatureFlightResponse{
				ID:   depatureAirport.ID.String(),
				Name: depatureAirport.Name,
				Code: depatureAirport.Code,
			},
			Arrival_airport_id: response.ArrivalFlightResponse{
				ID:   arrivalAirport.ID.String(),
				Name: arrivalAirport.Name,
				Code: arrivalAirport.Code,
			},
			Depature_time:      createdFlight.Depature_time.Format("2006-01-02 15:04:05"),
			Arrival_time:       createdFlight.Arrival_time.Format("2006-01-02 15:04:05"),
			Duration:           createdFlight.Duration.String(),
			Is_transit:         createdFlight.Is_transit,
			Transit_airport_id: createdFlight.Transit_airport,
			Base_price:         createdFlight.Base_price,
		},
	}
	c.JSON(http.StatusOK, response)

}
