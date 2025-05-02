package jobcontroller

import (
	"job-search-manager/internal/dao"
	jobmodel "job-search-manager/internal/jobs/models"
	"testing"
)

const (
	COMPANY_A = "Company A"
	TITLE_A   = "Position A"
	DESC_A    = "Loren Ipsum"
	NOTES_A   = "Loren Ipsum"
	JOB_ID_1  = "123ABC"
	JOB_ID_2  = "456DEF"
)

func assertJobEqual(t *testing.T, job *jobmodel.Job, expected *jobmodel.Job) {
	t.Helper()
	if job == nil {
		t.Fatal("Expected job, got nil")
	}
	if job.ID != expected.ID || job.Title != expected.Title || job.Description != expected.Description || job.Notes != expected.Notes {
		t.Fatalf("Job fields do not match\nExpected: %+v\nGot: %+v", expected, job)
	}
}

func assertCompany(t *testing.T, company *jobmodel.Company, expected *jobmodel.Company) {
	t.Helper()
	if company == nil {
		t.Fatal("Expected company, got nil")
	}

	if company.Name != expected.Name {
		t.Fatalf("Company fields do not match\nExpected: %+v\nGot: %+v", expected.Name, company.Name)
	}
	if len(company.Jobs) != len(expected.Jobs) {
		t.Fatalf("Expected %d jobs in Company, got %d", len(expected.Jobs), len(company.Jobs))
	}
	for i, job := range company.Jobs {
		if job.ID != expected.Jobs[i].ID {
			t.Fatalf("Expected job ID '%s' at position %d, got '%s'", expected.Jobs[i].ID, i, job.ID)
		}
	}
}

func Test_Controllers(t *testing.T) {

	var positionController = new(PositionController)
	var jobController = new(JobController)
	var companyController = new(CompanyController)

	dao.Init()
	defer dao.Close()

	t.Run("Add Position", func(t *testing.T) {

		newPosition := jobmodel.MakePosition(COMPANY_A, JOB_ID_1, TITLE_A, DESC_A, NOTES_A)

		job, err := jobController.GetById(newPosition.ID)
		if job != nil {
			t.Fatal("Job already exists", err)
		}

		positionCompany, _ := positionController.Create(newPosition)
		if len(positionCompany.Jobs) != 1 {
			t.Fatal("Expected only one job for the company")
		}
		if positionCompany.Jobs[0].ID != JOB_ID_1 {
			t.Fatal("Not expected job")
		}
		if positionCompany.Name != COMPANY_A {
			t.Fatal("Not expected company")
		}

		job, err = jobController.GetById(newPosition.ID)
		if job == nil {
			t.Fatal("Job not found", err)
		}
		expectedJob := jobmodel.Job{ID: JOB_ID_1, Title: TITLE_A, Description: DESC_A, Notes: NOTES_A}
		assertJobEqual(t, job, &expectedJob)

		company, err := companyController.GetByName(newPosition.Company)
		if company == nil {
			t.Fatal("Company not found", err)
		}
		expectedCompany := jobmodel.Company{Name: COMPANY_A, Jobs: []jobmodel.Job{{ID: JOB_ID_1}}}
		assertCompany(t, company, &expectedCompany)
	})

	t.Run("Do Not Add Duplicated Position", func(t *testing.T) {

		newPosition := jobmodel.MakePosition(COMPANY_A, JOB_ID_1, TITLE_A, DESC_A, NOTES_A)

		job, err := jobController.GetById(newPosition.ID)
		if job == nil {
			t.Fatal("Job does not exists", err)
		}

		positionCompany, _ := positionController.Create(newPosition)
		if positionCompany != nil {
			t.Fatal("Position should be already created")
		}

		job, err = jobController.GetById(newPosition.ID)
		if job == nil {
			t.Fatal("Job not found", err)
		}
		expectedJob := jobmodel.Job{ID: JOB_ID_1, Title: TITLE_A, Description: DESC_A, Notes: NOTES_A}
		assertJobEqual(t, job, &expectedJob)

		company, err := companyController.GetByName(newPosition.Company)
		if company == nil {
			t.Fatal("Company not found", err)
		}
		expectedCompany := jobmodel.Company{Name: COMPANY_A, Jobs: []jobmodel.Job{{ID: JOB_ID_1}}}
		assertCompany(t, company, &expectedCompany)
	})

	t.Run("Add Position Same Company", func(t *testing.T) {

		newPosition := jobmodel.MakePosition(COMPANY_A, JOB_ID_2, TITLE_A, DESC_A, NOTES_A)

		job, err := jobController.GetById(newPosition.ID)
		if job != nil {
			t.Fatal("Job already exists", err)
		}

		positionCompany, _ := positionController.Create(newPosition)
		if len(positionCompany.Jobs) != 2 {
			t.Fatal("Expected two jobs for the company")
		}
		if positionCompany.Jobs[1].ID != JOB_ID_2 {
			t.Fatal("Not expected job")
		}
		if positionCompany.Name != COMPANY_A {
			t.Fatal("Not expected company")
		}

		job, err = jobController.GetById(newPosition.ID)
		if job == nil {
			t.Fatal("Job not found", err)
		}
		expectedJob := jobmodel.Job{ID: JOB_ID_2, Title: TITLE_A, Description: DESC_A, Notes: NOTES_A}
		assertJobEqual(t, job, &expectedJob)

		company, err := companyController.GetByName(newPosition.Company)
		if company == nil {
			t.Fatal("Company not found", err)
		}
		expectedCompany := jobmodel.Company{Name: COMPANY_A, Jobs: []jobmodel.Job{{ID: JOB_ID_1}, {ID: JOB_ID_2}}}
		assertCompany(t, company, &expectedCompany)

	})
}
