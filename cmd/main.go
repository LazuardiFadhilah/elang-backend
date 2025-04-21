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

	// Inisialisasi Repository, Service, Handler
	airportRepo := repository.NewAirportRepository(config.DB)
	airportService := service.NewAirportService(airportRepo)
	airportHandler := handler.NewAirportHandler(airportService)

	// Setup Router
	r := router.SetupRouter(airportHandler)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running at http://localhost:" + port)
	log.Fatal(r.Run(":" + port))
}
