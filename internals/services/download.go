package services

import "go.mongodb.org/mongo-driver/mongo"

type DownloadService struct{
	repo *mongo.Collection
}

func NewDownloadService(collection *mongo.Collection) *DownloadService{
	return &DownloadService{repo:collection}

}