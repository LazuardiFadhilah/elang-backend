package repository

import (
	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FlightTierRepository interface {
	Create(flightTier *domain.FlightTier) error
	FindTierByFlightID(flightID uuid.UUID) ([]domain.FlightTier, error)
	Update(flightTier *domain.FlightTier) error
	Delete(id uuid.UUID) error
}

type flightTierRepository struct {
	db *gorm.DB
}

func NewFlightTierRepository(db *gorm.DB) FlightTierRepository {
	return &flightTierRepository{db}
}

func (r *flightTierRepository) Create(flightTier *domain.FlightTier) error {
	err := r.db.Create(flightTier).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *flightTierRepository) FindTierByFlightID(flightID uuid.UUID) ([]domain.FlightTier, error) {
	var flightTiers []domain.FlightTier
	err := r.db.Where("flight_id = ?", flightID).Find(&flightTiers).Error
	if err != nil {
		return nil, err
	}
	return flightTiers, nil
}

func (r *flightTierRepository) Update(flightTier *domain.FlightTier) error {
	err := r.db.Save(flightTier).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *flightTierRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&domain.FlightTier{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
