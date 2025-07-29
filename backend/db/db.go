package db

import (
	"errors"
	"fmt"
	"trip-planner-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Connect() (err error) {
	// Get config
	appConfig, err := config.GetConfig()
	if err != nil {
		return
	}

	// Connect to db
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", appConfig.DBHost, appConfig.DBUserName, appConfig.DBUserPassword, appConfig.DBName, appConfig.DBPort, appConfig.DBTimeZone)

	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return
}

func GetDB() (gorm.DB, error) {
	if dbInstance == nil {
		return gorm.DB{}, errors.New("not connected to database")
	}
	return *dbInstance, nil
}
