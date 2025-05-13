package main

import (
	config "job-search-manager/config"
	datahelper "job-search-manager/internal/data_helper"
	"log"
)

func main() {

	config.Init()

	log.Print("Helper")

	err := datahelper.ImportPositions()
	if err != nil {
		log.Fatal("Import Positions error: ", err)
	}
}
