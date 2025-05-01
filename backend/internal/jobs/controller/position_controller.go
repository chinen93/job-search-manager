package jobcontroller

import (
	jobmodel "job-search-manager/internal/jobs/models"
	"log"
)

type PositionController struct{}

func (positionController *PositionController) Create(position *jobmodel.Position) (company *jobmodel.Company, err error) {

	companyController := new(CompanyController)
	jobController := new(JobController)

	job, err := jobController.GetById(position.ID)
	if job != nil {
		log.Printf("Job '%v' already exist", position.ID)
		return nil, err
	}

	company, _ = companyController.GetByName(position.Company)
	if company == nil {
		company = new(jobmodel.Company)
		company.Name = position.Company
		companyController.Create(company)
	}

	newJob := new(jobmodel.Job)
	newJob.ID = position.ID
	newJob.Title = position.Title
	newJob.Notes = position.Notes
	newJob.Description = position.Description

	jobController.Create(newJob)

	companyController.AddJob(company, newJob)

	return company, nil
}
