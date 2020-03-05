package app

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internal/config"
	"github.com/manoj2210/distributed-download-system-backend/internal/controllers"
	"net/http"
)

func mapUrls(appConfig *config.AppConfig) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	downloadController:=controllers.NewDownloadController(appConfig)

	api := router.Group(appConfig.Server.NAME)
	{
		api.POST("/download", downloadController.Download)
		api.GET("/download/description/:grpID", downloadController.DownloadTableDetails)
		api.GET("/serve/:grpID/:uID", downloadController.ServeFiles)
		api.GET("/getScheduler/:grpID",downloadController.GetScheduler)
		api.GET("/acknowledge/:grpID/:uID/:file",downloadController.Acknowledge)
	}
	//router.GET("/getFileID/:grpID",downloadController.GetFileID)



	//router.GET("/ws",downloadController.DisplayStatus)

}
