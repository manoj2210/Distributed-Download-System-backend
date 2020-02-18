package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/errors"
	"github.com/manoj2210/distributed-download-system-backend/internals/helpers"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
)

func Download(c *gin.Context) {
	sam := models.DownloadPOSTRequest{}
	if err := c.ShouldBindJSON(&sam); err != nil {
		restErr := errors.NewBadRequestError("invalid request body")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(sam)
	//Check whether downloadable and download,push to db

	c.JSON(http.StatusOK, helpers.DownloadSuccess())
}
