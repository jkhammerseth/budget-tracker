package model

import (
	"time"
)

type Subcategory struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Name       string   `gorm:"type:varchar(255)" json:"name"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"-" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
