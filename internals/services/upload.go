package services

import "go.mongodb.org/mongo-driver/mongo"

type UploadService struct{
	Rls
  epo *mongo.Client
}

func NewUploadService(c *mongo.Client) *UploadService{
	return &UploadService{Repo:c}

}