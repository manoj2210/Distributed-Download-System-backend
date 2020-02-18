package controllers

import (
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/errors"
	"github.com/manoj2210/distributed-download-system-backend/internals/helpers"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
)

type DownloadController struct{
	DownloadService *services.DownloadService
}

func (ctrl *DownloadController)Download(c *gin.Context) {
	post := models.DownloadPOSTRequest{}
	if err := c.ShouldBindJSON(&post); err != nil {
		restErr := errors.NewBadRequestError("invalid request body")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, helpers.DownloadSuccess())
}

func NewDownloadController(config *config.AppConfig) *DownloadController{
	return &DownloadController{DownloadService:services.NewDownloadService(config.Downloads)}

}
