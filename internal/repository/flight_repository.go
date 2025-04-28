package repository

import (
	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type FlightRepository interface {
	Create(airport *domain.Flight) error
	FindAll() ([]domain.Flight, error)
	FindByCode(code string) (*domain.Flight, error)
	FindByID(id uuid.UUID) (*domain.Flight, error)
	Update(airport *domain.Flight) error
	Delete(id uuid.UUID) error
}

type flightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) FlightRepository {
	return &flightRepository{db}
}

func (r *flightRepository) Create(airport *domain.Flight) error {
	err := r.db.Create(airport).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *flightRepository) FindAll() ([]domain.Flight, error) {
	var flights []domain.Flight
	err := r.db.Preload("Airline").Preload("Depature_airport").Preload("Arrival_airport").Preload("Transit_airport").Preload("Flight_tiers").Find(&flights).Error
	if err != nil {
		return nil, err
	}
	return flights, nil
}

func (r *flightRepository) FindByID(id uuid.UUID) (*domain.Flight, error) {
	var flight domain.Flight
	err := r.db.Preload("Airline").Preload("Depature_airport").Preload("Arrival_airport").Preload("Transit_airport").Preload("Flight_tiers").Find(&flight, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &flight, nil
}

func (r *flightRepository) Update(airport *domain.Flight) error {
	err := r.db.Save(airport).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *flightRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&domain.Airport{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *flightRepository) FindByCode(code string) (*domain.Flight, error) {
	var flight domain.Flight
	err := r.db.Where("code = ?", code).First(&flight).Error
	if err != nil {
		return nil, err
	}
	return &flight, nil
}
