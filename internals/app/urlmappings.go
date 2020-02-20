package app

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/controllers"
	"net/http"
)

func mapUrls(appConfig *config.AppConfig) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	downloadController:=controllers.NewDownloadController(appConfig)

	router.POST("/download", downloadController.Download)
	router.GET("/download/description/:grpID",downloadController.DownloadTableDetails)
	router.GET("/serve/:hash/:grpID/:uID",downloadController.ServeFiles)
	router.GET("/getFileID/:grpID",downloadController.GetFileID)


	//router.GET("/ws",downloadController.DisplayStatus)

}
