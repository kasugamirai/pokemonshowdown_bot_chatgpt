package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"xy.com/pokemonshowdownbot/config"
	"xy.com/pokemonshowdownbot/models"
)

var (
	DB *gorm.DB
)

// InitDB initializes the database connection.
func InitDB() error {
	var err error
	dbConfig := config.Instance.DatabaseDSN

	// Create a new SQLite database connection
	DB, err = gorm.Open(sqlite.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the data models.
	err = migrateModels()
	if err != nil {
		return err
	}

	return nil
}

func migrateModels() error {
	err := DB.AutoMigrate(
		&models.Stickers{},
	)
	if err != nil {
		return err
	}

	return nil
}
