package service

import (
	"fmt"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
	"github.com/google/uuid"
)

type AirlineService interface {
	CreateAirline(airline *domain.Airline) (*domain.Airline, error)
	GetAllAirlines() ([]domain.Airline, error)
	GetAirlineByID(id string) (*domain.Airline, error)
	UpdateAirline(airline *domain.Airline) error
	DeleteAirline(id uuid.UUID) error
}

type airlineService struct {
	repo repository.AirlineRepository
}

func NewAirlineService(repo repository.AirlineRepository) AirlineService {
	return &airlineService{repo}
}

func (s *airlineService) CreateAirline(airline *domain.Airline) (*domain.Airline, error) {
	if airline.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	err := s.repo.Create(airline)
	if err != nil {
		return nil, err
	}
	return airline, nil
}

func (s *airlineService) GetAllAirlines() ([]domain.Airline, error) {
	airlines, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return airlines, nil
}

func (s *airlineService) GetAirlineByID(id string) (*domain.Airline, error) {
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

func (s *airlineService) UpdateAirline(airline *domain.Airline) error {
	existingAirline, err := s.repo.FindByID(airline.ID)
	fmt.Println("existingAirline", existingAirline)
	if err != nil {
		return err
	}

	if airline.Name == "" {
		airline.Name = existingAirline.Name
	}

	if airline.Logo_url == "" {
		airline.Logo_url = existingAirline.Logo_url
	}

	err = s.repo.Update(airline)
	if err != nil {
		return err
	}
	return nil
}

func (s *airlineService) DeleteAirline(id uuid.UUID) error {
	existingAirline, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	err = s.repo.Delete(existingAirline.ID)
	if err != nil {
		return err
	}
	return nil
}
