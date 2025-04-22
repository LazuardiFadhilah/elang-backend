package handler

import (
	"net/http"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/response"
	"github.com/LazuardiFadhilah/elang-backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *AirportHandler) GetAllAirports(c *gin.Context) {
	airports, err := h.service.GetAllAirports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "FAILED_TO_GET_AIRPORTS",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	var airportResponses []response.Airport
	for _, airport := range airports {
		airportResponses = append(airportResponses, response.Airport{
			ID:      airport.ID.String(),
			Name:    airport.Name,
			Code:    airport.Code,
			City:    airport.City,
			Country: airport.Country,
		})
	}

	response := response.AirportListResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_GET_AIRPORTS",
		Count:   len(airportResponses),
		Data:    airportResponses,
	}
	c.JSON(http.StatusOK, response)
}

func (h *AirportHandler) GetAirportByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_AIRPORT_ID",
			"status":  http.StatusBadRequest,
		})
		return
	}

	airport, err := h.service.GetAirportByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "AIRPORT_NOT_FOUND",
			"status":  http.StatusNotFound,
		})
		return
	}

	response := response.AirportResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_GET_AIRPORT",
		Data: response.Airport{
			ID:      airport.ID.String(),
			Name:    airport.Name,
			Code:    airport.Code,
			City:    airport.City,
			Country: airport.Country,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (h *AirportHandler) UpdateAirport(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_AIRPORT_ID",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var input domain.Airport
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_JSON_BODY",
			"status":  http.StatusBadRequest,
		})
		return
	}

	input.ID = id

	err = h.service.UpdateAirport(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "FAILED_TO_UPDATE_AIRPORT",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	response := response.AirportResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_UPDATE_AIRPORT",
		Data: response.Airport{
			ID:      input.ID.String(),
			Name:    input.Name,
			Code:    input.Code,
			City:    input.City,
			Country: input.Country,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (h *AirportHandler) DeleteAirport(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_AIRPORT_ID",
			"status":  http.StatusBadRequest,
		})
		return
	}

	existingAirport, err := h.service.GetAirportByID(id)
	if err != nil {
		if err.Error() == "airport not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"Message": "AIRPORT_NOT_FOUND",
				"status":  http.StatusNotFound,
			})
			return
		}
	}

	err = h.service.DeleteAirport(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "FAILED_TO_DELETE_AIRPORT",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	response := response.AirportResponse{
		Status:  http.StatusNoContent,
		Message: "SUCCESS_DELETE_AIRPORT",
		Data: response.Airport{
			ID:      existingAirport.ID.String(),
			Name:    existingAirport.Name,
			Code:    existingAirport.Code,
			City:    existingAirport.City,
			Country: existingAirport.Country,
		},
	}
	c.JSON(http.StatusNoContent, response)
}
