package handler

import (
	"net/http"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/response"
	"github.com/LazuardiFadhilah/elang-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AirportHandler struct {
	service service.AirportService
}

func NewAirportHandler(service service.AirportService) *AirportHandler {
	return &AirportHandler{service: service}
}

func (h *AirportHandler) CreateAirport(c *gin.Context) {
	var input domain.Airport
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_JSON_BODY",
			"status":  http.StatusBadRequest,
		})
		return
	}

	existingAirport, err := h.service.GetAirportByCode(input.Code)
	if err == nil && existingAirport != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Message": "AIRPORT_ALREADY_EXISTS",
			"status":  http.StatusConflict,
		})
		return
	}

	createdAirport, err := h.service.CreateAirport(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "FAILED_TO_CREATE_AIRPORT",
			"status":  http.StatusBadRequest,
		})
		return
	}

	response := response.AirportResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS_CREATE_AIRPORT",
		Data: response.Airport{
			ID:      createdAirport.ID.String(),
			Name:    createdAirport.Name,
			Code:    createdAirport.Code,
			City:    createdAirport.City,
			Country: createdAirport.Country,
		},
	}

	c.JSON(http.StatusCreated, response)
}
