package main

import (
	"user-service/helper"
	my_logger "user-service/logger"
	"user-service/models"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %v", err)
	}

	PSQL_CONNECTION := helper.GetENV("PSQL_CONNECTION", "")
	db, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: my_logger.GormLogger(),
	})
	if err != nil {
		logrus.Errorf("Error connecting to database: %v", err)
	}
	if err := db.Migrator().AutoMigrate(
		&models.User{},
	); err != nil {
		logrus.Errorf("Error migrating database table: %v", err)
	}

}
