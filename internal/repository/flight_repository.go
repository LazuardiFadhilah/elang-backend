package repository

import (
	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type FlightRepository interface {
	Create(airport *domain.Flight) error
	FindAll(filter domain.FlightFilter) ([]domain.Flight, error)
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

func (r *flightRepository) FindAll(filter domain.FlightFilter) ([]domain.Flight, error) {
	var flights []domain.Flight
	query := r.db.Model(&domain.Flight{})
	if filter.Code != "" {
		query = query.Where("flight_code = ?", filter.Code)
	}
	if filter.Depature_airport_id != "" {
		query = query.Where("depature_airport_id = ?", filter.Depature_airport_id)
	}
	if filter.Arrival_airport_id != "" {
		query = query.Where("arrival_airport_id = ?", filter.Arrival_airport_id)
	}
	if filter.Airline_id != "" {
		query = query.Where("airline_id = ?", filter.Airline_id)
	}
	if filter.Is_transit {
		query = query.Where("is_transit = ?", filter.Is_transit)
	}
	if filter.MinPrice != "" && filter.MaxPrice != "" {
		query = query.Where("base_price BETWEEN ? AND ?", filter.MinPrice, filter.MaxPrice)
	} else if filter.MinPrice != "" {
		query = query.Where("base_price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice != "" {
		query = query.Where("base_price <= ?", filter.MaxPrice)
	}
	err := query.Find(&flights).Error
	if err != nil {
		return nil, err
	}
	return flights, nil
}

func (r *flightRepository) FindByID(id uuid.UUID) (*domain.Flight, error) {
	var flight domain.Flight
	err := r.db.Find(&flight, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &flight, nil
}

func (r *flightRepository) Update(flight *domain.Flight) error {
	err := r.db.Save(flight).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *flightRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&domain.Flight{}, id).Error
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
