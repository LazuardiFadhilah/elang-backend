package service

import (
	"fmt"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
)

type FlightService interface {
	CreateFlight(flight *domain.Flight) (*domain.Flight, error)
}

type flightService struct {
	repo        repository.FlightRepository
	airlineRepo repository.AirlineRepository
	airportRepo repository.AirportRepository
}

func NewFlightService(repo repository.FlightRepository, airportRepo repository.AirportRepository, airlineRepo repository.AirlineRepository) FlightService {
	return &flightService{
		repo:        repo,
		airlineRepo: airlineRepo,
		airportRepo: airportRepo,
	}
}

func (s *flightService) CreateFlight(flight *domain.Flight) (*domain.Flight, error) {
	_, err := s.airlineRepo.FindByID(flight.Airline_id)
	if err != nil {
		return nil, fmt.Errorf("airline not found: %w", err)
	}
	_, err = s.airportRepo.FindByID(flight.Depature_airport_id)
	if err != nil {
		return nil, fmt.Errorf("depature airport not found: %w", err)
	}
	_, err = s.airportRepo.FindByID(flight.Arrival_airport_id)
	if err != nil {
		return nil, fmt.Errorf("arrival airport not found: %w", err)
	}

	err = s.repo.Create(flight)
	if err != nil {
		return nil, fmt.Errorf("failed to create flight: %w", err)
	}
	return flight, nil
}
