package server

import (
	"job-search-manager/internal/jobs"
	"job-search-manager/internal/testdata"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getJobs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testdata.Jobs)
}

func postJobs(c *gin.Context) {
	var newJob jobs.Job

	// Call BindJSON to bind the received JSON to
	// newJob
	if err := c.BindJSON(&newJob); err != nil {
		return
	}

	// Add the new job to the slice.
	testdata.Jobs = append(testdata.Jobs, newJob)
	c.IndentedJSON(http.StatusCreated, newJob)
}

// getJobByID locates the job whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getJobByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an job whose ID value matches the parameter.
	for _, a := range testdata.Jobs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "job not found"})
}
