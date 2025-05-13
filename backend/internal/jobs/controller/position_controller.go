package jobcontroller

import (
	"errors"
	jobmodel "job-search-manager/internal/jobs/models"
	"log"
)

type PositionController struct{}

func (positionController *PositionController) Create(position *jobmodel.Position) (company *jobmodel.Company, err error) {

	companyController := new(CompanyController)
	jobController := new(JobController)

	job, _ := jobController.GetById(position.ID)
	if job != nil {
		log.Printf("Job '%v' already exist", position.ID)
		return nil, errors.New("job already exists")
	}

	company, _ = companyController.GetByName(position.Company)
	if company == nil {
		company = jobmodel.MakeCompany(position.Company)
		companyController.Create(company)
	}

	newJob := jobmodel.MakeJob(position.ID, position.Title, position.Description, position.Notes)
	jobController.Create(newJob)

	companyController.AddJob(company, newJob)

	return company, nil
}
