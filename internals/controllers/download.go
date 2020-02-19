package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/errors"
	"github.com/manoj2210/distributed-download-system-backend/internals/helpers"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"github.com/manoj2210/distributed-download-system-backend/internals/services"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type DownloadController struct{
	DownloadService *services.DownloadService
	UploadService   *services.UploadService
}

func NewDownloadController(c *config.AppConfig) *DownloadController{
	return &DownloadController{DownloadService:services.NewDownloadService(c.Downloads),UploadService:services.NewUploadService(c.DB)}
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
	c.JSON(http.StatusCreated, helpers.DownloadSuccess())

	//Create a downloading queue table and set the writeCounter then access with websocket

	go helpers.StartDownload(post.GroupID,post.Url)
}

func (ctrl *DownloadController)DownloadtableDetails(c *gin.Context) {
	grpID:=c.Param("grpID")
	c.JSON(http.StatusOK,models.DownloadTable[grpID])
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


