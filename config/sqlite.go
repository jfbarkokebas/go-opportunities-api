package config

import (
	"os"

	"github.com/jfbarkokebas/go-opportunities-api.git/schemas"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite" // no terminal: go mod tidy
)

func checkBeforeCreateDbFile(dbPath string) error {

	//Checking if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not foun. Creating ...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}

		file.Close()
	}

	return nil
}

func InitializeSQLite() (*gorm.DB, error) {

	logger := GetLogger("sqlite ...")
	dbPath := "./db/main.db"

	err := checkBeforeCreateDbFile(dbPath)
	if err != nil {
		logger.Errorf("Failure to create database file error: %v", err)
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SQLite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
