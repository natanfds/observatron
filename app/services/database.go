package services

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDatabase - Função para inicializar o banco de dados
// Recebe um slice de models para realizar a migração
// Retorna um erro caso haja algum problema
func NewDatabase(modelsToMigrate []interface{}) (*gorm.DB, error) {
	dbConfig := sqlite.Open("./local.db")
	db, err := gorm.Open(dbConfig, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	for _, model := range modelsToMigrate {
		err = db.AutoMigrate(model)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
