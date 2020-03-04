package models

import (
	"time"
)


//var DownloadTable=make(map[string]*DownloadableFileDescription)

type DownloadableFileDescription struct {
	GroupId string `json:"groupID" bson:"groupID"`
	Url string    `json:"url" bson:"url""`
	Time time.Time  `json:"time" bson:"time""`
	Size int64  `json:"size" bson:"size""`
	Status string    `json:"status" bson:"status"`

}
//func UpdateStatus(grpID string,s string)error{
//      _,ok:=DownloadTable[grpID]
//      if ok{
//		  DownloadTable[grpID].Status=s
//		  return nil
//	  }
//	  return errors.New("Data Not available")
//}

func NewDownloadableFileDescription(u,g string)*DownloadableFileDescription{
	return &DownloadableFileDescription{g,u,time.Now(),0,"Downloading"}
}

//type WriteCounter struct {
//	Total int  `json:"total" bson:"total"`
//	L int      `json:"size" bson:"size"`
//	Percent string ``
//}

//func (wc *WriteCounter) Write(p []byte) (int, error) {
//	n := len(p)
//	wc.Total += n
//	wc.Percent=strconv.Itoa(int(float64(wc.Total/wc.L)*100))
//	return n, nil
//}

//func AddNewDownloadableFile(grpId string,df *DownloadableFileDescription)error{
//	for key,value := range DownloadTable{
//		if  time.Now().Sub(value.Time).Hours() >= 1{
//			delete(DownloadTable,key)
//		}
//	}
//	DownloadTable[grpId]=df
//	return nil
//}
//
//func GetDownloadableFile(grpID string)(*DownloadableFileDescription,error){
//	_,ok:=DownloadTable[grpID]
//	if ok{
//		return DownloadTable[grpID],nil
//	}
//	return nil,errors.New("Data Not available")
//}
