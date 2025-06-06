package server

import (
	"job-search-manager/internal/dao"
	jobserver "job-search-manager/internal/jobs/server"
	"job-search-manager/internal/testdata"

	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.Println("Server Starting")

	dao.Init()
	defer dao.Close()

	testdata.InitTest()

	router := gin.Default()
	router.GET("/jobs", jobserver.GetJobs)
	router.GET("/jobs/:id", jobserver.GetJobByID)
	router.POST("/jobs", jobserver.PostJobs)
	router.GET("/companies", jobserver.GetCompanies)
	router.GET("/companies/:id", jobserver.GetCompanyByID)
	router.POST("/companies", jobserver.PostCompanies)
	router.POST("/position", jobserver.PostPosition)

	router.Run("localhost:8080")

	log.Println("Server Started")
}
