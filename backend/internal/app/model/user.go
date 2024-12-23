package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Username  string `json:"username" gorm:"uniqueIndex"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login,omitempty"`
}
