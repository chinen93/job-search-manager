package testdata

import (
	jobcontroller "job-search-manager/internal/jobs/controller"
	jobmodel "job-search-manager/internal/jobs/models"
	"log"
)

var companyController = new(jobcontroller.CompanyController)
var jobController = new(jobcontroller.JobController)

func InitTest() {
	var Jobs = []jobmodel.Job{
		{ID: "1", Title: "Software Engineer", Description: "Lorem Ipsum 1", Notes: ""},
		{ID: "2", Title: "Software Engineer Full Stack", Description: "Lorem Ipsum 2", Notes: ""},
		{ID: "AB", Title: "Software Engineer Backend", Description: "Lorem Ipsum 3", Notes: ""},
	}

	var Companies = []jobmodel.Company{
		{ID: "1", Name: "Company A"},
		{ID: "2", Name: "Company B"},
		{ID: "3", Name: "Company C"},
	}

	var JobCompanyAssociation = map[string]jobmodel.Company{
		Jobs[0].ID: Companies[0],
		Jobs[1].ID: Companies[1],
		Jobs[2].ID: Companies[1],
	}

	for _, job := range Jobs {
		jobController.Create(&job)
	}

	for _, company := range Companies {
		companyController.Create(&company)
	}

	for jobID, company := range JobCompanyAssociation {

		log.Println(company)
		companyController.AddJob(&company, jobID)

		expected, _ := companyController.Get(company.ID)
		log.Println(expected)
	}

}
