package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jkhammerseth/budget-tracker/backend/internal/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		password := os.Getenv("DB_PASSWORD")
		dsn := fmt.Sprintf("host=localhost user=postgres dbname=budget-tracker-db password=%s port=5432 sslmode=disable TimeZone=UTC", password)
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		err = db.AutoMigrate(&model.User{}, &model.Expense{}, &model.Income{})
		if err != nil {
			log.Fatalf("Failed to auto-migrate: %v", err)
		}
	})
}

func GetDB() *gorm.DB {
	if db == nil {
		InitDB() // Ensure the db is initialized before attempting to return it
	}
	return db
}
