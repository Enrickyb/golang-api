package config

import (
	"github.com/Enrickyb/golang-api/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {

	logger := GetLogger("postgres")

	db, err := gorm.Open(postgres.Open("host=localhost port=32768 user=postgres dbname=postgres password=postgrespw sslmode=disable"), &gorm.Config{})

	if err != nil {
		logger.ErrorF("error database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("error automigration: %v", err)
		return nil, err
	}

	return db, nil

}
