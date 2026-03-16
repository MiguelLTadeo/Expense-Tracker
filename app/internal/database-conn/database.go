package databaseConn

import (
	"log"

	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/service/models"
	"github.com/MiguelLTadeo/Expense-Tracker.git/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := utils.GoDotEnvVariable("CONN_STR")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Expense{})

	return db
}
