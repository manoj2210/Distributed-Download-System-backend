package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internal/config"
	"github.com/manoj2210/distributed-download-system-backend/internal/errors"
	"github.com/manoj2210/distributed-download-system-backend/internal/helpers"
	"github.com/manoj2210/distributed-download-system-backend/internal/models"
	"github.com/manoj2210/distributed-download-system-backend/internal/services"
	"log"
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

	s:=models.NewScheduler(post.Url,post.GroupID)
	er:=ctrl.DownloadService.AddNewScheduler(s)
	if er!=nil{
		restErr:= errors.NewBadRequestError("Unable to Add a New Scheduler")
		c.JSON(restErr.Status, restErr)
		return
	}

	f:=models.NewDownloadableFileDescription(post.Url,post.GroupID)
	er=ctrl.DownloadService.AddNewDownloadableFile(f)
	if er!=nil{
		restErr:= errors.NewBadRequestError("Unable to Add a New Group ")
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

func (ctrl *DownloadController) ServeFiles(c *gin.Context) {
	uID:=c.Param("uID")
	grpID:=c.Param("grpID")

	//Scheduler Part 
	var file int64
	s,err:=ctrl.DownloadService.GetScheduler(grpID)
	if err != nil {
		restErr := errors.NewNotFoundError("No such GroupID0")
		c.JSON(restErr.Status, restErr)
		return
	}
	if s.Ptr+1==s.TotalChunks{
		f:=ctrl.DownloadService.CheckSchedulerForHoles(s)
		if f==-1{	
			restErr := errors.NewNotFoundError("All data scheduled")
			c.JSON(restErr.Status, restErr)
			return
		}
		file=f
	} else{
		s.Ptr += 1
		file=s.Ptr
	}
	r:=models.NewRecord(uID,file)
	s.Data = append(s.Data, r)
	err= ctrl.DownloadService.UpdateScheduler(s)
	if err != nil {
		log.Println(err)
		restErr := errors.NewNotFoundError("No such GroupID1")
		c.JSON(restErr.Status, restErr)
		return
	}
	m:=grpID+":"+strconv.Itoa(int(file-1))
	k, err := ctrl.DownloadService.ServeFile(m)
	if err != nil {
		restErr := errors.NewNotFoundError("No such GroupID2")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+m)
	c.Writer.Write(k.Bytes())
	return
}

func (ctrl *DownloadController) AckAndServe(c *gin.Context) {
	uID:=c.Param("uID")
	grpID:=c.Param("grpID")
	i:=c.Param("file")
	t, err := strconv.Atoi(i)
	f:=int64(t)

	//Scheduler Part
	var file int64
	s,err:=ctrl.DownloadService.GetScheduler(grpID)
	if err != nil {
		restErr := errors.NewNotFoundError("No such GroupID0")
		c.JSON(restErr.Status, restErr)
		return
	}

	if(f!=0){
		//ACK
		flag:=false
		for _,x:=range s.Data{
			if x.UserID==uID && x.FileNo==f {
				flag=true
			}
		}
		if flag {
			for idx, _ := range s.Data {
				if s.Data[idx].UserID == uID && s.Data[idx].FileNo == f {
					s.Data[idx].Acknowledged = true
				}
				if s.Data[idx].UserID != uID && s.Data[idx].FileNo == f {
					s.Data[idx].FileNo *= -1
				}
			}
		} else {
			restErr := errors.NewNotFoundError("No such Id with give Group and UID")
			c.JSON(restErr.Status, restErr)
			return
		}
	}

	if s.Ptr+1==s.TotalChunks{
		f:=ctrl.DownloadService.CheckSchedulerForHoles(s)
		if f==-1{
			restErr := errors.NewNotFoundError("All data scheduled")
			c.JSON(restErr.Status, restErr)
			return
		}
		file=f
	} else{
		s.Ptr += 1
		file=s.Ptr
	}
	r:=models.NewRecord(uID,file)
	s.Data = append(s.Data, r)
	err= ctrl.DownloadService.UpdateScheduler(s)
	if err != nil {
		log.Println(err)
		restErr := errors.NewNotFoundError("No such GroupID1")
		c.JSON(restErr.Status, restErr)
		return
	}
	m:=grpID+":"+strconv.Itoa(int(file-1))
	k, err := ctrl.DownloadService.ServeFile(m)
	if err != nil {
		restErr := errors.NewNotFoundError("No such GroupID2")
		c.JSON(restErr.Status, restErr)
		return
	}
	m=grpID+":"+strconv.Itoa(int(file))
	c.Header("Content-Disposition", "attachment; filename="+m)
	c.Writer.Write(k.Bytes())
	return
}

func (ctrl *DownloadController) GetScheduler(c *gin.Context){
	grpID:=c.Param("grpID")
	m,err:=ctrl.DownloadService.GetScheduler(grpID)
	if err != nil{
		restErr:= errors.NewNotFoundError("No Data available with that groupID")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK,m)
}

func (ctrl *DownloadController) Acknowledge(c *gin.Context){
	grpID:=c.Param("grpID")
	uID:=c.Param("uID")
	f:=c.Param("file")
	i, err := strconv.Atoi(f)
	err=ctrl.DownloadService.AcknowledgeScheduler(int64(i),grpID,uID)
	if err!=nil{
		restErr:= errors.NewNotFoundError("No File with fileId "+f+" available with that groupID")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, helpers.DownloadSuccess())
}

func (ctrl *DownloadController)Delete(c *gin.Context) {
	grpID:=c.Param("grpID")
	err:=ctrl.DownloadService.DeleteFiles(grpID)
	if err!=nil{
		restErr:= errors.NewInternalServerError("No Data for that groupId")
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, helpers.DownloadSuccess())
}