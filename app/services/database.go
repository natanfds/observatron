package services

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

var modelsToMigrate = []interface{}{}

func StartDatabase() error {
	dbConfig := sqlite.Open("./local.db")
	db, err := gorm.Open(dbConfig, &gorm.Config{})
	if err != nil {
		return err
	}

	DATABASE = db

	for _, model := range modelsToMigrate {
		err = db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
