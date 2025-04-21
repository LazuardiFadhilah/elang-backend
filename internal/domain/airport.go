package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Airport struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name    string    `gorm:"type:varchar(100)" json:"name"`
	Code    string    `gorm:"type:varchar(10)" json:"code"`
	City    string    `gorm:"type:varchar(100)" json:"city"`
	Country string    `gorm:"type:varchar(100)" json:"country"`
}

func (a *Airport) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
