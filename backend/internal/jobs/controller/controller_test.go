package jobcontroller

import (
	"job-search-manager/internal/dao"
	jobmodel "job-search-manager/internal/jobs/models"
	"testing"
)

func Test_Controllers(t *testing.T) {

	var positionController = new(PositionController)
	var jobController = new(JobController)
	var companyController = new(CompanyController)

	dao.Init()
	defer dao.Close()

	t.Run("Add Position", func(t *testing.T) {

		companyName := "Company A"
		jobID := "123ABC"
		jobTitle := "Position A"
		jobDescription := "Loren Ipsum"
		jobNotes := "Loren Ipsum"

		newPosition := new(jobmodel.Position)
		newPosition.Company = companyName
		newPosition.Title = jobTitle
		newPosition.ID = jobID
		newPosition.Description = jobDescription
		newPosition.Notes = jobNotes

		_, err := jobController.GetById(newPosition.ID)
		if err == nil {
			t.Fatal("Job already exists", err)
		}

		positionCompany, _ := positionController.Create(newPosition)
		t.Log(positionCompany)

		job, err := jobController.GetById(newPosition.ID)
		if job == nil {
			t.Fatal("Job not found", err)
		}
		if job.ID != jobID || job.Title != jobTitle || job.Description != jobDescription || job.Notes != jobNotes {
			t.Fatal("Job with different fields")
		}

		company, err := companyController.GetByName(newPosition.Company)
		if company == nil {
			t.Fatal("Company not found", err)
		}
		if company.Name != companyName {
			t.Fatal("Company with different fields")
		}

		if len(company.Jobs) != 1 {
			t.Fatal("Company has more jobs", len(company.Jobs))
		}

	})
}
