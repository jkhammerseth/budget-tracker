package model

import (
	"time"
)

type Category struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	Name          string        `gorm:"type:varchar(255)" json:"name"`
	Icon          string        `gorm:"type:varchar(255)" json:"icon"`
	Subcategories []Subcategory `gorm:"foreignKey:CategoryID" json:"subcategories"` // Explicit relationship
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
