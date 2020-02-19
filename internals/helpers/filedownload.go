package helpers

import (
	"fmt"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"log"
	//"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	//"strings"
)



//func (wc WriteCounter) PrintProgress() {
//	fmt.Printf("\r%s", strings.Repeat(" ", 35))
//	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
//}

func StartDownload(fileName string,fileUrl string) {

	//Adding to Download table
	file:=models.NewDownloadableFile(fileUrl)
	models.AddNewDownloadableFile(fileName,file)

	fmt.Println("Download Started")
	_=os.Mkdir("downloads/"+fileName,os.ModePerm)
	fileName="downloads/"+fileName+"/"+fileName
	err := DownloadFile(fileName, fileUrl,file.Counter)
	if err != nil {
		log.Println(err)
		os.Remove(fileName+ ".tmp")
		models.DownloadTable[fileName].Status=models.Error
		return
	}



	models.DownloadTable[fileName].Status=models.UploadedtoDB
	fmt.Println("Download Finished")
}

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
	counter.L=resp.ContentLength
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

