package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	ID                  uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Flight_code         string     `gorm:"type:varchar(10)" json:"flight_code"`
	Airline_id          uuid.UUID  `gorm:"type:uuid" json:"airline_id"`
	Depature_airport_id uuid.UUID  `gorm:"type:uuid" json:"depature_airport_id"`
	Arrival_airport_id  uuid.UUID  `gorm:"type:uuid" json:"arrival_airport_id"`
	Depature_time       time.Time  `gorm:"type:timestamptz" json:"depature_time"`
	Arrival_time        time.Time  `gorm:"type:timestamptz" json:"arrival_time"`
	Duration            string     `gorm:"type:interval" json:"duration"`
	Is_transit          bool       `gorm:"type:boolean" json:"is_transit"`
	Transit_airport_id  *uuid.UUID `gorm:"type:uuid" json:"transit_airport_id"`
	Base_price          int        `gorm:"type:int" json:"base_price"`

	// Relationships
	Airline          Airline      `gorm:"foreignKey:airline_id" json:"airline"`
	Depature_airport Airport      `gorm:"foreignKey:depature_airport_id" json:"depature_airport"`
	Arrival_airport  Airport      `gorm:"foreignKey:arrival_airport_id" json:"arrival_airport"`
	Transit_airport  *Airport     `gorm:"foreignKey:transit_airport_id" json:"transit_airport"`
	Flight_tiers     []FlightTier `gorm:"foreignKey:flight_id" json:"flight_tiers"`
}

func (a *Flight) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

type FlightFilter struct {
	Depature_airport_id string
	Arrival_airport_id  string
	Airline_id          string
	Code                string
	Is_transit          bool
	MinPrice            string
	MaxPrice            string
}
