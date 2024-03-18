package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName       string
	LastName        string
	Email           string `gorm:"uniqueIndex"`
	Username        string `gorm:"uniqueIndex"`
	Password        string `json:"-"`
	Theme           string
	LastLogin       string
	Profile_picture string
}
