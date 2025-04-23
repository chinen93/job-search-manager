package dao

import (
	"log"
	"os"

	"job-search-manager/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	jobmodel "job-search-manager/internal/jobs/models"
)

const (
	DATABASE_FILENAME      = "internal/dao/job-search-manager-sqlite.db"
	TEST_DATABASE_FILENAME = "internal/testdata/test-sqlite.db"
)

var (
	DB *gorm.DB
)

func Init() (err error) {
	log.Println("Database Starting")

	envmode := os.Getenv(config.ENV_MODE)

	var database_filename string

	switch envmode {
	case config.TEST_MODE:
		log.Println("Using Test Database")
		os.Remove(TEST_DATABASE_FILENAME)
		database_filename = TEST_DATABASE_FILENAME
	case config.PROD_MODE:
		database_filename = DATABASE_FILENAME
	}

	database, err := gorm.Open(sqlite.Open(database_filename), &gorm.Config{})
	DB = database

	database.AutoMigrate(&jobmodel.Job{})

	if err != nil {
		log.Println("failed to connect database")
	}

	db, err := database.DB()
	if err != nil {
		return
	}

	log.Printf("Database # Connections: %+v", db.Stats().OpenConnections)

	return db.Ping()
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		panic("failed to connect database")
	}

	db.Close()
}
