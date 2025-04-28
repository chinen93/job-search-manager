package jobserver

import (
	jobcontroller "job-search-manager/internal/jobs/controller"
	jobmodel "job-search-manager/internal/jobs/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var companyController = new(jobcontroller.CompanyController)

func GetCompanies(c *gin.Context) {
	companies, _ := companyController.GetAllCompanies()
	log.Println(companies)
	c.IndentedJSON(http.StatusOK, companies)
}

func PostCompanies(c *gin.Context) {
	var newCompany jobmodel.Company

	// Call BindJSON to bind the received JSON to
	// newCompany
	if err := c.BindJSON(&newCompany); err != nil {
		return
	}

	// Add the new job to the slice.
	companyController.CreateACompany(&newCompany)
	c.IndentedJSON(http.StatusCreated, newCompany)
}

// getCompanyByID locates the job whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")

	job, err := companyController.GetACompany(id)

	if err == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "company not found"})
	}

	c.IndentedJSON(http.StatusOK, job)
}
