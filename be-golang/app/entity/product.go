package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"type:char(36);" json:"id"`
	Product    string    `gorm:"type:varchar(255)" json:"Product,omitempty"`
	Department string    `gorm:"type:varchar(255)" json:"Department,omitempty"`
	Price      float64   `gorm:"type:decimal(10,2)" json:"Price,omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"-"`
}
