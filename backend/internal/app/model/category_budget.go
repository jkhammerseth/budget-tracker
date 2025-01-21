package model

import (
	"time"
)

type CategoryBudget struct {
	ID            uint         `gorm:"primaryKey"`
	CategoryID    uint         `gorm:"not null"` // Foreign key to parent Category
	Category      Category     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubcategoryID *uint        `gorm:""` // Nullable, in case the budget applies only to the parent category
	Subcategory   *Subcategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID        uint         `gorm:"not null"` // User-specific budget
	Budget        float64      `gorm:"not null"` // Monthly budget amount
	Month         int          `gorm:"not null"` // Month (1-12)
	Year          int          `gorm:"not null"` // Year
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
