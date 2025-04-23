package dao

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	jobmodel "job-search-manager/internal/jobs/models"
)

const (
	DATABASE_FILENAME = "job-search-manager-sqlite.db"
)

var (
	DB *gorm.DB
)

func Init() (err error) {
	log.Println("Database Starting")

	database, err := gorm.Open(sqlite.Open(DATABASE_FILENAME), &gorm.Config{})
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
