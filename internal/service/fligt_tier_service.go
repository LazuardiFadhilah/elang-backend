package service

import (
	"fmt"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
	"github.com/google/uuid"
)

type FlightTierService interface {
	CreateFlightTier(flightTier *domain.FlightTier) (*domain.FlightTier, error)
	FindAllFlightTiers(flight_id string) ([]domain.FlightTier, error)
}

type flightTierService struct {
	repo       repository.FlightTierRepository
	flightRepo repository.FlightRepository
}

func NewFlightTierService(repo repository.FlightTierRepository, flightRepo repository.FlightRepository) FlightTierService {
	return &flightTierService{
		repo:       repo,
		flightRepo: flightRepo,
	}
}

func (s *flightTierService) CreateFlightTier(flightTier *domain.FlightTier) (*domain.FlightTier, error) {
	_, err := s.flightRepo.FindByID(flightTier.Flight_id)
	if err != nil {
		return nil, fmt.Errorf("flight not found: %w", err)
	}

	err = s.repo.Create(flightTier)
	if err != nil {
		return nil, fmt.Errorf("failed to create flight tier: %w", err)
	}
	return flightTier, nil
}

func (s *flightTierService) FindAllFlightTiers(flight_id string) ([]domain.FlightTier, error) {
	uuid, err := uuid.Parse(flight_id)
	if err != nil {
		return nil, fmt.Errorf("invalid flight ID: %w", err)
	}

	flightTiers, err := s.repo.FindTierByFlightID(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to find flight tiers: %w", err)
	}
	return flightTiers, nil
}

func (s *flightTierService) UpdateFlightTier(flightTier *domain.FlightTier) error {
	_, err := s.flightRepo.FindByID(flightTier.Flight_id)
	if err != nil {
		return fmt.Errorf("flight not found: %w", err)
	}

	err = s.repo.Update(flightTier)
	if err != nil {
		return fmt.Errorf("failed to update flight tier: %w", err)
	}
	return nil
}

func (s *flightTierService) DeleteFlightTier(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid flight tier ID: %w", err)
	}

	err = s.repo.Delete(uuid)
	if err != nil {
		return fmt.Errorf("failed to delete flight tier: %w", err)
	}
	return nil
}
