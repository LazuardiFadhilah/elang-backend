package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlightService interface {
	CreateFlight(flight *domain.Flight) (*domain.Flight, error)
	FindAllFlights(filter domain.FlightFilter) ([]domain.Flight, error)
	FindByID(id string) (*domain.Flight, error)
	UpdateFlight(flight *domain.Flight) error
	DeleteFlight(id uuid.UUID) error
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

func (s *flightService) UpdateFlight(flight *domain.Flight) error {
	existingData, err := s.repo.FindByID(flight.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("fligth not Found")
		}
		return err
	}

	if flight.Flight_code == "" {
		flight.Flight_code = existingData.Flight_code
	}
	if flight.Airline_id == uuid.Nil {
		flight.Airline_id = existingData.Airline_id
		fmt.Println(flight.Airline_id)
	}

	if flight.Depature_airport_id == uuid.Nil {
		flight.Depature_airport_id = existingData.Depature_airport_id
	}
	if flight.Arrival_airport_id == uuid.Nil {
		flight.Arrival_airport_id = existingData.Arrival_airport_id
	}
	if flight.Depature_time == (time.Time{}) {
		flight.Depature_time = existingData.Depature_time
	}
	if flight.Arrival_time == (time.Time{}) {
		flight.Arrival_time = existingData.Arrival_time
	}
	if flight.Duration == "" {
		flight.Duration = existingData.Duration
	}

	return s.repo.Update(flight)
}

func (s *flightService) DeleteFlight(id uuid.UUID) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("airport not found")
		}
		return err
	}
	return s.repo.Delete(existing.ID)
}
