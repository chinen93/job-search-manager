package jobserver

import (
	jobcontroller "job-search-manager/internal/jobs/controller"
	jobmodel "job-search-manager/internal/jobs/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

var positionController = new(jobcontroller.PositionController)

func PostPosition(c *gin.Context) {
	var newPosition jobmodel.Position

	// Call BindJSON to bind the received JSON to
	// newPosition
	if err := c.BindJSON(&newPosition); err != nil {
		return
	}

	result, err := positionController.Create(&newPosition)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Position already exist"})
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}
