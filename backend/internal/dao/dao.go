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
	DATABASE_FILENAME      = "internal/dao/job_search_manager_sqlite.db"
	TEST_DATABASE_FILENAME = "internal/testdata/test_sqlite.db"
)

var (
	DB *gorm.DB
)

func selectDatabaseFilename() (database_filename string) {
	envmode := os.Getenv(config.ENV_MODE)

	switch envmode {

	case config.TEST_MODE:
		log.Println("Using Test Database")
		os.Remove(TEST_DATABASE_FILENAME)
		database_filename = TEST_DATABASE_FILENAME

	case config.PROD_MODE:
		database_filename = DATABASE_FILENAME

	}

	return
}

func Init() (err error) {
	log.Println("Database Starting")

	database_filename := selectDatabaseFilename()
	database, err := gorm.Open(sqlite.Open(database_filename), &gorm.Config{})

	DB = database.Debug()

	database.AutoMigrate(&jobmodel.Job{}, &jobmodel.Company{})

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
