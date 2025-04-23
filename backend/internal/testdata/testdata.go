package testdata

import (
	jobcontroller "job-search-manager/internal/jobs/controller"
	jobmodel "job-search-manager/internal/jobs/models"
)

var Jobs = []jobmodel.Job{
	{ID: "1", Title: "Software Engineer", Description: "Lorem Ipsum 1", Notes: ""},
	{ID: "2", Title: "Software Engineer Full Stack", Description: "Lorem Ipsum 2", Notes: ""},
	{ID: "AB", Title: "Software Engineer Backend", Description: "Lorem Ipsum 3", Notes: ""},
}

func InitTest() {
	for _, job := range Jobs {
		jobcontroller.CreateAJob(&job)
	}
}
