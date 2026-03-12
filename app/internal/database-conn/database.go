package databaseConn

import (
	"log"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:123456@localhost:5433/app"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Expense{})

	return db
}
