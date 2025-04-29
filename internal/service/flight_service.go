package service

import (
	"fmt"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
	"github.com/google/uuid"
)

type FlightService interface {
	CreateFlight(flight *domain.Flight) (*domain.Flight, error)
	FindAllFlights(filter domain.FlightFilter) ([]domain.Flight, error)
	FindByID(id string) (*domain.Flight, error)
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

func (s *flightService) FindAllFlights(filter domain.FlightFilter) ([]domain.Flight, error) {
	return s.repo.FindAll(filter)
}

func (s *flightService) FindByID(id string) (*domain.Flight, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	airline, err := s.repo.FindByID(uuidID)
	if err != nil {
		return nil, err
	}
	return airline, nil
}
