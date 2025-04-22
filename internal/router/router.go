package router

import (
	"github.com/LazuardiFadhilah/elang-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(airportHandler *handler.AirportHandler) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/airports", airportHandler.CreateAirport)
		v1.GET("/airports", airportHandler.GetAllAirports)
		v1.GET("/airports/:id", airportHandler.GetAirportByID)
		v1.PUT("/airports/:id", airportHandler.UpdateAirport)
		v1.DELETE("/airports/:id", airportHandler.DeleteAirport)
		// v1.GET("/airports", ...) <-- nanti tinggal tambahkan
	}

	return r
}
