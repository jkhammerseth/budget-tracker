package model

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID             uint `gorm:"primaryKey"`   // Use uint for auto-incrementing IDs
	UserID         uint `gorm:"type:integer"` // Ensure this matches the type of User.ID if it references it
	Amount         float64
	Interest       float64
	StartDate      time.Time `gorm:"type:date"` // Use time.Time for dates
	EndDate        time.Time `gorm:"type:date"` // Use time.Time for dates
	MonthlyPayment float64
	DownPayment    float64
	Comment        string `gorm:"type:varchar(255)"` // Define string length if necessary
	Remaining      float64
	Type           string `gorm:"type:varchar(50)"` // Mortgage, Car, Student, Personal
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Status         string         `gorm:"type:varchar(50)"` // Active, Paid, Overdue
	ExpenseIds     []uint
}
