package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/errors"
	"github.com/manoj2210/distributed-download-system-backend/internals/helpers"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"github.com/manoj2210/distributed-download-system-backend/internals/services"
	"net/http"
	"strconv"
)

type DownloadController struct{
	DownloadService   *services.DownloadService
}

func NewDownloadController(c *config.AppConfig) *DownloadController{
	return &DownloadController{DownloadService:services.NewDownloadService(c.DB)}
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

	f:=models.NewDownloadableFileDescription(post.Url)
	er:=models.AddNewDownloadableFile(post.GroupID,f)
	if er!=nil{
		restErr:= errors.NewBadRequestError("Unable to insert to DB")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, helpers.DownloadSuccess())

	//Create a downloading queue table and set the writeCounter then access with websocket

	go ctrl.DownloadService.DownloadFile(post.GroupID,post.Url,f)
}

func (ctrl *DownloadController)DownloadTableDetails(c *gin.Context) {
	grpID:=c.Param("grpID")
	//m,err:=ctrl.DownloadService.GetDownloadableFile(grpID)
	m,err:=models.GetDownloadableFile(grpID)
	if err != nil{
		restErr:= errors.NewNotFoundError("No such GroupID")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK,m)
}

func (ctrl *DownloadController)GetFileID(c *gin.Context) {
	grpID:=c.Param("grpID")
	o,er:=ctrl.DownloadService.FindDownloadableFile(grpID)
	if er != nil{
		restErr:= errors.NewNotFoundError("No such file")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK,o)
}

func (ctrl *DownloadController) ServeFiles(c *gin.Context) {
	hash:=c.Param("hash")
	n,_:=strconv.Atoi(c.Param("n"))
	k,err:=ctrl.DownloadService.ServeFile(hash,n)
	if err!=nil{
		restErr:= errors.NewNotFoundError("No such GroupID")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK,k)
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


