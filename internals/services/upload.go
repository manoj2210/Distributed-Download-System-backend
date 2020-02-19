package services

import "go.mongodb.org/mongo-driver/mongo"

type UploadService struct{
	repo *mongo.Client
}

func NewUploadService(c *mongo.Client) *UploadService{
	return &UploadService{repo:c}

}