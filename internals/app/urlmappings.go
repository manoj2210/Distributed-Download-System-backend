package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/controllers"
	models "github.com/manoj2210/distributed-download-system-backend/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func mapUrls(appConfig *config.AppConfig) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	downloadController:=controllers.NewDownloadController(appConfig)

	router.POST("/download", downloadController.Download)
	router.GET("/download/description/:grpID",downloadController.DownloadtableDetails)
	//router.GET("/ws",downloadController.DisplayStatus)

}
