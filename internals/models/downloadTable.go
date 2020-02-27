package models

import (
	"errors"
	"strconv"
	"time"
)


var DownloadTable=make(map[string]*DownloadableFileDescription)

type DownloadableFileDescription struct {
	Url string    `json:"url" bson:"url""`
	Time time.Time  `json:"time" bson:"time""`
	Counter *WriteCounter  `json:"counter" bson:"counter""`
	Status string    `json:"status" bson:"status"`

}
func UpdateStatus(grpID string,s string)error{
      _,ok:=DownloadTable[grpID]
      if ok{
		  DownloadTable[grpID].Status=s
		  return nil
	  }
	  return errors.New("Data Not available")
}

func NewDownloadableFileDescription(p string)*DownloadableFileDescription{
	return &DownloadableFileDescription{p,time.Now(),&WriteCounter{},"Downloading"}
}

type WriteCounter struct {
	Total int
	L int
	Percent string
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += n
	wc.Percent=strconv.Itoa(int((wc.Total/wc.L)*100))
	return n, nil
}

func AddNewDownloadableFile(grpId string,df *DownloadableFileDescription)error{
	for key,value := range DownloadTable{
		if  time.Now().Sub(value.Time).Hours() >= 1{
			delete(DownloadTable,key)
		}
	}
	DownloadTable[grpId]=df
	return nil
}

func GetDownloadableFile(grpID string)(*DownloadableFileDescription,error){
	_,ok:=DownloadTable[grpID]
	if ok{
		return DownloadTable[grpID],nil
	}
	return nil,errors.New("Data Not available")
}
