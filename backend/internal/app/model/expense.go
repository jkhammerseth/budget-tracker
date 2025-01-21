package model

import (
	"time"
)

type Expense struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	UniqueHash    string       `gorm:"uniqueIndex"`
	Name          string       `gorm:"type:varchar(255)" json:"name"`
	Amount        float64      `json:"amount"`
	CategoryID    uint         `gorm:"not null" json:"category_id"`
	Category      Category     `json:"category"`
	SubcategoryID *uint        `json:"subcategory_id,omitempty"`
	Subcategory   *Subcategory `json:"subcategory,omitempty"`
	Frequency     string       `gorm:"type:varchar(50)" json:"frequency"`
	StartDate     *time.Time   `gorm:"type:date" json:"start_date,omitempty"`
	EndDate       *time.Time   `gorm:"type:date" json:"end_date,omitempty"`
	PaymentDate   *time.Time   `gorm:"type:date" json:"payment_date,omitempty"`
	Paid          bool         `gorm:"type:boolean" json:"paid"`
	Comment       string       `gorm:"type:varchar(255)" json:"comment"`
	UserID        uint         `json:"user_id"`
	LoanID        *uint        `json:"loan_id,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}
