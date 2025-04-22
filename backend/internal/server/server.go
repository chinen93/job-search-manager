package server

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/jobs", getJobs)
	router.GET("/jobs/:id", getJobByID)
	router.POST("/jobs", postJobs)

	router.Run("localhost:8080")
}
