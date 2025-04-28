package domain

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type FlightTier struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Flight_id  uuid.UUID      `gorm:"type:uuid" json:"flight_id"`
	Tier       string         `gorm:"type:varchar(100)" json:"tier"`
	Price      int            `gorm:"type:int" json:"price"`
	Facilities pq.StringArray `gorm:"type:text[]" json:"facilities"`
}

func (a *FlightTier) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
