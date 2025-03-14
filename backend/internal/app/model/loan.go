package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID             uint `gorm:"primaryKey"`                                    // Use uint for auto-incrementing IDs
	UserID         uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Ensure this matches the type of User.ID if it references it
	Amount         float64
	Interest       float64
	StartDate      time.Time `gorm:"type:date"` // Use time.Time for dates
	EndDate        time.Time `gorm:"type:date"` // Use time.Time for dates
	MonthlyPayment float64
	DownPayment    float64
	Comment        string `gorm:"type:varchar(255)"` // Define string length if necessary
	Type           string `gorm:"type:varchar(50)"`  // Mortgage, Car, Student, Personal
	Status         string `gorm:"type:varchar(50)"`  // Active, Paid, Overdue
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
