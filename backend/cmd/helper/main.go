package main

import (
	config "job-search-manager/config"
	datahelper "job-search-manager/internal/data_helper"
	"log"
)

func main() {

	config.ConfigLog()
	config.ConfigEnv()

	log.Print("Helper")

	err := datahelper.ImportPositions()
	if err != nil {
		log.Fatal("Import Positions error: ", err)
	}
}
