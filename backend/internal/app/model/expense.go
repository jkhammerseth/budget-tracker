package model

import (
	"time"
)

type Expense struct {
	ID          uint   `gorm:"primaryKey"`        // Use uint for auto-incrementing IDs
	Name        string `gorm:"type:varchar(255)"` // Define string length if necessary
	Amount      float64
	Category    string     `gorm:"type:varchar(100)"` // Define string length if necessary
	Frequency   string     `gorm:"type:varchar(50)"`  // one-time, daily, weekly, monthly, yearly
	StartDate   *time.Time `gorm:"type:date"`         // When recurring payments start
	EndDate     *time.Time `gorm:"type:date"`         // When recurring payments end
	PaymentDate *time.Time `gorm:"type:date"`         // Specific day of the month/year (e.g., 1 for 1st of the month)
	Paid        bool       `gorm:"type:boolean"`      // Use boolean for true/false values
	Comment     string     `gorm:"type:varchar(255)"` // Define string length if necessary
	UserID      uint       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LoanID      *uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Ensure this matches the type of Loan.ID if it references it
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
