package handler

import (
	"net/http"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/response"
	"github.com/LazuardiFadhilah/elang-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AirlineHandler struct {
	airline service.AirlineService
}

func NewAirlineHandler(airline service.AirlineService) *AirlineHandler {
	return &AirlineHandler{
		airline: airline,
	}
}

func (h *AirlineHandler) CreateAirline(c *gin.Context) {
	var input domain.Airline
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"Message": "INVALID_JSON_BODY",
			"status":  http.StatusBadRequest,
		})
		return
	}

	createdAirline, err := h.airline.CreateAirline(&input)
	if err != nil {

		if err.Error() == "name is required" {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "NAME_IS_REQUIRED",
				"status":  http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "FAILED_TO_CREATE_AIRLINE",
			"status":  http.StatusBadRequest,
		})
		return
	}

	response := response.AirlineResponse{
		Status:  http.StatusCreated,
		Message: "SUCCESS_CREATE_AIRLINE",
		Data: response.Airline{
			ID:       createdAirline.ID.String(),
			Name:     createdAirline.Name,
			Logo_url: createdAirline.Logo_url,
		},
	}
	c.JSON(http.StatusCreated, response)
}

func (h *AirlineHandler) GetAllAirlines(c *gin.Context) {
	airlines, err := h.airline.GetAllAirlines()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "FAILED_TO_GET_AIRLINES",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	var airlineResponses []response.Airline
	for _, airline := range airlines {
		airlineResponses = append(airlineResponses, response.Airline{
			ID:       airline.ID.String(),
			Name:     airline.Name,
			Logo_url: airline.Logo_url,
		})
	}

	response := response.AirlinesResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_GET_AIRLINES",
		Count:   len(airlineResponses),
		Data:    airlineResponses,
	}
	c.JSON(http.StatusOK, response)
}

func (h *AirlineHandler) GetAirlineByID(c *gin.Context) {
	id := c.Param("id")
	airline, err := h.airline.GetAirlineByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "AIRLINE_NOT_FOUND",
			"status":  http.StatusNotFound,
		})
		return
	}

	response := response.AirlineResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_GET_AIRLINE",
		Data: response.Airline{
			ID:       airline.ID.String(),
			Name:     airline.Name,
			Logo_url: airline.Logo_url,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (h *AirlineHandler) UpdateAirline(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_AIRLINE_ID",
			"status":  http.StatusBadRequest,
		})
		return
	}

	var input domain.Airline
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"Message": "INVALID_JSON_BODY",
			"status":  http.StatusBadRequest,
		})
		return
	}
	input.ID = id

	err = h.airline.UpdateAirline(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "FAILED_TO_UPDATE_AIRLINE",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	response := response.AirlineResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_UPDATE_AIRLINE",
		Data: response.Airline{
			ID:       input.ID.String(),
			Name:     input.Name,
			Logo_url: input.Logo_url,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (h *AirlineHandler) DeleteAirline(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "INVALID_AIRLINE_ID",
			"status":  http.StatusBadRequest,
		})
		return
	}

	existingAirline, err := h.airline.GetAirlineByID(id.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "AIRLINE_NOT_FOUND",
			"status":  http.StatusNotFound,
		})
		return
	}

	err = h.airline.DeleteAirline(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": "FAILED_TO_DELETE_AIRLINE",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	response := response.AirlineResponse{
		Status:  http.StatusOK,
		Message: "SUCCESS_DELETE_AIRLINE",
		Data: response.Airline{
			ID:       id.String(),
			Name:     existingAirline.Name,
			Logo_url: existingAirline.Logo_url,
		},
	}
	c.JSON(http.StatusOK, response)
}
