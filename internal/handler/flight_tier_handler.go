package handler

import (
	"net/http"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/response"
	"github.com/LazuardiFadhilah/elang-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FlightTierHandler struct {
	service       service.FlightTierService
	flightService service.FlightService
}

func NewFlightTierHandler(service service.FlightTierService, flightService service.FlightService) *FlightTierHandler {
	return &FlightTierHandler{
		service:       service,
		flightService: flightService,
	}
}

func (h *FlightTierHandler) CreateFlightTier(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "INVALID_FLIGHT_ID"})
		return
	}
	recentFlight, err := h.flightService.FindByID(id.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FLIGHT_NOT_FOUND"})
		return
	}
	var input domain.FlightTier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_INPUT"})
		return
	}

	input.Flight_id = recentFlight.ID

	createdFlightTier, err := h.service.CreateFlightTier(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FAILED_TO_CREATE_FLIGHT_TIER"})
		return
	}
	response := response.FlightTierResponse{
		Status:  http.StatusOK,
		Message: "FLIGHT_TIER_CREATED",
		Data: response.FlightTier{
			ID:         createdFlightTier.ID.String(),
			Flight:     createdFlightTier.Flight_id.String(),
			Tier:       createdFlightTier.Tier,
			Price:      createdFlightTier.Price,
			Facilities: createdFlightTier.Facilities},
	}
	c.JSON(http.StatusOK, response)
}
