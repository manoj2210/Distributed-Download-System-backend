package helpers

import (
	"fmt"
	"log"

	//"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	//"strings"
)

type WriteCounter struct {
	Total int64
	l int64
}



func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += (int64(n)/wc.l)*100
	//wc.PrintProgress()
	return n, nil
}

//func (wc WriteCounter) PrintProgress() {
//	fmt.Printf("\r%s", strings.Repeat(" ", 35))
//	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
//}

func StartDownload(fileName string,fileUrl string,counter *WriteCounter) {
	fmt.Println("Download Started")
	os.Mkdir("downloads/"+fileName,os.ModePerm)
	fileName="downloads/"+fileName+"/"+fileName
	err := DownloadFile(fileName, fileUrl,counter)
	if err != nil {
		log.Println(err)
		os.Remove(fileName+ ".tmp")
		return
	}

	fmt.Println("Download Finished")
}

func DownloadFile(filepath string, url string, counter *WriteCounter) error {

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
	counter.l=resp.ContentLength
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

