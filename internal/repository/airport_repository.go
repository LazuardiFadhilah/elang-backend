package repository

import (
	"github.com/LazuardiFadhilah/elang-backend/internal/domain"

	"gorm.io/gorm"
)

type AirportRepository interface {
	Create(airport *domain.Airport) error
	FindAll() ([]domain.Airport, error)
	FindByCode(code string) (*domain.Airport, error)
	FindByID(id uint) (*domain.Airport, error)
	Update(airport *domain.Airport) error
	Delete(id uint) error
}

type airportRepository struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) AirportRepository {
	return &airportRepository{db}
}

func (r *airportRepository) Create(airport *domain.Airport) error {
	err := r.db.Create(airport).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *airportRepository) FindAll() ([]domain.Airport, error) {
	var airports []domain.Airport
	err := r.db.Find(&airports).Error
	if err != nil {
		return nil, err
	}
	return airports, nil
}

func (r *airportRepository) FindByID(id uint) (*domain.Airport, error) {
	var airport domain.Airport
	err := r.db.First(&airport, id).Error
	if err != nil {
		return nil, err
	}
	return &airport, nil
}

func (r *airportRepository) Update(airport *domain.Airport) error {
	err := r.db.Save(airport).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *airportRepository) Delete(id uint) error {
	err := r.db.Delete(&domain.Airport{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *airportRepository) FindByCode(code string) (*domain.Airport, error) {
	var airport domain.Airport
	err := r.db.Where("code = ?", code).First(&airport).Error
	if err != nil {
		return nil, err
	}
	return &airport, nil
}
