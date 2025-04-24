package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flight struct {
	ID                  uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	Flight_code         string        `gorm:"type:varchar(10)" json:"flight_code"`
	Airline_id          uuid.UUID     `gorm:"type:uuid;uniqueIndex" json:"airline_id"`
	Depature_airport_id uuid.UUID     `gorm:"type:uuid" json:"departure_airport_id"`
	Arrival_airport_id  uuid.UUID     `gorm:"type:uuid" json:"arrival_airport_id"`
	Depature_time       time.Time     `gorm:"type:timestapmtz" json:"departure_time"`
	Arrival_time        time.Time     `gorm:"type:timestapmtz" json:"arrival_time"`
	Duration            time.Duration `gorm:"type:interval" json:"duration"`
	Is_transit          bool          `gorm:"type:boolean" json:"is_transit"`
	Transit_airport_id  uuid.UUID     `gorm:"type:uuid" json:"transit_airport_id"`
	Base_price          int           `gorm:"type:int" json:"base_price"`
}

func (a *Flight) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
