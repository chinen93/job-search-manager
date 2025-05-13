package datahelper

import (
	dao "job-search-manager/internal/dao"
	jobcontroller "job-search-manager/internal/jobs/controller"
	"log"
)

func ImportPositions() error {

	dao.Init()
	defer dao.Close()

	positionController := new(jobcontroller.PositionController)

	positions, err := parsePositions()
	if err != nil {
		return err
	}

	for _, position := range positions {
		log.Println(position)
		positionController.Create(position)
	}

	return nil
}
