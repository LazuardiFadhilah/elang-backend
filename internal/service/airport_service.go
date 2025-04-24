package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
)

type AirportService interface {
	CreateAirport(airport *domain.Airport) (*domain.Airport, error)
	GetAllAirports() ([]domain.Airport, error)
	GetAirportByCode(code string) (*domain.Airport, error)
	GetAirportByID(id uuid.UUID) (*domain.Airport, error)
	UpdateAirport(airport *domain.Airport) error
	DeleteAirport(id uuid.UUID) error
}

type airportService struct {
	repo repository.AirportRepository
}

func NewAirportService(repo repository.AirportRepository) AirportService {
	return &airportService{repo}
}

func (s *airportService) GetAirportByCode(code string) (*domain.Airport, error) {
	return s.repo.FindByCode(code)
}
func (s *airportService) CreateAirport(airport *domain.Airport) (*domain.Airport, error) {
	// Validasi manual
	if airport.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if airport.Code == "" {
		return nil, fmt.Errorf("code is required")
	}
	if airport.City == "" {
		return nil, fmt.Errorf("city is required")
	}
	if airport.Country == "" {
		return nil, fmt.Errorf("country is required")
	}

	// Optional: validasi panjang kode (misal harus 3 huruf kayak CGK)
	if len(airport.Code) != 3 {
		return nil, fmt.Errorf("code must be 3 characters")
	}

	err := s.repo.Create(airport)
	if err != nil {
		return nil, err
	}
	return airport, nil
}

func (s *airportService) GetAllAirports() ([]domain.Airport, error) {
	return s.repo.FindAll()
}

func (s *airportService) GetAirportByID(id uuid.UUID) (*domain.Airport, error) {
	return s.repo.FindByID(id)
}

func (s *airportService) UpdateAirport(airport *domain.Airport) error {
	existing, err := s.repo.FindByID(airport.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("airport not found")
		}
		return err
	}

	if airport.Name == "" {
		airport.Name = existing.Name
	}

	if airport.Code == "" {
		airport.Code = existing.Code
	}
	if len(airport.Code) != 3 || strings.ToUpper(airport.Code) != airport.Code {
		return fmt.Errorf("code must be 3 uppercase characters")
	}
	if airport.City == "" {
		airport.City = existing.City
	}
	if airport.Country == "" {
		airport.Country = existing.Country
	}

	return s.repo.Update(airport)
}

func (s *airportService) DeleteAirport(id uuid.UUID) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("airport not found")
		}
		return err
	}
	return s.repo.Delete(existing.ID)
}
