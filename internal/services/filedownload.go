package services

import (
	"fmt"
	"github.com/manoj2210/distributed-download-system-backend/internal/models"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func DownloadFile(filepath string, url string,counter *models.WriteCounter) error {

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
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func StartDownload(d *DownloadService,file *models.DownloadableFileDescription) {
	fileName:=file.GroupId
	fileUrl:=file.Url

	_=os.Mkdir("downloads/"+fileName,os.ModePerm)
	filepath:="downloads/"+fileName+"/"+fileName
	err := DownloadFile(filepath, fileUrl,file.Counter)
	if err != nil {
		log.Println(err)
		os.Remove(filepath+ ".tmp")
		d.UpdateStatus(fileName,"Error")
		return
	}

	fi, err := os.Stat(filepath)
	if err != nil {
		log.Fatal("error with fileSize")
	}
	size := fi.Size()

	file.Counter.L=int(size)
	d.UpdateStatus(fileName,"Splitting")

	out,err:=exec.Command("split","-b","1m",filepath).Output()
	fmt.Println(out)
	if err !=nil{
		log.Println(err)
	}

	out,err=exec.Command("rm",filepath).Output()
	fmt.Println(out)
	if err !=nil{
		log.Println(err)
	}

	out,err=exec.Command("ls","-S","downloads/"+fileName).Output()
	str:=string(out)

	strings.ReplaceAll(str," ","")
	g:=strings.Split(str,"\n")
	//fmt.Println(g)

	if err !=nil{
		log.Println(err)
	}

	d.UpdateStatus(fileName,"Uploading")
	for idx,itm := range g{
		go UploadFileToDB("downloads/"+fileName+"/"+itm,fileName+":"+strconv.Itoa(idx),d.repo)
	}

	d.UpdateStatus(fileName,"Uploaded")
}

