package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LazuardiFadhilah/elang-backend/config"
	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/handler"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
	"github.com/LazuardiFadhilah/elang-backend/internal/router"
	"github.com/LazuardiFadhilah/elang-backend/internal/service"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	// Auto migrate models
	config.DB.AutoMigrate(&domain.Airport{})
	config.DB.AutoMigrate(&domain.Airline{})
	config.DB.AutoMigrate(&domain.Flight{})

	// Auto migrate models
	if err := config.DB.AutoMigrate(&domain.Airport{}, &domain.Airline{}, &domain.Flight{}); err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	} else {
		fmt.Println("Auto migration successful")
	}

	// Inisialisasi Repository, Service, Handler
	airportRepo := repository.NewAirportRepository(config.DB)
	airportService := service.NewAirportService(airportRepo)
	airportHandler := handler.NewAirportHandler(airportService)

	airlineRepo := repository.NewAirlineRepository(config.DB)
	airlineService := service.NewAirlineService(airlineRepo)
	airlineHandler := handler.NewAirlineHandler(airlineService)

	flightRepo := repository.NewFlightRepository(config.DB)
	flightService := service.NewFlightService(flightRepo, airportRepo, airlineRepo)
	flightHandler := handler.NewFlightHandler(flightService, airlineService, airportService)

	// Setup Router
	r := router.SetupRouter(airportHandler, airlineHandler, flightHandler)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running at http://localhost:" + port)
	log.Fatal(r.Run(":" + port))
}
