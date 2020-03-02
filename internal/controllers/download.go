package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internal/config"
	"github.com/manoj2210/distributed-download-system-backend/internal/errors"
	"github.com/manoj2210/distributed-download-system-backend/internal/helpers"
	"github.com/manoj2210/distributed-download-system-backend/internal/models"
	"github.com/manoj2210/distributed-download-system-backend/internal/services"
	"net/http"
	//	"bytes"
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

	f:=models.NewDownloadableFileDescription(post.Url,post.GroupID)
	er:=ctrl.DownloadService.AddNewDownloadableFile(f)
	if er!=nil{
		restErr:= errors.NewBadRequestError("Unable to Add a New Group")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, helpers.DownloadSuccess())

	//Create a downloading queue table and set the writeCounter then access with websocket

	go ctrl.DownloadService.DownloadFile(f)
}

func (ctrl *DownloadController)DownloadTableDetails(c *gin.Context) {
	grpID:=c.Param("grpID")
	m,err:=ctrl.DownloadService.GetDownloadableFile(grpID)
	if err != nil{
		restErr:= errors.NewNotFoundError("No Data available with that groupID")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK,m)
}

//func (ctrl *DownloadController)GetFileID(c *gin.Context) {
//	grpID:=c.Param("grpID")
//	o,er:=ctrl.DownloadService.FindDownloadableFile(grpID)
//	if er != nil{
//		restErr:= errors.NewNotFoundError("No such file")
//		c.JSON(restErr.Status, restErr)
//		return
//	}
//	c.JSON(http.StatusOK,o)
//}

func (ctrl *DownloadController) ServeFiles(c *gin.Context) {
	//hash:=c.Param("hash")
	uID:=c.Param("uID")
	grpID:=c.Param("grpID")
	//_,ok:=models.SchedulerArray[grpID]
	//if ok {
	//	n := models.SchedulerArray[grpID].Allocate(uID)
	//	fmt.Println(models.SchedulerArray[grpID])
	//	if n!=-1{
	//
	//	}else{
	//		restErr := errors.NewNotFoundError("No Data")
	//		c.JSON(restErr.Status, restErr)
	//		return
	//	}
	//}
	m:=grpID+":"+uID
	k, err := ctrl.DownloadService.ServeFile(m)
	if err != nil {
		restErr := errors.NewNotFoundError("No such GroupID")
		c.JSON(restErr.Status, restErr)
		return
	}
	//l:=bytes.NewReader(k.Bytes())
	//extraHeaders := map[string]string{
	//	"Content-Disposition": `attachment; filename="gopher.png"`,
	//}
	//c.DataFromReader(http.StatusOK,1000000,"application/octet-stream",l,extraHeaders)
	c.Writer.Write(k.Bytes())
	return

	restErr := errors.NewNotFoundError("No Data")
	c.JSON(restErr.Status, restErr)
	return
}

