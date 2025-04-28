package jobserver

import (
	jobcontroller "job-search-manager/internal/jobs/controller"
	jobmodel "job-search-manager/internal/jobs/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jobController = new(jobcontroller.JobController)

func GetJobs(c *gin.Context) {
	jobs, _ := jobController.GetAllJob()
	c.IndentedJSON(http.StatusOK, jobs)
}

func PostJobs(c *gin.Context) {
	var newJob jobmodel.Job

	// Call BindJSON to bind the received JSON to
	// newJob
	if err := c.BindJSON(&newJob); err != nil {
		return
	}

	// Add the new job to the slice.
	jobController.CreateAJob(&newJob)
	c.IndentedJSON(http.StatusCreated, newJob)
}

// getJobByID locates the job whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetJobByID(c *gin.Context) {
	id := c.Param("id")

	job, err := jobController.GetAJob(id)

	if err == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
	}

	c.IndentedJSON(http.StatusOK, job)
}
