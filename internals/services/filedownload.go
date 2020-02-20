package services

import (
	"fmt"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func DownloadFile(filepath string, url string, counter *models.WriteCounter) error {

	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter.L,_=strconv.Atoi(resp.Header.Get("Content-Length"))
	if counter.L==0{
		counter.L=1
	}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	fmt.Print("\n")
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func StartDownload(fileName string,fileUrl string,client *mongo.Client,file *models.DownloadableFileDescription) {

	//Adding to Download table
	fmt.Println("Download Started")
	_=os.Mkdir("downloads/"+fileName,os.ModePerm)
	fileName1:="downloads/"+fileName+"/"+fileName
	err := DownloadFile(fileName1, fileUrl,file.Counter)
	if err != nil {
		log.Println(err)
		os.Remove(fileName1+ ".tmp")
		//models.UpdateStatus(fileName,"Error",client)
		models.UpdateStatus(fileName,"Error")
		return
	}

	fmt.Println("Download Finished")
	//models.UpdateStatus(fileName,"Uploading",client)
	models.UpdateStatus(fileName,"Uploading")
	fmt.Println("Uploading")
	err= UploadFileToDB(fileName1,fileName,client)
	if err !=nil{
		log.Println(err)
	}
	//models.UpdateStatus(fileName,"Uploaded",client)
	models.UpdateStatus(fileName,"Uploaded")
}

