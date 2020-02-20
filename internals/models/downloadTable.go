package models

import (
	"time"
)

const (
	Downloading=0
	Downloaded=1
	Error=2
	UploadedtoDB=3
)

var DownloadTable=make(map[string]*DownloadableFile)

type DownloadableFile struct {
	Url string    `json:"url"`
	Time time.Time  `json:"time"`
	Counter *WriteCounter  `json:"counter"`
	Status uint8    `json:"status"`
}

func NewDownloadableFile(p string)*DownloadableFile{
	return &DownloadableFile{p,time.Now(),&WriteCounter{},Downloading}
}

type WriteCounter struct {
	Total int64
	L int64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += (int64(n)/wc.L)*100
	//wc.PrintProgress()
	return n, nil
}

func AddNewDownloadableFile(grpId string,df *DownloadableFile){
	for key,value := range DownloadTable{
		if  time.Now().Sub(value.Time).Hours() >= 1{
			delete(DownloadTable,key)
		}
	}
	DownloadTable[grpId]=df
}
