package testdata

import (
	jobcontroller "job-search-manager/internal/jobs/controller"
	jobmodel "job-search-manager/internal/jobs/models"
)

var companyController = new(jobcontroller.CompanyController)
var jobController = new(jobcontroller.JobController)

var Jobs = []jobmodel.Job{
	{ID: "1", Title: "Software Engineer", Description: "Lorem Ipsum 1", Notes: ""},
	{ID: "2", Title: "Software Engineer Full Stack", Description: "Lorem Ipsum 2", Notes: ""},
	{ID: "AB", Title: "Software Engineer Backend", Description: "Lorem Ipsum 3", Notes: ""},
}

var Companies = []jobmodel.Company{
	{ID: "1", Name: "Company A", Jobs: []jobmodel.Job{}},
	{ID: "2", Name: "Company B", Jobs: []jobmodel.Job{}},
	{ID: "3", Name: "Company C", Jobs: []jobmodel.Job{}},
}

var JobCompanyAssociation = map[*jobmodel.Job]*jobmodel.Company{
	&Jobs[0]: &Companies[0],
	&Jobs[1]: &Companies[1],
	&Jobs[2]: &Companies[1],
}

func InitTest() {

	for index, job := range Jobs {
		jobCreated, _ := jobController.Create(&job)
		Jobs[index] = *jobCreated
	}

	for _, company := range Companies {
		companyController.Create(&company)
	}

	for job, company := range JobCompanyAssociation {
		companyController.AddJob(company, job)
	}
}
