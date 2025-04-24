package repository

import (
	"github.com/LazuardiFadhilah/elang-backend/internal/domain"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type AirlineRepository interface {
	Create(airline *domain.Airline) error
	FindAll() ([]domain.Airline, error)
	FindByID(id uuid.UUID) (*domain.Airline, error)
	Update(airline *domain.Airline) error
	Delete(id uuid.UUID) error
}

type airlineRepository struct {
	db *gorm.DB
}

func NewAirlineRepository(db *gorm.DB) AirlineRepository {
	return &airlineRepository{db}
}

func (r *airlineRepository) Create(airline *domain.Airline) error {
	err := r.db.Create(airline).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *airlineRepository) FindAll() ([]domain.Airline, error) {
	var airlines []domain.Airline
	err := r.db.Find(&airlines).Error
	if err != nil {
		return nil, err
	}
	return airlines, nil
}

func (r *airlineRepository) FindByID(id uuid.UUID) (*domain.Airline, error) {
	var airline domain.Airline
	err := r.db.First(&airline, id).Error
	if err != nil {
		return nil, err
	}
	return &airline, nil
}

func (r *airlineRepository) Update(airline *domain.Airline) error {
	err := r.db.Save(airline).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *airlineRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&domain.Airline{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
