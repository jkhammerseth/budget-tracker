package model

import (
	"time"
)

type Expense struct {
	ID        uint   `gorm:"primaryKey"`        // Use uint for auto-incrementing IDs
	Name      string `gorm:"type:varchar(255)"` // Define string length if necessary
	Amount    float64
	Category  string    `gorm:"type:varchar(100)"` // Define string length if necessary
	Frequency string    `gorm:"type:varchar(50)"`  // one-time, daily, weekly, monthly
	Date      time.Time `gorm:"type:date"`         // Use time.Time for dates
	Paid      bool      `gorm:"type:boolean"`      // Use boolean for true/false values
	Comment   string    `gorm:"type:varchar(255)"` // Define string length if necessary
	UserID    uint      `gorm:"type:integer"`      // Ensure this matches the type of User.ID if it references it
}
