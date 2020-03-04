package models

import (
	"time"
)

type DownloadableFileDescription struct {
	GroupId string `json:"groupID" bson:"groupID"`
	Url string    `json:"url" bson:"url""`
	Time time.Time  `json:"time" bson:"time""`
	Size int64  `json:"size" bson:"size""`
	Status string    `json:"status" bson:"status"`

}

func NewDownloadableFileDescription(u,g string)*DownloadableFileDescription{
	return &DownloadableFileDescription{g,u,time.Now(),0,"Downloading"}
}

