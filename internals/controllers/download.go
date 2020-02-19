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
	if err:=helpers.ValidateDownloadRequest(&post);err!=nil{
		restErr:= errors.NewBadRequestError("invalid URL")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, helpers.DownloadSuccess())
	go helpers.StartDownload(post.GroupID,post.Url,&helpers.WriteCounter{})
}


//func Echo(ws *websocket.Conn) {
//
//	helpers.DownloadFile()
//	for {
//		if err := websocket.JSON.Send(ws, info); err != nil {
//			fmt.Println("Error sending message")
//			break
//		}
//		// if BytesTransferred == 100 break
//	}
//}


//func (ctrl *DownloadController)DisplayStatus(c *gin.Context){
//	handler := websocket.Handler(Echo)
//	handler.ServeHTTP(c.Writer, c.Request)
//}

func NewDownloadController(config *config.AppConfig) *DownloadController{
	return &DownloadController{DownloadService:services.NewDownloadService(config.Downloads)}
}
