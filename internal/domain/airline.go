package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Airline struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"type:varchar(100)" json:"name"`
	Logo_url string    `gorm:"type:varchar(255)" json:"logo_url"`
}

func (a *Airline) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
