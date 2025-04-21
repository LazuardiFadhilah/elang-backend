package service

import (
	"fmt"

	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/LazuardiFadhilah/elang-backend/internal/repository"
)

type AirportService interface {
	CreateAirport(airport *domain.Airport) (*domain.Airport, error)
	GetAllAirports() ([]domain.Airport, error)
	GetAirportByCode(code string) (*domain.Airport, error)
	GetAirportByID(id uint) (*domain.Airport, error)
	UpdateAirport(airport *domain.Airport) error
	DeleteAirport(id uint) error
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

func (s *airportService) GetAirportByID(id uint) (*domain.Airport, error) {
	return s.repo.FindByID(id)
}

func (s *airportService) UpdateAirport(airport *domain.Airport) error {
	return s.repo.Update(airport)
}

func (s *airportService) DeleteAirport(id uint) error {
	return s.repo.Delete(id)
}
